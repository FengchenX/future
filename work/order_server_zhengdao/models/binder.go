package models

import "time"

type Binder struct {
	ID             uint      "gorm:PRIMARY KEY"
	CompanyId      uint      `gorm:"not null"`
	CompanyName    string    `gorm:"not null"`
	BrenchName     string    `gorm:"not null"`
	UserAddress    string    `gorm:"not null"`
	CompanyAddress string    `gorm:"not null"`
	BindTime       time.Time `gorm:"not null"`
}
