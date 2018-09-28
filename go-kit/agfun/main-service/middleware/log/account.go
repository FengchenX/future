package log

import (
	"fmt"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
)

func (mw logmw) CreateAccount(req api.CreateAccountReq) (resp api.Resp, err error) {
	resp, err = mw.AppService.CreateAccount(req)
	fmt.Println("111111111111111", req, resp, err)
	return
}

func (mw logmw) Login(req api.LoginReq) (resp api.Resp, err error) {
	resp, err = mw.AppService.Login(req)
	return
}