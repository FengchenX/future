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