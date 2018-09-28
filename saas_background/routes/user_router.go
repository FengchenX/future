//author xinbing
//time 2018/9/11 14:52
//
package routes

import "ibs_service/saas_background/controller"

func initUserRouter() {
	userGroup := router.Group("/users")
	userGroup.POST("/login", controller.Login)
	userGroup.POST("/logout", controller.Logout)
	userGroup.POST("/add", controller.AddUser)
	userGroup.GET("/curr", controller.GetCurrUser)
	userGroup.GET("/menus", controller.GetUserMenus)
	userGroup.GET("/companies", controller.GetUserCompanies)
}
