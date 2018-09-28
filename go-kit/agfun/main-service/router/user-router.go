package router

import (
	"fmt"
	"github.com/feng/future/go-kit/agfun/main-service/transport"
	"github.com/feng/future/go-kit/agfun/main-service/service"
)

func initUserRouter(svc service.AppService) {
	userGroup := router.Group("/users")
	fmt.Println(userGroup)
	userGroup.POST("/create", transport.CreateAccount(svc))
}
