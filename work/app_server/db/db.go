package db

import (
	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	//gorm 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	
	"sync"
	"time"
	"sub_account_service/app_server_v2/model"
)

//DbClient 数据库客户端
var DbClient *Db

//Db db对象
type Db struct {
	addr   string       // the addr of db server
	Lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

//InitDb 初始化db
func InitDb(addr string) {
	glog.Infoln("starting db", addr)
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
	}
	glog.Infoln("InitDB***************connecting db success!")
	mydb.Client = db
	DbClient = mydb

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(false)
	mydb.AutoCreate()

	go timer1(addr)
}

//CreateTable 创建表
func (DB *Db) CreateTable(models interface{}) {
	DB.Client.CreateTable(models)
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

//AutoCreate 自动创建
func (DB *Db) AutoCreate() {
	glog.Infoln("init AutoMigrate mysql db tables")
	if mydb := DbClient.Client.AutoMigrate(&Schedule{}); mydb.Error != nil {
		glog.Errorln("AutoCreate************schedule:", mydb.Error)
	}
	//DbClient.Client.AutoMigrate(&Reflash{})
	if mydb := DbClient.Client.AutoMigrate(&AppOrder{}); mydb.Error != nil {
		glog.Errorln("AutoCreate*************appOrder:", mydb, mydb.Error)
	}
	if mydb := DbClient.Client.AutoMigrate(&NumberOrder{}); mydb.Error != nil {
		glog.Errorln("AutoCreate*************NumOrder:", mydb.Error)
	}
	if mydb := DbClient.Client.AutoMigrate(&model.Rs{}); mydb.Error != nil {
		glog.Errorln("AutoCreate****************rs:", mydb.Error)
	}
	if mydb := DbClient.Client.AutoMigrate(&Reset{}); mydb.Error != nil {
		glog.Errorln("AutoCreate******************reset:", mydb.Error)
	}
	if mydb := DbClient.Client.AutoMigrate(&model.PaiBan{}); mydb.Error != nil {
		glog.Errorln("AutoCreate*********************paiban:", mydb.Error)
	}
}

//Reset 重置
type Reset struct {
	gorm.Model
	SubCode   string
	UserAddr  string
	LastTime  time.Time
	Job       string
	ResetWay  int64 //0 不重置， 1每天重置, 2每月1号重置
	ResetTime int64
	Publisher string
}


