package apigateway

import (
	"github.com/go-kit/kit/endpoint"
	consulsd "github.com/go-kit/kit/sd/consul"
	"flag"
	"time"
	"github.com/hashicorp/consul/api"
	kitlog "github.com/go-kit/kit/log"
	"log"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"context"
	"os"
)
var (
	// httpAddr     = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
	// consulAddr   = flag.String("consul.addr", "", "Consul agent address")
	httpAddr     = flag.String("httpaddr", ":8000", "Address for HTTP (JSON) server")
	consulAddr   = flag.String("consuladdr", "", "Consul agent address")
	retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
	//retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
	retryTimeout = flag.Duration("retry.timeout", 2*time.Second, "per-request timeout, including retries")
)

func init() {
	consulConfig := api.DefaultConfig()
	if len(*consulAddr) > 0 {
		consulConfig.Address = *consulAddr
	}
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}
	client = consulsd.NewClient(consulClient)

	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	
}

var logger kitlog.Logger


var client consulsd.Client
// Transport domain.
var tracer = stdopentracing.GlobalTracer() //no-op
var zipkinTracer, _ = stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))	
var ctx = context.Background()

type GatewayService interface {
	GetAccount() endpoint.Endpoint
}

type GatewaySvc struct{}

type SvcMiddleware func(GatewayService) GatewayService

func (GatewaySvc) GetAccount() endpoint.Endpoint {
	var (
		tags = []string{"appsvc"}
		passingOnly = true
		getAccount endpoint.Endpoint
		instancer = consulsd.NewInstancer(client, logger, "appsvc", tags, passingOnly)
	)
	{
		
	}
}