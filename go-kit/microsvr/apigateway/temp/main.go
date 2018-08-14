package main

import (
	"syscall"
	"os/signal"
	"net/url"
	"strings"
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
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"fmt"
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
		{
			factory := addsvcFactory(addendpoint.MakeConcatEndpoint, tracer, zipkinTracer, logger)
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			endpoints.ConcatEndpoint = retry
		}
		r.PathPrefix("/addsvc").Handler(http.StripPrefix("/addsvc", addtransport.NewHTTPHandler(endpoints, tracer, zipkinTracer, logger)))
	}

	{
		var (
			tags = []string{"appsvc"}
			passingOnly = true
			getAccount endpoint.Endpoint

			instancer = consulsd.NewInstancer(client, logger, "appsvc", tags, passingOnly)
		)
		{
			factory := appsvcFactory(ctx, "GET", "/getaccount") 
			endpointer := sd.NewEndpointer(instancer, factory, logger)
			balancer := lb.NewRoundRobin(endpointer)
			retry := lb.Retry(*retryMax, *retryTimeout, balancer)
			getAccount = retry
		}

		r.Handle("/appsvc/getaccount", httptransport.NewServer(getAccount, decodeGetAccountRequest, encodeJSONResponse))
	}

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c) 
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, r)
	}()

	logger.Log("exit", <-errc)
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
		case "/":
			enc, dec = encodeJSONRequest, decodeGetAccountResponse
		default:
			return nil, nil, fmt.Errorf("unknown stringsvc path %q", path)
		}
		return httptransport.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}


func encodeJSONRequest(_ context.Context, req *http.Request, request interface{}) error {
	// Both uppercase and count requests are encoded in the same way:
	// simple JSON serialization to the request body.
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeGetAccountResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response struct {
		StatusCode  uint32
		UserAccount model.UserAccount
		Msg         string
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeGetAccountRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request struct {
		UserAddress string
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}