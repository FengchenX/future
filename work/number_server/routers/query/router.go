package query

import (
	"github.com/gin-gonic/gin"
	"sub_account_service/number_server/routers/query/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())


	r.GET("/orders/batch", api.BatchGetOrderList)

	r.GET("/getLatestVersion", api.GetLatestVersion)

	return r
}
