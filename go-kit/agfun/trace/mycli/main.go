package main

import (
	"context"
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"

	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/feng/future/go-kit/agfun/trace/svc11"
)

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}
	opentracing.InitGlobalTracer(tracer)
	client := svc11.NewHTTPClient(tracer, svc1Endpoint)

	span := opentracing.StartSpan("Run")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	span.LogEvent("Call Concat")
	res1, err := client.Concat(ctx, "Hello", "World!")
	fmt.Printf("Concat: %s Err: %+v\n", res1, err)
	
	span.Finish()
	collector.Close()
}

const (
	// Our service name.
	serviceName = "myCli"

	// Host + port of our service.
	hostPort = "0.0.0.0:0"

	// Endpoint to send Zipkin spans to.
	zipkinHTTPEndpoint = "http://localhost:9411/api/v1/spans"

	// Debug mode.
	debug = false

	// Base endpoint of our SVC1 service.
	svc1Endpoint = "http://localhost:61001"

	// same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = true

	// make Tracer generate 128 bit traceID's for root spans.
	traceID128Bit = true
)