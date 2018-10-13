package router

import (
	"fmt"

	"github.com/feng/future/go-kit/agfun/agfun-server/controller"
	"github.com/feng/future/go-kit/agfun/agfun-server/service"
)

func initUserRouter(svc service.AppService) {
	userGroup := router.Group("/users")
	fmt.Println(userGroup)
	userGroup.POST("/create", transport.CreateAccount(svc))
	userGroup.POST("/login", transport.Login(svc))
	userGroup.GET("/query", transport.Account(svc))
}
