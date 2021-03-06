// Code generated by protoc-gen-go. DO NOT EDIT.
// source: obj.proto

package model

// 用户账户信息结构体
type UserAccount struct {
	Address   string
	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
}

// 账本详情结构体
type AccountBook struct {
	Address         string
	OrderId         string
	Money           float64
	Rflag           bool
	TransferDetails string
	Radio           float64
	Acco            *UserAccount
	SubWay          int64
}

// 订单分账结构体
type OrderDliver struct {
	OrderId string
	Abs     []*AccountBook
}

// 好友信息结构体
type Friend struct {
	Name    string
	Phone   string
	Address string
}

// 发布分配比例结构体
type Sd struct {
	Owner        string
	CreateTime   int64
	Status       uint32
	Rss          []*Rs
	UserAccounts []*UserAccount
	ScheduleId   string
	Message      string
}


// 发布排班结构体
type PaiBan struct {
	UserAddress  string
	JobName      string
}

// 每个人的分配比例结构体
type Rs struct {
	// 参与分账账户
	Accounts string
	// 分账的等级 : 总量分配时优先级较高的优先分到钱，同时比例分配的优先级永远低于总量分配，1级最高优先级,其次为23456
	Level int64
	// 分账的比例 : 比例分配时填写比例，总量分配时填写分配总数
	Radios float64
	// 分配方式 : 0为按比例分配，1为按总量分配
	SubWay int64
	// 重置方式 : 2为每月1号重置。 1为每天重置  0为不重置
	ResetWay int64
	// 重置时间 : 重置时间为整数1~6，代表每天凌晨1点~6点启动更新
	ResetTime int64
	// 已经分配数 (只有定额才有)
	GetMoney float64
	// 参与分账职位
	Job string
}

type Bill struct {
	Name    string
	Money   float64
	Ratio   float64
	SubWay  uint32
	PayAcco string
}