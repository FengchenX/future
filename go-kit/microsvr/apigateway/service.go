package main

import (
	"fmt"
	"net/url"
	"strings"
	"io"
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
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	
	"github.com/go-kit/kit/sd/lb"
)
var (
	// httpAddr     = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
	// consulAddr   = flag.String("consul.addr", "", "Consul agent address")
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
		tags = []string{"appserver"}
		passingOnly = true
		getAccount endpoint.Endpoint
		instancer = consulsd.NewInstancer(client, logger, "appserver", tags, passingOnly)
	)
	{
		factory := appsvcFactory(ctx, "GET", "/appserver/getaccount")
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(*retryMax, *retryTimeout, balancer)
		getAccount = retry
	}
	return getAccount
}

func appsvcFactory(ctx context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path
		var (
			enc httptransport.EncodeRequestFunc
			dec httptransport.DecodeResponseFunc
		)
		switch path {
		case "/appserver/getaccount":
			enc, dec = encodeJSONRequest, decodeGetAccountResponse
		default:
			return nil, nil, fmt.Errorf("unknown appsvc path %q", path)
		}
		return httptransport.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}
