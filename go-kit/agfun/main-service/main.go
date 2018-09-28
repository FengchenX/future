package main

import (
	"flag"
	"github.com/feng/future/go-kit/agfun/main-service/service"
	"github.com/feng/future/go-kit/agfun/main-service/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	//kitlogmw "github.com/feng/future/go-kit/agfun/main-service/log"
	"fmt"
	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	//go log.RenameLogFile()

	// Only log the warning severity or above.
	//logrus.SetLevel(logrus.WarnLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	var (
		listen = flag.String("listen", ":8080", "http listen address")
		//proxy =flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	var svc service.AppService
	svc = &service.AppSvc{}
	//svc = kitlogmw.LoggingMiddleware()(svc)

	mux := http.NewServeMux()

	mux.Handle("/appserver/", transport.MakeHandler(svc))

	http.HandleFunc("/check", consulCheck)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 1)
	go func() {
		logrus.Infoln("transport", "http", "address", *listen, "msg", "listening")
		errs <- http.ListenAndServe(*listen, nil)
	}()
	logrus.Infoln("terminated", <-errs)
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

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}
