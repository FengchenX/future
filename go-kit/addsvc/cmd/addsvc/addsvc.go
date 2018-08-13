package main

import (
	"flag"
	"os"
	"github.com/go-kit/kit/log"
	"fmt"
	"text/tabwriter"
	
)

func main() {
	fs := flag.NewFlagSet("addsvc", flag.ExitOnError)
	var (
		debugAddr = fs.String("debug.addr", ":8080", "Debug and metrics listen address")
		httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
		grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
		thriftAddr     = fs.String("thrift-addr", ":8083", "Thrift listen address")
		jsonRPCAddr    = fs.String("jsonrpc-addr", ":8084", "JSON RPC listen address")
		thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
		thriftBuffer   = fs.Int("thrift-buffer", 0, "0 for unbuffered")
		thriftFramed   = fs.Bool("thrift-framed", false, "true to enable framing")
		zipkinV2URL    = fs.String("zipkin-url", "", "Enable Zipkin v2 tracing (zipkin-go) using a Reporter URL e.g. http://localhost:9411/api/v2/spans")
		zipkinV1URL    = fs.String("zipkin-v1-url", "", "Enable Zipkin v1 tracing (zipkin-go-opentracing) using a collector URL e.g. http://localhost:9411/api/v1/spans")
		lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
		appdashAddr    = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var tracer 
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}