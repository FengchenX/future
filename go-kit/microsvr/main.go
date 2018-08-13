package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"flag"
	"github.com/feng/future/go-kit/microsvr/app-server/service"
	"net/http"
	"github.com/feng/future/go-kit/microsvr/app-server/transport"
	"github.com/go-kit/kit/log"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "http listen address")
		//proxy =flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	var svc service.AppService
	svc = service.AppSvr{}

	mux := http.NewServeMux()
	mux.Handle("/appserver/", transport.MakeHandler(svc, logger))

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 1)
	go func() {
		logger.Log("transport", "http", "address", *listen, "msg", "listening")
		errs <- http.ListenAndServe(*listen, nil)
	}()
	logger.Log("terminated", <- errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}