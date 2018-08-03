package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"sub_account_service/finance/config"
	"sub_account_service/finance/lib"
	m "sub_account_service/finance/manager"
)

var NOTNEETLOGIN = map[string]bool{ //不需要登陆的接口链接
	"/login":  true,
	"/logout": true,
}

func Init() {
	// Starts a new Gin instance with no middle-ware
	route := gin.New()

	route.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	route.Use(gin.Recovery())

	//	// starts a new Grpc Client
	//	cli := client.NewClient(config.Optional.ApiAddress + config.Optional.ApiPort)
	//	go cli.Run()

	// template page
	route.LoadHTMLGlob("asserts/html/*")
	route.GET("/index", func(c *gin.Context) {
		glog.Infoln("index")
		c.HTML(http.StatusOK, "index_change.html", gin.H{
			"title": "Main website",
		})
	})

	/**
	 * 返回json数据都使用lib.result中的方法
	 */
	// Web handlers

	route.GET("/download/income", Handler(m.DownIncomeExcel))
	route.GET("/download/book", Handler(m.DownBookExcel))
	route.GET("/download/payout", Handler(m.DownExpExcel))

	route.POST("/login", Handler(m.Login))         //用户登录
	route.POST("/querybill", Handler(m.QueryBill)) //查进出帐流水
	route.POST("/getbooks", Handler(m.GetBook))    //获取账本
	route.POST("/pagetransmoney", Handler(m.PageTransMoney))
	route.POST("/queryaccount", Handler(m.AccountList))
	route.POST("/querydist", Handler(m.DisProfile))
	route.POST("/queryrefund", Handler(m.Refund))
	route.GET("/manager/profile", m.Profile)

	// Handle all requests using net/http
	http.Handle("/", route)

	gin.SetMode(gin.ReleaseMode)

	route.Run(config.Optional.LocalAddress + ":" + config.Optional.WebPort)
}

//Handler Reset handlefunc,change owners function to 'gin style' handle function by using golang Anonymous
// functions,this function have recover.
func Handler(f func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusOK, lib.Result.Fail(-2, "服务器发生错误"))
			}
		}()
		if !NOTNEETLOGIN[c.Request.RequestURI] { //需要登陆
			if !strings.Contains(c.Request.RequestURI, "/download") {
				sess := m.GlobalSessions.Session(c)
				if sess == nil {
					glog.Errorln(c.Request.RequestURI + "*******未找到session")
					c.JSON(332, lib.Result.Fail(-1, "请先登录"))
					return
				}
			}
		}
		f(c)
	}
}
