package main

import (
	"flag"
	//	"os"
	//	"runtime/pprof"

	"github.com/golang/glog"

	bch "sub_account_service/finance/blockchain"
	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/router"
	third "sub_account_service/finance/third_part_pay"
	"sub_account_service/finance/utils"
)

var _BUILD_ = "unknown"
var _VERSION_ = "unknown"

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// 设置同时执行的cup数
	utils.SetGOMAXPROCS()
	/*	fd, err := os.OpenFile("./pprof_log", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0750)
		if err != nil {
			glog.Errorln(err)
		}
		defer func() {
			fd.Close()
		}()
		if err := pprof.StartCPUProfile(fd); err != nil {
			glog.Errorln(err)
		}*/

	flag.Parse()
	//添加查看版本号参数
	glog.Infoln("Build:", _BUILD_)
	glog.Infoln("Version:", _VERSION_)
	defer glog.Flush()

	var path string
	flag.StringVar(&path, "config", "../conf/config.toml", "config path")
	config.ParseToml(path) // initing config
	flag.Parse()

	db.InitDb(config.Opts().MysqlStr)
	db.InitRedis()

	third.InitScheduleArrears()
	third.InitThirdPartPay()
	bch.InitBlockchain()

	glog.Infoln("starting web service")
	router.Init()
}
