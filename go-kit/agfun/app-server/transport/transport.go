package transport

import (
	"github.com/feng/future/go-kit/agfun/app-server/service"
	"net/http"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"context"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/feng/future/go-kit/agfun/app-server/endpoint"
)

//MakeHandler 创建handler
func MakeHandler(svc service.AppService) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	getAccountHandler := kithttp.NewServer(
		endpoint.MakeAccountEndpoint(svc),
		decodeAccountRequest,
		encodeResponse,
		opts...,
	)
	createAccountHandler := kithttp.NewServer(
		endpoint.MakeCreateAccountEndpoint(svc),
		decodeCreateAccountRequest,
		encodeResponse,
		opts...,
	)

	updateAccountHandler := kithttp.NewServer(
		endpoint.MakeUpdateAccountEndpoint(svc),
		decodeUpdateAccountRequest,
		encodeResponse,
		opts...,
	)
	r := mux.NewRouter()

	r.Handle("/appserver/query-account", getAccountHandler).Methods("GET")
	r.Handle("/appserver/create-account", createAccountHandler).Methods("POST")
	r.Handle("/appserver/update-account", updateAccountHandler).Methods("POST")

	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// switch err {
	// 	//todo
	// default:
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	w.WriteHeader(http.StatusInternalServerError)
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