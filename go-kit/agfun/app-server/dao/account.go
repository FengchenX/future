package dao

import (
	"github.com/feng/future/go-kit/agfun/app-server/model"
	"github.com/sirupsen/logrus"
)

//CreateAccount 创建账户
func CreateAccount(account *model.UserAccount) error {
	if err := createModel(account); err != nil {
		logrus.Errorln("CreateAccount ", err)
		return err
	}
	return nil
}
