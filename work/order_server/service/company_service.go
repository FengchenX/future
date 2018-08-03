package service

import (
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/entity"
	"github.com/golang/glog"
	"sub_account_service/order_server/utils"
)

var simpleCompanyCache = utils.CreateSimpleCache()
// 根据公司名，分店名查询
func GetCompanyByNameAndBranchName(name,branchName string,fromCache bool) *entity.Company{
	var company *entity.Company
	if fromCache {
		inter := simpleCompanyCache.Get(name+branchName)
		if inter != nil {
			company,ok := inter.(*entity.Company)
			if !ok {
				glog.Errorln("inter cast to company error!")
			}
			return company
		}
	}
	company = &entity.Company{}
	err := db.DbClient.Client.Where("`name` = ? AND branch_name = ?",name,branchName).First(company).Error
	if err != nil {
		glog.Errorln("GetServiceByNameAndBranchName Error:",err)
		return nil
	}
	if company.ID != 0 {
		simpleCompanyCache.Put(name+branchName,company)
		return company
	}
	return nil
}

var appIdDevelopCache = utils.CreateSimpleCache()
func GetDevelopKeyByAppId(appId string) string {
	inter := appIdDevelopCache.Get(appId)
	if inter != nil {
		strPoint,ok := inter.(*string)
		if !ok {
			glog.Errorln("inter cast to string error")
			return ""
		}
		return *strPoint
	}
	company := &entity.Company{}
	db.DbClient.Client.Where("app_id = ?",appId).First(company)
	if company.ID != 0 {
		developKey := company.DevelopKey
		appIdDevelopCache.Put(appId,&developKey)
	}
	return company.DevelopKey
}