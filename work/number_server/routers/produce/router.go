package produce

import (
	"sub_account_service/number_server/routers/produce/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(gin.ReleaseMode)

	r.POST("/addOrder", api.AddOrder)

	return r
}
