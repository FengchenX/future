package main

import (
	_ "github.com/lunny/godbc"
	"sub_account_service/bxg_database_scanner/config"
	"sub_account_service/bxg_database_scanner/schedules"
	"sub_account_service/bxg_database_scanner/db"
)

func main() {
	db.InitDb(config.GetInstance().SqlServerUrl) //初始化db连接
	schedules.Start() //启动定时任务
}