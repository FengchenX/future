package entity

import "time"

type Company struct {
	ID        int
	Name      string    	`gorm:"not null"`
	BranchName string		`gorm:"not null"`
	PassWord  string
	Address   string
	Phone     string
	Describe string    `gorm:"type:text"`
	ThirdTradeNoPrefix string
	AppId			 string	   `gorm:"type:varchar(64);unique;not null"`
	DevelopKey		 string	   `gorm:"not null"`
	DefaultSubNumber string	   `gorm:"not null"`
	CreateTime       time.Time `gorm:"type:datetime;not null"`
	UpdateTime		 time.Time `gorm:"type:datetime;not null"`
}
