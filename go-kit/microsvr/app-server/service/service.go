package service

import (
	"fmt"
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"github.com/feng/future/go-kit/microsvr/app-server/config"
	"github.com/feng/future/go-kit/microsvr/app-server/util"
)

// AppService app服务接口
type AppService interface {
	GetAccount(userAddr string) (uint32, string, model.UserAccount)
}

//AppSvr app服务实例
type AppSvr struct{}


//SvrMiddleware is a chainable behavior modifier for appservice
type SvrMiddleware func(AppService) AppService

func doPost(url string, reqobj interface{}, respobj interface{}) error {
	url = "http://" + config.AppInst().APIAddr + config.AppInst().APIPort + url
	fmt.Println("URL:>", url)

	return util.DoPost(url, reqobj, respobj)
}
