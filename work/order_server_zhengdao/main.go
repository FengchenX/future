package main

import (
	"flag"
	"net/http"
	"sub_account_service/order_server_zhengdao/client"
	"sub_account_service/order_server_zhengdao/config"
	"sub_account_service/order_server_zhengdao/db"
	m "sub_account_service/order_server_zhengdao/manager"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Reset handlefunc,change owners function to 'gin style' handle function by using golang Anonymous
// functions,this function have recover.
func Handler(cli *client.Client, f func(c *gin.Context, cli *client.Client)) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorln("panic", err)
				c.JSON(http.StatusInternalServerError, "发生错误")
			}
		}()
		f(c, cli)
	}
}

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	flag.Parse()
	var path string
	flag.StringVar(&path, "config", "../conf/config.toml", "config path")
	glog.Infoln(path)
	config.ParseToml(path) // initing config
	glog.Infoln("starting web service")

	db.InitDb(config.Opts().MysqlStr)

	m.InitAutoTB()

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// starts a new Grpc Client
	cli := client.NewClient(config.Optional.ApiAddress + config.Optional.ApiPort)
	go cli.Run()

	m.InitTransfer()

	// Three_party handlers
	r.POST("/saveorder", Handler(cli, m.SaveOrder)) // Save orders from Three_party api
	r.POST("/savebill", Handler(cli, m.SaveBill))   // Save bills from Three_party api

	r.POST("/getmoney", Handler(cli, m.SaveOrder)) // GetMoney Request

	r.POST("/queryorder", Handler(cli, m.QueryPofile))

	r.GET("/ping", func(c *gin.Context) {
		glog.Infoln("main************Ping")
		c.JSON(http.StatusOK, gin.H{"server": "Pong"})
	})

	// Handle all requests using net/http
	http.Handle("/", r)
	r.Run(config.Optional.LocalAddress + config.Optional.ThreePort)
}
