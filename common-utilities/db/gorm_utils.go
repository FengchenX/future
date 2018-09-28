//author xinbing
//time 2018/9/4 17:38
package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

// 当只连接一个数据源的时候，可以直接使用GormClient
// 否则应当自己持有管理InitGormDB返回的GormDB
var GormClient *GormDB
type GormDB struct {
	dbConfig *DBConfig
	lock   sync.RWMutex // lock
	Client *gorm.DB     // mysql client
}

// 本方法会给GormClient赋值，多次调用GormClient指向最后一次调用的GormDB
func InitGormDB(dbConfig *DBConfig) *GormDB {
	logrus.Infoln("starting db")
	if err := dbConfig.check(); err != nil {
		logrus.WithError(err).Errorln("error db config!")
		return nil
	}
	myDB := &GormDB{
		dbConfig: dbConfig,
	}
	db, err := gorm.Open("mysql", dbConfig.DBAddr)
	if err != nil {
		logrus.Fatalln("db initing fail", err)
		return nil
	}
	err = db.DB().Ping()
	if err != nil {
		logrus.Fatalln("db ping fail", err)
		return nil
	}
	logrus.WithField("addr",dbConfig.DBAddr ).Infoln("connecting db success!")
	myDB.Client = db
	myDB.initByDBConfigs()
	myDB.initAutoDB()
	go myDB.timer1()
	GormClient = myDB //gormClient
	return myDB
}

//重连接
func (p *GormDB) reConnect() {
	db, err := gorm.Open("mysql", p.dbConfig.DBAddr)
	if err != nil {
		logrus.Fatalln("db reconnect open addr fail", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		logrus.Fatalln("db reconnect ping fail", err)
		return
	}
	p.initByDBConfigs()
	logrus.WithField("db addr",p.dbConfig.DBAddr).Infoln("reconnect db success!")
}

// 初始化参数
func (p *GormDB) initByDBConfigs() {
	p.Client.DB().SetMaxIdleConns(p.dbConfig.MaxIdleConns)
	p.Client.DB().SetMaxOpenConns(p.dbConfig.MaxOpenConns)
	p.Client.LogMode(p.dbConfig.LogMode)
}

//auto create table
func (p *GormDB) initAutoDB() {
	if p.dbConfig.AutoCreateTables != nil && len(p.dbConfig.AutoCreateTables) > 0 {
		logrus.WithField("addr",p.dbConfig.DBAddr).Infoln("begin initAutoDB")
	}
	for _,item := range p.dbConfig.AutoCreateTables {
		p.autoCreate(item)
	}
}

func (p *GormDB) autoCreate(it interface{}) {
	err := p.Client.AutoMigrate(it).Error
	if err != nil {
		logrus.Errorln("auto create ",it," error",err)
	}
}

func (p *GormDB) timer1() {
	timer1 := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-timer1.C:
			err := p.Client.DB().Ping()
			if err != nil {
				logrus.Errorln("mysql connect fail,err:", err)
				logrus.Infoln("reconnect beginning...")
				p.reConnect()
			}
		}
	}
}