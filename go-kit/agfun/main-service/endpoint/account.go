package endpoint

import (
	"context"
	"fmt"
	// "github.com/feng/future/go-kit/agfun/main-service/entity"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
	"github.com/feng/future/go-kit/agfun/main-service/service"
	// "github.com/go-kit/kit/endpoint"
)

//MakeAccountEndpoint 生成Account断点
func MakeAccountEndpoint(svc service.AppService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AccountReq)
		var resp api.AccountResp
		fmt.Println(req)
		return resp, nil
	}
}

//MakeCreateAccountEndpoint 生成CreateAccount端点
func MakeCreateAccountEndpoint(svc service.AppService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.CreateAccountReq)
		var resp api.Resp
		resp, _ = svc.CreateAccount(req)

		return resp, nil
	}
}

//MakeUpdateAccountEndpoint 更新账户端点
func MakeUpdateAccountEndpoint(svc service.AppService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.UpdateAccountReq)
		var resp api.Resp
		fmt.Println(req)
		return resp, nil
	}
}

//MakeLoginEndpoint 登录端点
func MakeLoginEndpoint(svc service.AppService) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.LoginReq)
		var resp api.Resp
		resp, _ = svc.Login(req)
		return resp, nil
	}
}
