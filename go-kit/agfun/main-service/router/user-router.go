package router

import (
	"fmt"
)

func initUserRouter() {
	userGroup := router.Group("/users")
	fmt.Println(userGroup)
	//userGroup.POST("/login", )
}