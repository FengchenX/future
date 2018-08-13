package service

import (
	"github.com/feng/future/go-kit/microsvr/app-server/model"
)


//GetAccount 获取账号信息
func (AppSvr) GetAccount(userAddr string) (uint32, string, model.UserAccount) {
	var req struct {
		UserAddress string
	}
	req.UserAddress = userAddr
	var resp struct {
		StatusCode  uint32
		UserAccount model.UserAccount
		Msg         string
	}
	if err := doPost("/getaccount", req, &resp); err != nil {
		resp.Msg = err.Error()
		resp.StatusCode = AccountDoPostErr
	}
	return resp.StatusCode, resp.Msg, resp.UserAccount
}

