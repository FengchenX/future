package service

import (
	"github.com/feng/future/go-kit/agfun/app-server/model"
)

// AppService app服务接口
type AppService interface {
	CreateAccount(account, password string) (int, string)
	Account(account string) (int, string, model.UserAccount)
	UpdateAccount(account string, userAccount model.UserAccount) (int, string)
	DeleteAccount(account string) (int, string)
	Login()
}

//AppSvc app服务实例
type AppSvc struct{}

//SvcMiddleware is a chainable behavior modifier for appservice
type SvcMiddleware func(AppService) AppService

// func doPost(url string, reqobj interface{}, respobj interface{}) error {
// 	url = "http://" + config.AppInst().APIAddr + config.AppInst().APIPort + url
// 	fmt.Println("URL:>", url)

// 	return util.DoPost(url, reqobj, respobj)
// }
