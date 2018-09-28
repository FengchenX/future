package entity

import (
	"github.com/jinzhu/gorm"
	//"time"
)

// UserAccount 用户账户信息结构体
type UserAccount struct {
	gorm.Model

	Account  string
	Password string

	PermanentID string //用户所有操作使用此ID

	Name      string
	BankCard  string
	WeChat    string
	Alipay    string
	Telephone string
	Email     string
}
