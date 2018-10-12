package service

import (
	"github.com/feng/future/go-kit/agfun/main-service/protocol"
)

// AppService app服务接口
type AppService interface {
	CreateAccount(req protocol.CreateAccountReq) (protocol.Resp, error)
	Account(req protocol.AccountReq) (protocol.Resp, error)
	// UpdateAccount(req protocol.UpdateAccountReq) (protocol.UpdateAccountResp, error)
	// DeleteAccount(req protocol.) (int, string)
	Login(req protocol.LoginReq) (protocol.Resp, error)
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
