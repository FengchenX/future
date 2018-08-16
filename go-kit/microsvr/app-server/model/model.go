package model

import (
	"github.com/jinzhu/gorm"
	"time"
)


// UserAccount 用户账户信息结构体
type UserAccount struct {
	Address   string
	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
}

//Rs 每个人的分配比例结构体
type Rs struct {
	//ID uint `gorm:"primary_key"`
	gorm.Model
	ScheID uint `gorm:"not null"` 
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

//Schedule 发布信息结构体
type Schedule struct {
	gorm.Model
	SubCode    string         `gorm:"type:text;not null"`                // 排班编号
	Publisher  string         `gorm:"not null"`                          // 发布者
	CreateTime time.Time      `gorm:"not null default CURRENT_TIESTAMP"` // 创建时间
	Status     int64          `gorm:"not null"`                          // 是否当前
	Menthon    string         `gorm:"not null"`                          // 備註信息
	Hash       string         `gorm:"not null"`                          // Hash
	Times      int64          `gorm:"not null"`                          // Hash
	Rs         []Rs     `gorm:"ForeignKey:ScheID;AssociationForeignKey:ID"`
	PaiBan     []PaiBan `gorm:"ForeignKey:ScheID;AssociationForeignKey:ID"`
}

//PaiBan 发布排班结构体
type PaiBan struct {
	gorm.Model 
	ScheID uint `gorm:"not null"`
	UserAddress string
	JobName     string
}