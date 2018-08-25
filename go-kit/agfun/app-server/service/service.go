package service

import (
	"fmt"
	"github.com/feng/future/go-kit/agfun/app-server/model"
	"github.com/feng/future/go-kit/agfun/app-server/config"
	"github.com/feng/future/go-kit/agfun/app-server/util"
)

// AppService app服务接口
type AppService interface {
	GetAccount(userAddr string) (uint32, string, model.UserAccount)
	SetAccount(userKeyStore, userParse, keyString string, userAccount model.UserAccount) (uint32, string)
}

//AppSvc app服务实例
type AppSvc struct{}


//SvrMiddleware is a chainable behavior modifier for appservice
type SvcMiddleware func(AppService) AppService

func doPost(url string, reqobj interface{}, respobj interface{}) error {
	url = "http://" + config.AppInst().APIAddr + config.AppInst().APIPort + url
	fmt.Println("URL:>", url)

	return util.DoPost(url, reqobj, respobj)
}
