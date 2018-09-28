package log

import (
	"fmt"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
)

func (mw logmw) CreateAccount(req api.CreateAccountReq) (resp api.CreateAccountResp, err error) {
	resp, err = mw.AppService.CreateAccount(req)
	fmt.Println("111111111111111", req, resp, err)
	return
}
