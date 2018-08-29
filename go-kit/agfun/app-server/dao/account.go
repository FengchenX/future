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

const accountSQL = "account = ?"
//Account 获取账户
func Account(account string) (model.UserAccount, error) {
	var myAccount model.UserAccount
	if mydb := DBInst().Where(accountSQL, account).First(&myAccount); mydb.Error != nil {
		logrus.Errorln("Account ", mydb.Error)
		return myAccount, mydb.Error
	}
	return myAccount, nil
}
