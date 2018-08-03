package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
	"sub_account_service/order_server/entity"
	"sub_account_service/order_server/config"
)

var DbClient *Db

type Db struct {
	addr   string       // the addr of db server
	Lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

//auto create table
func InitAutoDB() {
	if config.GetConfig().AutoMigrate {
		glog.Infoln("init AutoMigrate mysql db tables")
		err := DbClient.Client.AutoMigrate(&entity.Company{}).Error
		if err != nil {
			glog.Errorln("AutoMigrate Company error,",err)
		}
		err = DbClient.Client.AutoMigrate(&entity.Order{}).Error
		if err != nil {
			glog.Errorln("AutoMigrate Order error,",err)
		}
		err = DbClient.Client.AutoMigrate(&entity.OrderDetail{}).Error
		if err != nil {
			glog.Errorln("AutoMigrate OrderDetail error",err)
		}
		err = DbClient.Client.AutoMigrate(&entity.Users{}).Error
		if err != nil {
			glog.Errorln("AutoMigrate Users error",err)
		}
		err = DbClient.Client.AutoMigrate(&entity.UserCompanyRel{}).Error
		if err != nil {
			glog.Errorln("AutoMigrate UserCompanyRel error",err)
		}
	}
}

func InitDb(addr string) {
	glog.Infoln("starting db")
	mydb := &Db{}
	mydb.addr = addr
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		glog.Fatalln("db initing fail", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		glog.Fatalln("db ping fail", err)
		return
	} else {
		glog.Infoln("connecting db success!")
	}
	mydb.Client = db
	DbClient = mydb

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)

	go timer1(addr)
}

func (this *Db) CreateTable(models interface{}) {
	this.Client.CreateTable(models)
}

func timer1(addr string) {
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
