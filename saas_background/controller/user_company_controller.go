//author xinbing
//time 2018/9/11 16:58
//
package controller

import (
	"github.com/gin-gonic/gin"
	"ibs_service/saas_background/service"
	"net/http"
	"strconv"
)

func GetUserCompanies(ctx *gin.Context) {
	userIdStr := ctx.Query("userId")
	userId, _ := strconv.Atoi(userIdStr)
	resp := service.GetCompaniesByUserId(uint(userId))
	ctx.JSON(http.StatusOK, resp)
}
