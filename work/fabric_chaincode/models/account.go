package models

import "sub_account_service/fabric_chaincode/types"

//账户
type Account struct {
	ModelType   types.ModelType `json:"model_type"`
	AccountAddr string          `json:"account_addr"`
	Name        string          `json:"name"`
	Telephone   string          `json:"telephone"` // 电话号码
	BankCard    string          `json:"bank_card"` // 银行卡，保持只有一个，可更换
	WeChat      string          `json:"we_chat"`   // 微信
	Alipay      string          `json:"alipay"`    // 支付宝
	Balance     int             `json:"balance"`   //余额
}

type Schedule struct {
	ModelType types.ModelType `json:"model_type"`
	IssueCode string `json:"issue_code"`
	Role 	[32]byte `json:"role"`//职位
	Joiner string `json:"joiner"` //账户名
}

type Issue struct {
	ModelType types.ModelType `json:"model_type"`
	SubCode string	`json:"sub_code"`
	Role [32]byte `json:"role"`//职位
	Ratio uint `json:"ratio"` //分配比率
	SubWay uint `json:"sub_way"`// 分配方式，0表示按比例分配，1表示按定额分配。
	QuotaWay uint `json:"quota_way"`// 分配方式，0表示按比例分配，1表示按定额分配。
	ResetTime uint `json:"reset_time"`// 分配数据重置数据，按日的话，每天的这个时间，按月的话，每月1号的这个时间；1-6之间，表1点到6点之间更新
}


//流水
type OrderInfo struct {
	ModelType types.ModelType `json:"model_type"`
	OrderId   string          `json:"order_id"`
	Content   string          `json:"content"`
	CreateAt  string          `json:"create_at"`
	Confirm   bool            `json:"confirm"`
}

//分账
type Settlement struct {
	ModelType       types.ModelType `json:"model_type"`
	JoinAddr        string          `json:"join_addr"`
	Ratio           uint            `json:"ratio"`
	SubWay          uint            `json:"sub_way"`
	Calculate       uint            `json:"calculate"`
	SubCode         string          `json:"sub_code"`
	OrderId         string          `json:"order_id"`
	TotalConsume    uint            `json:"total_consume"`
	Confirm         bool            `json:"confirm"`
	TransferDetails string          `json:"transfer_details"`
}

//定额分配表
type Quota struct {
	ModelType       types.ModelType `json:"model_type"`
	SubCode         string          `json:"sub_code"`
	Money			uint			`json:"money"`
	AccountAddr	string				`json:"account_addr"`
	Role        string 				`json:"role"`
}


