package models

import "time"

type Company struct {
	ID               uint      "gorm:PRIMARY KEY"
	CompanyName      string    `gorm:"not null"`
	CompanyBanchName string    `gorm:"not null"`
	CompanyPassWord  string    `gorm:"not null"`
	CompanyAddress   string    `gorm:"not null"`
	CompanyPhone     string    `gorm:"not null"`
	CompanyDespcrite string    `gorm:"type:text;not null"`
	CreateTime       time.Time `gorm:"not null"`
}
