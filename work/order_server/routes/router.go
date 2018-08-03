package routes

import (
	"fmt"
	"net/http"
	"sub_account_service/order_server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"sub_account_service/order_server/config"
	"sub_account_service/order_server/handlers/manage_handler"
	"sub_account_service/order_server/service"
	"sub_account_service/order_server/utils"
	"sub_account_service/order_server/utils/store"
)
// Reset handlefunc,change owners function to 'gin style' handle function by using golang Anonymous
// functions,this function have recover.
func webSecurityWrapper(f func(c *gin.Context),needLogin bool) func(c *gin.Context){
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusInternalServerError, "inter error!")
			}
		}()
		if needLogin {
			ssr,_ := c.GetCookie("ssr")
			if ssr == "" || store.UserCache.Get(ssr) == nil { //为空
				c.JSON(332, utils.Result.Fail("未登录"))
				return
			}
		}
		f(c)
	}
}

func appSecurityWrapper(f func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusInternalServerError, "inter error!")
			}
		}()
		appId := c.Query("AppId")
		developKey := service.GetDevelopKeyByAppId(appId)
		fmt.Println(appId)
		timestamp := c.Query("Timestamp")
		sign := c.Query("Sign")
		if utils.MD5(timestamp+developKey) != sign {
			c.JSON(http.StatusOK, utils.Result.New(10000, "invalid sign", nil))
			return
		}
		f(c)
	}
}
var router *gin.Engine
func init() {
	router = gin.New()
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//app handlers
	router.POST("orders/save", appSecurityWrapper(handlers.SaveOrder))
	router.POST("orders/getLatestOrderNo",appSecurityWrapper(handlers.GetLatestOrderNo))
	//web handlers
	router.GET("orders", webSecurityWrapper(manage_handler.QueryOrders, true))
	router.GET("orders/details",webSecurityWrapper(manage_handler.GetOrderDetails, true))
	router.POST("login",webSecurityWrapper(manage_handler.Login, false))
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "wait!what are you doing?I'm os！")
	//})

	router.Static("/html", "./assets/html")
	http.Handle("/", router)
}

func Run() {
	router.Run(fmt.Sprintf(":%v",config.GetConfig().Port))
}
