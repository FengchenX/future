package service

import (
	"github.com/feng/future/go-kit/agfun/app-server/dao"
	"github.com/feng/future/go-kit/agfun/app-server/entity"
	"github.com/sirupsen/logrus"
	"github.com/feng/future/go-kit/agfun/app-server/protocol/api"
)

//CreateAccount 创建账户
func (app *AppSvc) CreateAccount(req api.CreateAccountReq) (api.CreateAccountResp, error) {
	var resp api.CreateAccountResp
	var err error
	newAccount := entity.UserAccount{
		Account:  req.Account,
		Password: req.Password,
	}
	if err = dao.CreateAccount(&newAccount); err != nil {
		logrus.Errorln("CreateAccount newAccount", newAccount)
		resp.Code = 11100
		resp.Msg = err.Error()
		return resp, err
	}
	resp.Code = 0
	resp.Msg = "success"
	return resp, err
}

//Account 获取账户信息
func (app *AppSvc) Account(req api.AccountReq) (api.CreateAccountResp, error) {
	var resp api.CreateAccountResp
	var err error
	var userAccount entity.UserAccount
	if userAccount, err = dao.Account(req.Account); err != nil {
		resp.Code = 11100
		resp.Msg = err.Error()
		return resp, err
	}
	resp.Code = 0
	resp.Msg = "success"
	return resp, err
}

//UpdateAccount 更新账户
func (app *AppSvc) UpdateAccount(req api.UpdateAccountReq) (api.UpdateAccountResp, error) {
	var resp api.UpdateAccountResp
	var err error
	userAccount := entity.UserAccount{
		
	}
	if err := dao.UpdateAccount(req.Account, req.); err != nil {
		code = 11100
		msg = err.Error()
		return code, msg
	}
	code = 0
	msg = "success"
	return code, msg
}

//DeleteAccount 删除账户
// func (app *AppSvc) DeleteAccount(account string) (int, string) {
// 	return 0, ""
// }
