package api

import (
	
)

//CreateAccountReq 创建账户请求
type CreateAccountReq struct {
	Account       string
	Password  string
}

//CreateAccountResp 创建账户响应
type CreateAccountResp struct {
	RespBase
}