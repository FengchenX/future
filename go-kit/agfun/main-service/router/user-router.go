package router

import (
	"fmt"

	"github.com/feng/future/go-kit/agfun/main-service/controller"
	"github.com/feng/future/go-kit/agfun/main-service/service"
)

func initUserRouter(svc service.AppService) {
	userGroup := router.Group("/users")
	fmt.Println(userGroup)
	userGroup.POST("/create", transport.CreateAccount(svc))
	userGroup.POST("/login", transport.Login(svc))
	userGroup.GET("/query", transport.Account(svc))
}
