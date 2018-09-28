package service

import (
	"github.com/feng/future/go-kit/agfun/main-service/dao"
	"github.com/feng/future/go-kit/agfun/main-service/entity"
	// "github.com/sirupsen/logrus"
	"common-utilities/encrypt"
	"common-utilities/utilities"
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
	"github.com/feng/future/go-kit/agfun/main-service/store"
)

//CreateAccount 创建账户
func (app *AppSvc) CreateAccount(req api.CreateAccountReq) (api.Resp, error) {
	var resp api.Resp
	var err error
	userAccount := entity.UserAccount{
		Account:  req.Account,
		Password: req.Password,
	}
	if err = dao.CreateAccount(&userAccount); err != nil {
		panic(err)
	}
	createResp := api.CreateAccountResp{}
	resp.Success("success", createResp)
	return resp, err
}

//Account 获取账户信息
func (app *AppSvc) Account(req api.AccountReq) (api.Resp, error) {
	// var resp api.CreateAccountResp
	// var err error
	// var userAccount entity.UserAccount
	// if userAccount, err = dao.Account(req.Account); err != nil {
	// 	resp.Code = 11100
	// 	resp.Msg = err.Error()
	// 	return resp, err
	// }
	// resp.Code = 0
	// resp.Msg = "success"
	// return resp, err
	panic("todo")
}

//UpdateAccount 更新账户
func (app *AppSvc) UpdateAccount(req api.UpdateAccountReq) (api.Resp, error) {
	// var resp api.UpdateAccountResp
	// var err error
	// userAccount := entity.UserAccount{

	// }
	// if err := dao.UpdateAccount(req.Account, req.); err != nil {
	// 	code = 11100
	// 	msg = err.Error()
	// 	return code, msg
	// }
	// code = 0
	// msg = "success"
	// return code, msg
	panic("todo")
}

//DeleteAccount 删除账户
// func (app *AppSvc) DeleteAccount(account string) (int, string) {
// 	return 0, ""
// }

func (app *AppSvc) Login(req api.LoginReq) (api.Resp, error) {
	var resp api.Resp
	var err error
	var loginResp api.LoginResp

	myAccount, e := dao.Account(req.UserName)
	if e != nil {
		panic(err)
	}
	if myAccount.ID == 0 || req.Pwd != myAccount.Password {
		return resp.Failed("用户名或密码错误"), err
	}
	accessToken := encrypt.SHA1(utilities.GetRandomStr(32) + req.Pwd)
	store.CacheUser(accessToken, myAccount.ID)
	loginResp.AccessToken = accessToken
	return resp.Success("success", loginResp), err
}
