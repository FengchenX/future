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
func (app *AppSvc)Account(account string) (int, string, model.UserAccount) {
	return 0, "", model.UserAccount{}
}
func (app *AppSvc)UpdateAccount(account string, userAccount model.UserAccount) (int, string) {
	return 0, ""
}
func (app *AppSvc)DeleteAccount(account string) (int, string) {
	return 0, ""
}

