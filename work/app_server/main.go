package main

import (
	"flag"
	"github.com/golang/glog"
	"sub_account_service/app_server_v2/api"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/db"
	"sub_account_service/app_server_v2/lib"
	"sub_account_service/app_server_v2/router"

	"time"

	//todo 测试
	//"sub_account_service/app_server_v2/model"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	glog.Infoln("main**********************v4")
	db.InitDb(config.ConfInst().Mysql)
	api.Init()
	go lib.NewKeyMap()

	api.SaveOrders()

	//Test_Adds()
	//Test_InsertSchedule()
	router.Init()

}

func Test_Adds() {
	api.InitReset()
	publisher := "0xf4fe3beb14ac163cb649f69a0d26043f7dae4053"
	var rs []db.Reset
	
	rs = append(rs, db.Reset{
		SubCode: "AC:b4c8bde8d473f171d10c73d89d0800cc",
		UserAddr: "0xf4fe3beb14ac163cb649f69a0d26043f7dae4053",
		LastTime: time.Now().Add(-24*time.Hour),
		Job: "厨师",
		ResetWay: 1,
		ResetTime: 10,
		Publisher: publisher,
	})
	rs = append(rs, db.Reset{
		SubCode: "AC:b4c8bde8d473f171d10c73d89d0800cc",
		UserAddr: "0x8cd010F82E365E1b34925eCfcFFfa170793e7E43",
		LastTime: time.Now().Add(-24*time.Hour),
		Job: "gongsi",
		ResetWay: 1,
		ResetTime: 10,
		Publisher: publisher,
	})
	for i, r := range rs {
		
		if r.ResetWay == 2 {
			// now := time.Now()
			// day := now.Day()
			// dua := time.Duration(day - 1)
			// x := time.Hour * (-1) * 24 * dua
			// rs[i].LastTime = now.Add(x)
		} else if r.ResetWay == 1 {
			now := time.Now()
			h := now.Hour()
			dua := time.Duration(h + 24 - int(r.ResetTime))
			rs[i].LastTime = now.Add(-dua * time.Hour)
		}
	}
	// api.Adds(publisher, rs)
	// c := make(chan int)
	// <- c
}

func Test_InsertSchedule() {
	// var rss []model.Rs
	// rs := model.Rs {
	// 	Accounts: "",
	// 	Level: 1,
	// 	Radios: 100.0,
	// 	SubWay: 0,
	// 	Job: "chushi",
	// }
	// rss = append(rss, rs)
	// rs = model.Rs {
	// 	Accounts: "",
	// 	Level: 1,
	// 	Radios: 100.0,
	// 	SubWay: 0,
	// 	Job: "chushi111",
	// }
	// rss = append(rss, rs)
	// db.InsertSchedule(rss, "AC:111111111111111111111111", "0x22222222222", "hash", "mehton", false)

	var schedule db.Schedule
	db.DbClient.Client.Where("id = 1").Preload("Rs").First(&schedule)
	glog.Infoln(schedule)
}