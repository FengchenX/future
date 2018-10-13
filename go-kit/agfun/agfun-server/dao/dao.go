package dao

import (
	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	//gorm 驱动
	"github.com/feng/future/go-kit/agfun/agfun-server/config"
	"github.com/feng/future/go-kit/agfun/agfun-server/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

func init() {
	DBInst()
	autoMigrate()
}

//DB 数据库对象
var DB *gorm.DB

//DBInst 数据库单例
func DBInst() *gorm.DB {
	if DB == nil {
		logrus.Infoln("starting db", config.AppInst().MysqlAddr)
		temp, err := gorm.Open("mysql", config.AppInst().MysqlAddr)
		if err != nil {
			logrus.Errorln("db initing fail******", err)
			return nil
		}
		err = temp.DB().Ping()
		if err != nil {
			logrus.Errorln("db initing fail******", err)
			return nil
		}
		logrus.Infoln("InitDB**************connect db success!")
		temp.DB().SetMaxIdleConns(10)
		temp.DB().SetMaxOpenConns(100)
		temp.LogMode(false)
		DB = temp
	}
	return DB
}

func autoMigrate() {
	if mydb := DBInst().AutoMigrate(&entity.UserAccount{}); mydb.Error != nil {
		logrus.Error("autoMigrate UserAccount", mydb.Error)
	}
}

func createModel(value interface{}) error {
	if DBInst().NewRecord(value) {
		if mydb := DBInst().Create(value); mydb.Error != nil {
			return mydb.Error
		}
	}
	return nil
}
