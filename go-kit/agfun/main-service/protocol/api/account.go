package api

import ()

//CreateAccountReq 创建账户请求
type CreateAccountReq struct {
	Account  string
	Password string
}

//CreateAccountResp 创建账户响应
type CreateAccountResp struct {
	RespBase
}

//AccountReq 查询账户请求
type AccountReq struct {
	Account string
}

//AccountResp 查询账户响应
type AccountResp struct {
	RespBase
	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
	Email     string
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
	RespBase
}
