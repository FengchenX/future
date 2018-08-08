package service

import (
	"github.com/feng/future/go-kit/microsvr/app_server/model"
)


//GetAccount 获取账号信息
func (AppSvr) GetAccount(userAddr string) (uint32, string, model.UserAccount) {
	
	return 1, "", model.UserAccount{}
}