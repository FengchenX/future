package transport

import (
	"github.com/feng/future/go-kit/microsvr/app_server/service"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"context"
	"encoding/json"
)

//MakeHandler 创建handler
func MakeHandler(svc service.AppService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	getAccountHandler := kithttp.NewServer(
		makeGetAccountEndpoint(svc),
		decodeGetAccountRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("appserver/getaccount", getAccountHandler).Methods("POST")

	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
		//todo
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}