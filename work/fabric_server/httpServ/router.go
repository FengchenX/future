package httpServ

import (
	"net/http"
	"sub_account_service/fabric_server/api"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Handler(f func(c *gin.Context)) func(c *gin.Context) {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusOK, map[string]interface{}{
					"StatusCode": -1,
					"Msg":        "服务器发生错误",
				})
			}
		}()
		f(c)
	}
}

func Serve(api *api.ApiService) {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/", Handler(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world!")
	}))

	router.POST("/getaccount", Handler(api.GetAccount))
	router.POST("/setaccount", Handler(api.SetAccount))

	router.POST("/getaccountbook", Handler(api.GetAccountBook)) // 订单查询服务
	router.POST("/getabbyid", Handler(api.GetABById))           // 根据orderid查询账本
	router.POST("/getabbysh", Handler(api.GetABBySh))           // 根据sheduleid查询账本

	router.POST("/getethbalance", Handler(api.GetEthBalance)) // 获取以太币余额
	router.POST("/getallschedule", Handler(api.GetAllSchedule))
	router.POST("/gettrans", Handler(api.GetTrans))
	router.POST("/getmoney", Handler(api.GetMoney))

	router.POST("/reloadconfig", Handler(api.ReloadConfig))

	router.POST("/setschedule", Handler(api.SetSchedule)) // 发布分配表
	router.POST("/getschedule", Handler(api.GetSchedule)) // 查询分配表

	router.POST("/newscheduleid", Handler(api.NewScheduleId))
	router.POST("/resetquo", Handler(api.ResetQuo)) // 2.5.重设按总量分账
	router.POST("/getquo", Handler(api.GetQuo))     // 2.6.查询按总量分账

	router.POST("/threesetbill", Handler(api.ThreeSetBill)) // 订单上链
	router.POST("/threeconfirm", Handler(api.ThreeConfirm)) // 上链确认

	// 2018.07.14 新增排班表等接口 by GrApple
	router.POST("/setpaiban", Handler(api.SetPaiBan))                   // v3.发布排班  <--（客户端）v2新增
	router.POST("/getpaiban", Handler(api.GetPaiBan))                   // v3.查询排班  <--（客户端）v2新增
	router.POST("/set_subcode_quota", Handler(api.SetSubCodeQuotaData)) // v3.设置每个账户的已分配定额数
	router.POST("/change_samrt_payer", Handler(api.ChanngeSmartPayer))  // v3.修改财务平台的付款账户地址

	gin.SetMode(gin.DebugMode)
	router.Run(":20000")
}
