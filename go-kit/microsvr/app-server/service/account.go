package service

import (
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"github.com/sirupsen/logrus"
	"github.com/feng/future/go-kit/microsvr/app-server/util"
)


//GetAccount 获取账号信息
func (AppSvc) GetAccount(userAddr string) (uint32, string, model.UserAccount) {
	var req struct {
		UserAddress string
	}
	req.UserAddress = userAddr
	var resp struct {
		StatusCode  uint32
		UserAccount model.UserAccount
		Msg         string
	}
	logrus.Infoln("GetAccount*****************req", req)
	if err := doPost("/getaccount", req, &resp); err != nil {
		resp.Msg = err.Error()
		resp.StatusCode = AccountDoPostErr
	}
	logrus.Infoln("GetAccount****************resp", resp)
	return resp.StatusCode, resp.Msg, resp.UserAccount
}

//SetAccount 设置账户
func SetAccount(userKeyStore, userParse, keyString string, userAccount model.UserAccount) (uint32, string) {
	var req struct {
		UserKeyStore string
		UserParse string
		KeyString string
		UserAccount model.UserAccount
	}
	var resp struct {
		StatusCode uint32
		Msg        string
	}
	if keyString == "" {
		keyString = util.KeyMap.Calc(userKeyStore, userParse)
	}	
	err := doPost("/setaccount", req, &resp)
	if err != nil {
		logrus.Errorln("SetAccount**************doPost err", err)
		resp.Msg = err.Error()
		resp.StatusCode = 10001 //todo
	}
	logrus.Infoln("SetAccount****************resp:", resp)
	return resp.StatusCode, resp.Msg
}

