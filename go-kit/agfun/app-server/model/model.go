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
