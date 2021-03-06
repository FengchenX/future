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
	SetAccount(userKeyStore, userParse, keyString string, userAccount model.UserAccount) (uint32, string)
	GetEthBalance(userAddr string) (uint32, string, string)
	SetSchedule(userAddress, 
		userKeyStore, 
		userParse, 
		keyString, 
		scheduleName string, 
		Rss []model.Rs, 
		message string) (uint32, string, string, string)
	
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
