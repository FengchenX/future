package api

import (
	"github.com/feng/future/go-kit/agfun/main-service/entity"
)

//CreateAccountReq 创建账户请求
type CreateAccountReq struct {
	Account  string
	Password string
}

//CreateAccountResp 创建账户响应
type CreateAccountResp struct {
}

//AccountReq 查询账户请求
type AccountReq struct {
	Accesstoken string
}

//AccountResp 查询账户响应
type AccountResp struct {
	entity.UserAccount
}

//UpdateAccountReq 更新账户请求
type UpdateAccountReq struct {
	Account   string
	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
	Email     string
}

//UpdateAccountResp 更新账户请求响应
type UpdateAccountResp struct {
}

type LoginReq struct {
	UserName string
	Pwd      string
}

type LoginResp struct {
	AccessToken string
}
