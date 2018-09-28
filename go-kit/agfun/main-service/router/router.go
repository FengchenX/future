
package router

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func init() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(loginCheckMidWare)
	initUserRouter()
}

var router *gin.Engine


func loginCheckMidWare(ctx *gin.Context) {
	// uri := ctx.Request.RequestURI
	// if !loginPass[uri] {
	// 	if store.GetCurrUserId(ctx) <= 0 { //为空
	// 		ctx.String(http.StatusUnauthorized, "未登录")
	// 		ctx.Abort()
	// 		return
	// 	}
	// }
	// ctx.Next()
}

func Start() {
	router.Run(":8080")
}