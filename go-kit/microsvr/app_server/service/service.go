package service

import (
	"github.com/feng/future/go-kit/microsvr/app_server/model"
)
// AppService app服务接口
type AppService interface {
	GetAccount(userAddr string) (uint32, string, model.UserAccount)
}

//AppSvr app服务实例
type AppSvr struct{}

