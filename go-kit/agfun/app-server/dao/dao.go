package dao

import (
	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	//gorm 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/feng/future/go-kit/agfun/app-server/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//DB 数据库对象
var DB *gorm.DB

//DBInst 数据库单例
func DBInst() *gorm.DB {
	if DB == nil {
		logrus.Infoln("starting db", config.AppInst().MysqlAddr)
		DB, err := gorm.Open("mysql", config.AppInst().MysqlAddr)
		if err != nil {
			logrus.Errorln("db initing fail******", err)
			return nil
		}
		err = DB.DB().Ping()
		if err != nil {
			logrus.Errorln("db initing fail******", err)
			return nil
		}
		logrus.Infoln("InitDB**************connect db success!")
		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(100)
		DB.LogMode(false)
		autoMigrate()		
	}
	return DB
}

func autoMigrate() {

}

