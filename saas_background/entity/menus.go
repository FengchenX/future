//author xinbing
//time 2018/9/11 10:00
//菜单相关
package entity

import (
	"github.com/jinzhu/gorm"
)

type SaasMenus struct {
	gorm.Model
	Name 		string
	ParentID 	uint	`gorm:"index:index_parent_id"`
	Icon 		string	//图标
	Link		string  //链接
	Orders		int		//排序
}

