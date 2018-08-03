package db

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"time"
	"github.com/go-xorm/core"
)

var DbClient *Db

// ceshi1
type Db struct {
}
var Engine *xorm.Engine

func InitDb(addr string) {
	engine, err := xorm.NewEngine("odbc", addr)
	if err != nil {
		fmt.Println("连接数据库失败!", err)
		return
	}
	engine.SetMapper(core.GonicMapper{})
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(20)
	Engine = engine
	go autoConnect(addr)
}

func autoConnect(addr string) {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <- timer.C:
			err := Engine.Ping()
			if err != nil {
				fmt.Println("mssql connect fail,",err)
				InitDb(addr)
			}
		}
	}
}
