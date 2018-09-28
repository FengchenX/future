//author xinbing
//time 2018/9/11 11:21
//
package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ibs_service/saas_background/saas_config"
	"ibs_service/saas_background/store"
	"net/http"
)

var router *gin.Engine

func init() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(loginCheckMidWare)
	initUserRouter() //user routes
}

func Start() {
	router.Run(fmt.Sprintf(":%v", saas_config.GetConfigInstance().Port))
}

var loginPass = map[string]bool{
	"/users/login": true,
	"/users/logout": true,
}

func loginCheckMidWare(ctx *gin.Context) {
	uri := ctx.Request.RequestURI
	if !loginPass[uri] {
		if store.GetCurrUserId(ctx) <= 0 { //为空
			ctx.String(http.StatusUnauthorized, "未登录")
			ctx.Abort()
			return
		}
	}
	ctx.Next()
}