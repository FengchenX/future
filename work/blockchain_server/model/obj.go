// Code generated by protoc-gen-go. DO NOT EDIT.
// source: obj.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户账户信息结构体
type UserAccount struct {
	Address   string `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	BankCard  string `protobuf:"bytes,3,opt,name=BankCard" json:"BankCard,omitempty"`
	WeChat    string `protobuf:"bytes,4,opt,name=WeChat" json:"WeChat,omitempty"`
	Alipay    string `protobuf:"bytes,5,opt,name=Alipay" json:"Alipay,omitempty"`
	Telephone string `protobuf:"bytes,6,opt,name=Telephone" json:"Telephone,omitempty"`
}

// 账本详情结构体
type AccountBook struct {
	Address         string       `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	OrderId         string       `protobuf:"bytes,2,opt,name=OrderId" json:"OrderId,omitempty"`
	Money           float64      `protobuf:"fixed64,3,opt,name=Money" json:"Money,omitempty"`
	Rflag           bool         `protobuf:"varint,4,opt,name=Rflag" json:"Rflag,omitempty"`
	TransferDetails string       `protobuf:"bytes,5,opt,name=TransferDetails" json:"TransferDetails,omitempty"`
	Radio           float64      `protobuf:"fixed64,6,opt,name=Radio" json:"Radio,omitempty"`
	Acco            *UserAccount `protobuf:"bytes,7,opt,name=Acco" json:"Acco,omitempty"`
	SubWay          int64        `protobuf:"varint,8,opt,name=SubWay" json:"SubWay,omitempty"`
}

// 订单分账结构体
type OrderDliver struct {
	OrderId string         `protobuf:"bytes,1,opt,name=OrderId" json:"OrderId,omitempty"`
	Abs     []*AccountBook `protobuf:"bytes,2,rep,name=Abs" json:"Abs,omitempty"`
}

// 好友信息结构体
type Friend struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=Phone" json:"Phone,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
}

// 发布分配比例结构体
type Sd struct {
	Owner        string         `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	CreateTime   int64          `protobuf:"varint,2,opt,name=CreateTime" json:"CreateTime,omitempty"`
	Status       uint32         `protobuf:"varint,3,opt,name=Status" json:"Status,omitempty"`
	Rss          []*Rs          `protobuf:"bytes,4,rep,name=Rss" json:"Rss,omitempty"`
	UserAccounts []*UserAccount `protobuf:"bytes,5,rep,name=UserAccounts" json:"UserAccounts,omitempty"`
	ScheduleId   string         `protobuf:"bytes,6,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	Message      string         `protobuf:"bytes,7,opt,name=Message" json:"Message,omitempty"`
}


// 发布排班结构体
type PaiBan struct {
	UserAddress  string         `protobuf:"bytes,4,opt,name=UserAddress" json:"UserAddress,omitempty"`
	JobName      string         `protobuf:"bytes,5,opt,name=JobName" json:"JobName,omitempty"`
}

// 每个人的分配比例结构体
type Rs struct {
	// 参与分账账户
	Accounts string `protobuf:"bytes,1,opt,name=Accounts" json:"Accounts,omitempty"`
	// 分账的等级 : 总量分配时优先级较高的优先分到钱，同时比例分配的优先级永远低于总量分配，1级最高优先级,其次为23456
	Level int64 `protobuf:"varint,2,opt,name=Level" json:"Level,omitempty"`
	// 分账的比例 : 比例分配时填写比例，总量分配时填写分配总数
	Radios float64 `protobuf:"fixed64,3,opt,name=Radios" json:"Radios,omitempty"`
	// 分配方式 : 0为按比例分配，1为按总量分配
	SubWay int64 `protobuf:"varint,4,opt,name=SubWay" json:"SubWay,omitempty"`
	// 重置方式 : 2为每月1号重置。 1为每天重置  0为不重置
	ResetWay int64 `protobuf:"varint,5,opt,name=ResetWay" json:"ResetWay,omitempty"`
	// 重置时间 : 重置时间为整数1~6，代表每天凌晨1点~6点启动更新
	ResetTime int64 `protobuf:"varint,6,opt,name=ResetTime" json:"ResetTime,omitempty"`
	// 已经分配数 (只有定额才有)
	GetMoney float64 `protobuf:"fixed64,7,opt,name=GetMoney" json:"GetMoney,omitempty"`
	// 参与分账职位
	Job string `protobuf:"bytes,8,opt,name=Job" json:"Job,omitempty"`
}

type Bill struct {
	Name    string  `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Money   float64 `protobuf:"fixed64,2,opt,name=Money" json:"Money,omitempty"`
	Ratio   float64 `protobuf:"fixed64,3,opt,name=Ratio" json:"Ratio,omitempty"`
	SubWay  uint32  `protobuf:"varint,4,opt,name=SubWay" json:"SubWay,omitempty"`
	PayAcco string  `protobuf:"bytes,5,opt,name=PayAcco" json:"PayAcco,omitempty"`
}