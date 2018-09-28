//author xinbing
//time 2018/9/11 16:08
//
package service

import (
	"github.com/sirupsen/logrus"
	"ibs_service/saas_background/dbs"
	"ibs_service/saas_background/dto"
	"ibs_service/saas_background/utils"
)

func GetUserMenus(userId uint, currUser bool) *utils.Resp {
	var menuDTOSlice []*dto.MenuDTO
	err := dbs.SaasGormDB.Client.Table("saas_menus").Joins("JOIN saas_user_menus um ON um.menu_id = saas_menus.id").
		Where("um.user_id = ? AND saas_menus.parent_id = ?", userId, 0).Order("orders ASC").Scan(&menuDTOSlice).Error
	if err != nil {
		logrus.WithError(err).Errorln("GetUserMenus failed!")
		return utils.Resp{}.Failed("GetUserMenus failed!")
	}
	for _,item := range menuDTOSlice {
		recFindChildren(item, userId)
	}
	return utils.Resp{}.Success("获取成功！", &menuDTOSlice)
}

// 直接递归查询
func recFindChildren(menuDTO *dto.MenuDTO, userId uint) {
	var menuDTOSlice []*dto.MenuDTO
	err := dbs.SaasGormDB.Client.Table("saas_menus").Joins("JOIN saas_user_menus um ON um.menu_id = saas_menus.id").
		Where("parent_id = ? AND um.user_id = ?", menuDTO.ID, userId).Order("orders ASC").Scan(&menuDTOSlice).Error
	if err != nil {
		logrus.WithError(err).Errorln("query failed!")
		return
	}
	menuDTO.Children = menuDTOSlice
	for _, item := range menuDTOSlice {
		recFindChildren(item, userId)
	}
}
