//author xinbing
//time 2018/9/11 14:42
//
package dbs

import (
	"common-utilities/db"
	"ibs_service/saas_background/entity"
	"ibs_service/saas_background/saas_config"
)

var SaasGormDB *db.GormDB

var OrderServerGormDB *db.GormDB

func init() {
	initSaasDB()
	initOrderServerDB()
}

func initSaasDB() {
	SaasGormDB = db.InitGormDB(&db.DBConfig{
		DBAddr: saas_config.GetConfigInstance().MySqlAddr,
		LogMode: true,
		AutoCreateTables:[]interface{}{
			&entity.SaasUsers{}, &entity.SaasUserCompanyRels{},
			&entity.SaasMenus{}, &entity.SaasUserMenus{},
		},
	})
}

func initOrderServerDB() {
	OrderServerGormDB = db.InitGormDB(&db.DBConfig{
		DBAddr: saas_config.GetConfigInstance().OrderServerMySqlAddr,
		LogMode: true,
	})
}
