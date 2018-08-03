package main

import (
	"sub_account_service/order_server/config"
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/schedules"
	"sub_account_service/order_server/routes"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	//init db
	db.InitDb(config.GetConfig().MysqlStr)
	db.InitAutoDB() //auto create table
	//end init db
	//init routers
	schedules.StartAddOrderSchedule() //开启定时任务
	routes.Run()                      //启动
}
