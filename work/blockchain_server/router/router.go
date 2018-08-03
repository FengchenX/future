package router

import (
	"net/http"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"fmt"
	API "sub_account_service/blockchain_server/api"
	"sub_account_service/blockchain_server/config"
)

func Handler(f func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusOK, map[string]interface{}{
					"StatusCode":-1,
					"Msg":"服务器发生错误",
				})
			}
		}()
		f(c)
	}
}

func Init() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/", Handler(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world!")
	}))

	router.POST("/getaccount", Handler(API.GetAccount))
	router.POST("/setaccount", Handler(API.SetAccount))

	router.POST("/getaccountbook", Handler(API.GetAccountBook)) // 财务平台getAccountBook
	router.POST("/getabbyid", Handler(API.GetABById))           // 根据orderid查询账本
	router.POST("/getabbysh", Handler(API.GetABBySh))           // 根据sheduleid查询账本

	router.POST("/getethbalance", Handler(API.GetEthBalance)) // 获取以太币余额
	router.POST("/getallschedule", Handler(API.GetAllSchedule))
	router.POST("/gettrans", Handler(API.GetTrans))
	router.POST("/getmoney", Handler(API.GetMoney))

	router.POST("/reloadconfig", Handler(API.ReloadConfig))

	router.POST("/setschedule", Handler(API.SetSchedule)) // 发布分配表
	router.POST("/getschedule", Handler(API.GetSchedule)) // 查询分配表

	router.POST("/newscheduleid", Handler(API.NewScheduleId))
	router.POST("/resetquo", Handler(API.ResetQuo)) // 2.5.重设按总量分账
	router.POST("/getquo", Handler(API.GetQuo))     // 2.6.查询按总量分账

	router.POST("/threesetbill", Handler(API.ThreeSetBill)) // 订单上链
	router.POST("/threeconfirm", Handler(API.ThreeConfirm)) // 上链确认

	// 2018.07.14 新增排班表等接口 by GrApple
	router.POST("/setpaiban", Handler(API.SetPaiBan))                   // v3.发布排班  <--（客户端）v2新增
	router.POST("/getpaiban", Handler(API.GetPaiBan))                   // v3.查询排班  <--（客户端）v2新增
	router.POST("/set_subcode_quota", Handler(API.SetSubCodeQuotaData)) // v3.设置每个账户的已分配定额数
	router.POST("/change_samrt_payer", Handler(API.ChanngeSmartPayer)) // v3.修改财务平台的付款账户地址

	gin.SetMode(gin.ReleaseMode)
	port := fmt.Sprintf(":%v", config.Optional.Port)
	router.Run(port)

	/*
		s := &http.Server{
			Addr:           port,
			Handler:        router, // < here Gin is attached to the HTTP server
			//  ReadTimeout:    10 * time.Second,
			//  WriteTimeout:   10 * time.Second,
			//  MaxHeaderBytes: 1 << 20,
		}
		s.SetKeepAlivesEnabled(true)
		s.ListenAndServe()
	*/
}
