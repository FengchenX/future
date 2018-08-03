package router

import (
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"sub_account_service/app_server_v2/api"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/controllers"
)

func Init() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.POST("/getaccount", api.GetAccount) //10001
	router.POST("/setaccount", api.SetAccount) //20001

	router.POST("/getethbalance", api.GetEthBalance) // 获取以太币余额 //30001
	router.POST("/getallschedule", api.GetAllSchedule) //40001
	router.POST("/getmoney", api.GetMoney) //50001

	router.POST("/setschedule", api.SetSchedule) // 发布分配表服务 //60001
	
	router.POST("/setpaiban", api.SetPaiBan) // 2.7.发布排班  //70001
	router.POST("/getpaiban", api.GetPaiBan) // 2.8.查询排班  //80001

	router.GET("/manager/profile", controllers.Profile)

	gin.SetMode(gin.ReleaseMode)
	port := config.ConfInst().LocalPort
	router.Run(port)
}
