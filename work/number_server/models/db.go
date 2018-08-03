package models

import (
	//导入mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	//导入gorm mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sub_account_service/number_server/config"
	"sync"
	"time"
)

//DbClient 客户端
var DbClient *Db

//AutoMigrate 自动迁移
var AutoMigrate = false

var db *gorm.DB

//Db 数据库
type Db struct {
	addr   string       // the addr of db server
	Lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

//InitDb 初始化数据库
func InitDb(addr string) {
	mydb := &Db{}
	mydb.addr = addr
	var err error
	db, err = gorm.Open("mysql", addr)
	if err != nil {
		glog.Errorln("db initing fail", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		glog.Errorln("db ping fail", err)
		return
	} else {
		glog.Infoln("connecting db success!")
	}
	mydb.Client = db
	DbClient = mydb

	db.DB().SetMaxIdleConns(1000)
	db.DB().SetMaxOpenConns(2000)
	db.LogMode(false)

	AutoMigrate = true

	go timer(addr)
	glog.Infoln("starting db :", addr)
}

func timer(addr string) {
	timer1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer1.C:
			err := DbClient.Client.DB().Ping()
			if err != nil {
				glog.Errorln("mysql connect fail,err:", err)
				InitDb(addr)
			}
		}
	}
}

func Setup() {
	InitDb(config.Opts().Mysql)
}
