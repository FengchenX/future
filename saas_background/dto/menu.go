//author xinbing
//time 2018/9/11 16:21
//
package dto

import "ibs_service/saas_background/entity"

type MenuDTO struct {
	entity.SaasMenus
	Children []*MenuDTO
}
