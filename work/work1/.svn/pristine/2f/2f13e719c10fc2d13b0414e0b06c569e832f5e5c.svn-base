package db

import (
	"time"
)

// 发布信息结构体
type HistoryDeploy struct {
	Id         uint      "gorm:PRIMARY KEY"
	CreateTime time.Time "gorm:not null" // 创建时间
	Address    string    "gorm:not null" // 地址
	Account    string    "gorm:not null" // 账户
	Version    string    "gorm:not null" // 版本
	SubVersion string    "gorm:not null" // 子版本號
	RequsetIp  string    "gorm:not null" // 請求地址
	Menthon    string    "gorm:not null" // 請求方式（WEB_SERVER,MANUAL，SYSTEM）
}

// 发布信息结构体
type RpcHistory struct {
	Id         uint      "gorm:PRIMARY KEY"
	CreateTime time.Time "gorm:not null" // 创建时间
	Address    string    "gorm:not null" // 地址+port
	Version    string    "gorm:not null" // 版本
	SubVersion string    "gorm:not null" // 子版本號
}
