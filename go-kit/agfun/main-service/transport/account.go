package transport

import (
	"context"
	"github.com/feng/future/go-kit/agfun/main-service/endpoint"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
	"github.com/feng/future/go-kit/agfun/main-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func decodeAccountRequest(ctx *gin.Context) (interface{}, error) {

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

	panic("todo")
}

func decodeLoginRequest(ctx *gin.Context) (interface{}, error) {
	var request api.LoginReq
	if err := ctx.BindJSON(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//CreateAccount 创建账户
func CreateAccount(svc service.AppService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, _ := decodeCreateAccountRequest(ctx)
		resp, _ := endpoint.MakeCreateAccountEndpoint(svc)(context.Background(), req)
		ctx.JSON(http.StatusOK, resp)
	}
}

//Login 登录
func Login(svc service.AppService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, _ := decodeLoginRequest(ctx)
		resp, _ := endpoint.MakeLoginEndpoint(svc)(context.Background(), req)
		ctx.JSON(http.StatusOK, resp)
	}
}
