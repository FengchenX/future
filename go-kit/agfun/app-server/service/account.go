package service

import (
	"github.com/feng/future/go-kit/agfun/app-server/model"
	"github.com/feng/future/go-kit/agfun/app-server/dao"
	"github.com/sirupsen/logrus"
)

//CreateAccount 创建账户
func (app *AppSvc)CreateAccount(account, password string) (int, string) {
	var code int
	var msg string
	newAccount := model.UserAccount {
		Account: account,
		Password: password,
	}
	if err := dao.CreateAccount(&newAccount); err != nil {
		logrus.Errorln("CreateAccount newAccount", newAccount)
		code = 11100
		msg = err.Error()
		return code, msg
	}
	code = 0
	msg = "success"
	return code, msg
}

//Account 获取账户信息
func (app *AppSvc)Account(account string) (int, string, model.UserAccount) {
	var code int
	var msg string
	var userAccount model.UserAccount
	if userAccount, err := dao.Account(account); err != nil {
		code = 11100
		msg = err.Error()
		return code, msg, userAccount
	}
	code = 0
	msg = "success"
	return code, msg, userAccount
}

//UpdateAccount 更新账户
func (app *AppSvc)UpdateAccount(account string, userAccount model.UserAccount) (int, string) {
	var code int
	var msg string
	
	return 0, ""
}

//DeleteAccount 删除账户
func (app *AppSvc)DeleteAccount(account string) (int, string) {
	return 0, ""
}

