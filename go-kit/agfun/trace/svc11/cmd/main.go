package main

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/feng/future/go-kit/agfun/trace/svc11"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"os"
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v/n", err)
		os.Exit(-1)
	}
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v/n", err)
		os.Exit(-1)
	}

	opentracing.InitGlobalTracer(tracer)
	service := svc11.NewService()
	svc11.RouterInit(hostPort, tracer, service)
}

const (
	serviceName = "svc11"
	// Host + port of our service.
	hostPort = "127.0.0.1:61001"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// Base endpoint of our SVC2 service.
	svc2Endpoint = "http://localhost:61002"

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)
