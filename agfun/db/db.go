package db

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	//引入gorm
	"github.com/feng/future/agfun/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func init() {
	stdAddr := config.Conf().MysqlStr
	stdDB = NewDB(stdAddr)
	go timer(stdAddr)
}

var stdDB *DB

//DB 数据库
type DB struct {
	*gorm.DB        // mysql client
	addr     string // the addr of db server
}

//NewDB addr:数据库地址和密码"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
func NewDB(addr string) *DB {
	glog.Infoln("NewDB****start")
	var temp = &DB{}
	temp.addr = addr
	db, err := gorm.Open("mysql", addr)
	if err != nil {
		glog.Fatalln("NewDB*******Init Fail", err)
	}
	glog.Infoln("NewDB********success")
	temp.DB = db
	return temp
}

//CreateTable 创建表
func (db *DB) CreateTable(models interface{}) {
	db.CreateTable(models)
}

func timer(addr string) {
	timer1 := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer1.C:
			err := stdDB.DB.DB().Ping()
			if err != nil {
				glog.Errorln("mysql connect fail,err:", err)
				stdDB = NewDB(addr)
			}
		}
	}
}
