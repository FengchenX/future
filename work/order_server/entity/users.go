package entity

import "github.com/jinzhu/gorm"

//用户表
type Users struct {
	gorm.Model
	Username string
	Pwd		 string
	NickName string
}

//用户关系表
type UserCompanyRel struct {
	UserId   int `gorm:"primary key"`
	CompanyId int `gorm:"primary key"`
}
