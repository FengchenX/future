//author xinbing
//time 2018/9/11 16:00
//用户菜单
package controller

import (
	"github.com/gin-gonic/gin"
	"ibs_service/saas_background/service"
	"ibs_service/saas_background/store"
	"net/http"
	"strconv"
)

func GetUserMenus(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	userId, _ := strconv.Atoi(userIdStr)
	currUserId := store.GetCurrUserId(ctx)
	resp := service.GetUserMenus(uint(userId),currUserId == uint(userId))
	ctx.JSON(http.StatusOK, resp)
}
