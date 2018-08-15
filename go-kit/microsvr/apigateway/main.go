package main

import (
	"net/http"
	"flag"	
)

func main() {
	var (
		httpAddr     = flag.String("httpaddr", ":8000", "Address for HTTP (JSON) server")
	)
	flag.Parse()

	var gatewaysvc GatewayService
	gatewaysvc = GatewaySvc{}

	mux := http.NewServeMux()

	mux.Handle("/gateway/", makeHandler(gatewaysvc, logger))

	http.Handle("/", accessControl(mux))
	errs := make(chan error, 1)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
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


