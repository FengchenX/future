package transport

import (
	"context"
	"net/http"

	"github.com/feng/future/go-kit/agfun/main-service/endpoint"
	"github.com/feng/future/go-kit/agfun/main-service/protocol"
	"github.com/feng/future/go-kit/agfun/main-service/service"
	"github.com/gin-gonic/gin"
)

func decodeAccountRequest(ctx *gin.Context) (interface{}, error) {
	var request protocol.AccountReq
	request.Accesstoken = ctx.Query("AccessToken")
	return request, nil
}

func decodeCreateAccountRequest(ctx *gin.Context) (interface{}, error) {
	var request protocol.CreateAccountReq
	if err := ctx.BindJSON(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateAccountRequest(ctx *gin.Context) (interface{}, error) {

	panic("todo")
}

func decodeLoginRequest(ctx *gin.Context) (interface{}, error) {
	var request protocol.LoginReq
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

func Account(svc service.AppService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, _ := decodeAccountRequest(ctx)
		resp, _ := endpoint.MakeAccountEndpoint(svc)(context.Background(), req)
		ctx.JSON(http.StatusOK, resp)
	}
}
