package manager

import (
	"github.com/golang/glog"
	"sub_account_service/finance/db"
)

func InitAutoTB() {
	if db.AutoMigrate == true {
		glog.Infoln("init AutoMigrate mysql db tables")
	}
}
