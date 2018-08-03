package db

import (
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"sub_account_service/finance/models"
)

//DbClient 客户端
var DbClient *Db

//AutoMigrate 自动迁移
var AutoMigrate = false

//Db 数据库
type Db struct {
	addr   string       // the addr of db server
	Lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

//InitDb 初始化数据库
func InitDb(addr string) {
	glog.Infoln("starting db :", addr)
	mydb := &Db{}
	mydb.addr = addr
	db, err := gorm.Open("mysql", addr)
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

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)

	AutoMigrate = true
	//创建表
	createTable()

	go timer1(addr)
}

//CreateTable 创建表
func (this *Db) CreateTable(models interface{}) {
	DbClient.Client.AutoMigrate(models)
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

func createTable() {
	DbClient.Client.AutoMigrate(&models.User{})
	DbClient.Client.AutoMigrate(&models.ExpensesBill{})
	DbClient.Client.AutoMigrate(&models.IncomeStatement{})
	DbClient.Client.AutoMigrate(&models.UserBill{})
	DbClient.Client.AutoMigrate(&models.UserAccount{})
	DbClient.Client.AutoMigrate(&models.RefundTrade{})
	DbClient.Client.AutoMigrate(&models.ScheduleAccount{})
}
