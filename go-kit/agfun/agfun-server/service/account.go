package service

import (
	"github.com/feng/future/go-kit/agfun/agfun-server/dao"
	"github.com/feng/future/go-kit/agfun/agfun-server/entity"

	// "github.com/sirupsen/logrus"
	"common-utilities/encrypt"
	"common-utilities/utilities"

	"github.com/feng/future/go-kit/agfun/agfun-server/protocol"
	"github.com/feng/future/go-kit/agfun/agfun-server/store"
)

//CreateAccount 创建账户
func (app *AppSvc) CreateAccount(req protocol.CreateAccountReq) (protocol.Resp, error) {
	var resp protocol.Resp
	var err error
	userAccount := entity.UserAccount{
		Account: req.Account,
		Pwd:     req.Pwd,
	}
	if err = dao.CreateAccount(&userAccount); err != nil {
		panic(err)
	}
	createResp := protocol.CreateAccountResp{}

	return resp.Success("success", createResp), err
}

//Account 获取账户信息
func (app *AppSvc) Account(req protocol.AccountReq) (protocol.Resp, error) {
	var resp protocol.Resp
	var err error
	id := store.GetUserId(req.Accesstoken)
	if id == 0 {
		return resp.Failed("no this user"), err
	}
	myAccount, e := dao.AccountById(id)
	if e != nil {
		return resp.Failed("no this user"), e
	}
	accountResp := protocol.AccountResp{
		UserAccount: entity.UserAccount{
			Name:      myAccount.Name,
			BankCard:  myAccount.BankCard,
			WeChat:    myAccount.WeChat,
			Telephone: myAccount.Telephone,
		},
	}
	return resp.Success("success", accountResp), nil
}

//UpdateAccount 更新账户
func (app *AppSvc) UpdateAccount(req protocol.UpdateAccountReq) (protocol.Resp, error) {
	// var resp protocol.UpdateAccountResp
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

//Login 登录
func (app *AppSvc) Login(req protocol.LoginReq) (protocol.Resp, error) {
	var resp protocol.Resp
	var err error
	var loginResp protocol.LoginResp

	myAccount, e := dao.Account(req.Account)
	if e != nil {
		panic(err)
	}
	if myAccount.ID == 0 || req.Pwd != myAccount.Pwd {
		return resp.Failed("用户名或密码错误"), err
	}
	accessToken := encrypt.SHA1(utilities.GetRandomStr(32) + req.Pwd)
	store.CacheUser(accessToken, myAccount.ID)
	loginResp.AccessToken = accessToken
	return resp.Success("success", loginResp), err
}
