package main

import (
	"io"
	"github.com/gorilla/mux"
	"flag"
	"time"
	"github.com/go-kit/kit/log"
	"os"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"context"
	"github.com/feng/future/go-kit/microsvr/apigateway/pkg/addendpoint"
	"github.com/feng/future/go-kit/microsvr/apigateway/pkg/addservice"
	"github.com/feng/future/go-kit/microsvr/apigateway/pkg/addtransport"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"google.golang.org/grpc"
	"github.com/go-kit/kit/sd/lb"
)

func main() {
	var (
		// httpAddr     = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
		// consulAddr   = flag.String("consul.addr", "", "Consul agent address")
		httpAddr     = flag.String("httpaddr", ":8000", "Address for HTTP (JSON) server")
		consulAddr   = flag.String("consuladdr", "", "Consul agent address")
		retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
		//retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
		retryTimeout = flag.Duration("retry.timeout", 2*time.Second, "per-request timeout, including retries")
	)
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(*consulAddr) > 0 {
			consulConfig.Address = *consulAddr
		}
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	// Transport domain.
	tracer := stdopentracing.GlobalTracer() //no-op
	zipkinTracer, _ := stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))	
	ctx := context.Background()
	r := mux.NewRouter()

	{
		var (
			tags = []string{"appsvc"}
			passingOnly = true
			endpoints = addendpoint.Set{}
			instancer = consulsd.NewInstancer(client, logger, "addsvc", tags, passingOnly)
		)
		{
			factory := addsvcFactory(addendpoint.MakeSumEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			endpoints.SumEndpoint = retry
		}
	}
}

func addsvcFactory(makeEndpoint func(addservice.Service) endpoint.Endpoint, tracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := addtransport.NewGRPCClient(conn, tracer, zipkinTracer, logger)
		endpoint := makeEndpoint(service)
		return endpoint, conn, nil
	}
}

