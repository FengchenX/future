package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
	"net/http"
	"github.com/gin-gonic/gin"
)

func decodeAccountRequest(ctx *gin.Context) (interface{}, error) {
	// var request api.AccountReq
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// vars :=  mux.Vars(r)
	// vars := r.URL.Query()
	// fmt.Println("decodeAccountRequest", vars)

	// request.Account = vars["Account"][0]
	// return request, nil
	panic("todo")
}

func decodeCreateAccountRequest(ctx *gin.Context) (interface{}, error) {
	// var request api.CreateAccountReq
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// return request, nil
	panic("todo")
}

func decodeUpdateAccountRequest(ctx *gin.Context) (interface{}, error) {
	// var request api.UpdateAccountReq
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// return request, nil
	panic("todo")
}
