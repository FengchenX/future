package transport

import (
	"github.com/feng/future/go-kit/microsvr/app_server/service"
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/feng/future/go-kit/microsvr/app_server/model"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

func makeGetAccountEndpoint(svc service.AppService) endpoint.Endpoint {
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

func decodeGetAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ReqGetAccount
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
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