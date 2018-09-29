package service

import (
	"fmt"
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
	
	return resp.Success("success", createResp), err
}

//Account 获取账户信息
func (app *AppSvc) Account(req api.AccountReq) (api.Resp, error) {
	var resp api.Resp
	var err error
	id := store.GetUserId(req.Accesstoken)
	if id == 0 {
		return resp.Failed("no this user"), err
	}
	myAccount, e := dao.AccountById(id)
	if e != nil {
		return resp.Failed("no this user"), e
	}
	accountResp := api.AccountResp{
		UserAccount: entity.UserAccount {
			Name: myAccount.Name,
			BankCard: myAccount.BankCard,
			WeChat: myAccount.WeChat,
			Telephone: myAccount.Telephone,
		},
	}
	return resp.Success("success", accountResp), nil
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
	fmt.Println("11111111111111111", accessToken)
	return resp.Success("success", loginResp), err
}
