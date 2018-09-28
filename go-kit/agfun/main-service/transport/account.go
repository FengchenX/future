package transport

import (
	"context"
	// "fmt"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/feng/future/go-kit/agfun/main-service/endpoint"
	"github.com/feng/future/go-kit/agfun/main-service/service"
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
	var request api.CreateAccountReq
	if err := ctx.BindJSON(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateAccountRequest(ctx *gin.Context) (interface{}, error) {
	// var request api.UpdateAccountReq
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	return nil, err
	// }
	// return request, nil
	panic("todo")
}


func CreateAccount(svc service.AppService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, _ := decodeCreateAccountRequest(ctx)
		resp, _ := endpoint.MakeCreateAccountEndpoint(svc)(context.Background(), req)
		ctx.JSON(http.StatusOK, resp)
	}
}