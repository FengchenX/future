package endpoint

import (
	"github.com/feng/future/go-kit/microsvr/app-server/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"context"
)


func MakeGetAccountEndpoint(svc service.AppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReqGetAccount)
		statusCode, msg, userAccount := svc.GetAccount(req.UserAddress)
		return RespGetAccount{
			StatusCode: statusCode,
			UserAccount: userAccount,
			Msg: msg,
		}, nil
	}
}

//ReqGetAccount 客户端 查询支付账号
type ReqGetAccount struct {
	UserAddress string
}

//RespGetAccount 服务端 查询支付账号
type RespGetAccount struct {
	StatusCode  uint32
	UserAccount model.UserAccount
	Msg         string
}

func MakeSetAccountEndpoint(svc service.AppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReqSetAccount)
		statusCode, msg := svc.SetAccount(req.UserKeyStore, req.UserParse, req.KeyString, req.UserAccount)
		return RespSetAccount{
			StatusCode: statusCode,
			Msg: msg,
		}, nil
	}
}

//ReqSetAccount 客户端 绑定支付账户
type ReqSetAccount struct {
	UserKeyStore string
	UserParse    string
	KeyString    string
	UserAccount  model.UserAccount
}

//RespSetAccount 服务端 绑定支付账户
type RespSetAccount struct {
	StatusCode uint32
	Msg        string
}