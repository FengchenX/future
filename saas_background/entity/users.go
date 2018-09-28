//author xinbing
//time 2018/9/11 10:05
//用户相关
package entity

import "github.com/jinzhu/gorm"

// 用户表
type SaasUsers struct {
	gorm.Model
	Username	string	`gorm:"index:index_username"`
	Pwd 		string
	HeadImg		string
}

// 用户与
type SaasUserCompanyRels struct {
	UserID	    uint	`gorm:"index:index_user_id"`
	CompanyID	uint	`gorm:"index:index_company_id"`
}

// 用户菜单表
type SaasUserMenus struct {
	UserID		uint	`gorm:"index:index_user_id"`
	MenuID		uint	`gorm:"index:index_menu_id"`
}


