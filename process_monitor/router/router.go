package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sub_account_service/process_monitor/monitor"
	"fmt"
	"sub_account_service/process_monitor/config"
)

func Start() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,monitor.Cache)
	})
	http.Handle("/",router)
	router.Run(fmt.Sprintf(":%v",config.GetConfigInstance().LocalPort))
}