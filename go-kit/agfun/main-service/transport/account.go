package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feng/future/go-kit/agfun/app-server/protocol/api"
	"net/http"
	//"github.com/gorilla/mux"
)

func decodeAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.AccountReq
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// vars :=  mux.Vars(r)
	vars := r.URL.Query()
	fmt.Println("decodeAccountRequest", vars)

	request.Account = vars["Account"][0]
	return request, nil
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.CreateAccountReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.UpdateAccountReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
