package db

import (
	"time"

	//"github.com/golang/glog"
)

// 发布信息结构体
type AccountBook struct {
	Id         uint      `gorm:"PRIMARY KEY"`
	SubCode    string    `gorm:"type:text;not null"` // 排班编号
	CreateTime time.Time `gorm:"not null"`           // 创建时间
	Price      float64   `gorm:"not null"`           // 价格
	Details    string    `gorm:"type:text;not null"` // 附加信息
	Account    string    `gorm:"type:text;not null"` // 账户信息
}

// func InsertAccountBook(ab *AccountBook) {
// 	db := DbClient.Client.Create(ab)
// 	if db.Error != nil {
// 		glog.Errorln(db.Error)
// 	}
// }

// func GetAccountBook(where, arg1, arg2 string) []*AccountBook {
// 	var all []*AccountBook
// 	db := DbClient.Client.Model(&AccountBook{}).Where(where, arg1, arg2).Find(&all)
// 	if db.Error != nil {
// 		glog.Errorln(db.Error)
// 	}
// 	return all
// }
