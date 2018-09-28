package service

import (
	"github.com/feng/future/go-kit/agfun/main-service/protocol/api"
)

// AppService app服务接口
type AppService interface {
	CreateAccount(req api.CreateAccountReq) (api.CreateAccountResp, error)
	// Account(req api.AccountReq) (api.CreateAccountResp, error)
	// UpdateAccount(req api.UpdateAccountReq) (api.UpdateAccountResp, error)
	// DeleteAccount(req api.) (int, string)
	// Login()
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
