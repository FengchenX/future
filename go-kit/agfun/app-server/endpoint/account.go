package endpoint

import (
	"github.com/feng/future/go-kit/agfun/app-server/service"
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/feng/future/go-kit/agfun/app-server/protocol/api"
)


//MakeAccountEndpoint 生成Account断点
func MakeAccountEndpoint(svc service.AppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AccountReq)
		var resp api.AccountResp
		statusCode, msg, userAccount := svc.Account(req.Account)
		resp.Code = statusCode
		resp.Msg = msg
		resp.Name = userAccount.Name
		resp.BankCard = userAccount.BankCard
		resp.WeChat = userAccount.WeChat
		resp.Alipay = userAccount.Alipay
		resp.Telephone = userAccount.Telephone
		resp.Email = userAccount.Email
		return resp, nil
	}
}

//MakeCreateAccountEndpoint 生成CreateAccount端点
func MakeCreateAccountEndpoint(svc service.AppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.CreateAccountReq)
		statusCode, msg := svc.CreateAccount(req.Account, req.Password)
		var resp api.CreateAccountResp
		resp.Code = statusCode
		resp.Msg = msg
		return resp, nil
	}
}

