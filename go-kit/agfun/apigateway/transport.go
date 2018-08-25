package main

import (
	"github.com/gorilla/mux"
	kitlog "github.com/go-kit/kit/log"
	"net/http"
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"bytes"
	"io/ioutil"
	"github.com/feng/future/go-kit/agfun/app-server/model"
)

//makeHandler 创建handler
func makeHandler(svc GatewayService, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	getAccountHandler := kithttp.NewServer(
		makeGetAccountEndpoint(svc),
		decodeGetAccountRequest,
		encodeJSONResponse,
		opts...,
	)
	r := mux.NewRouter()

	r.Handle("/gateway/v1/getaccount", getAccountHandler).Methods("POST")

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

func encodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}