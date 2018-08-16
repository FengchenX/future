package endpoint

import (
	"github.com/feng/future/go-kit/microsvr/app-server/service"
	"github.com/go-kit/kit/endpoint"
	"context"
)

//MakeGetBalanceEndpoint 生成GetBalaceEndpoint端点
func MakeGetEthBalanceEndpoint(svc service.AppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReqGetEthBalance)
		statusCode, msg, balance := svc.GetEthBalance(req.UserAddress)
		return RespGetEthBalance {
			StatusCode: statusCode,
			Msg: msg,
			Balance: balance,
		}, nil
	}
}

//ReqGetEthBalance 客户端 获取以太坊的余额
type ReqGetEthBalance struct {
	UserAddress string
}

//RespGetEthBalance 服务端 获取以太坊的余额
type RespGetEthBalance struct {
	StatusCode uint32
	Balance    string
	Msg        string
}