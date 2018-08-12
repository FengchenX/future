package transport

import (
	"github.com/feng/future/go-kit/microsvr/app_server/service"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"context"
	"encoding/json"
	"bytes"
	"io/ioutil"
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

	r.Handle("/appserver/getaccount", getAccountHandler).Methods("POST")

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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}