//author xinbing
//time 2018/9/11 13:52
// 用户相关
package controller

import (
	"github.com/gin-gonic/gin"
	"ibs_service/saas_background/entity"
	"ibs_service/saas_background/service"
	"ibs_service/saas_background/store"
	"ibs_service/saas_background/utils"
	"net/http"
)

//登录
func Login(ctx *gin.Context) {
	var userDTO entity.SaasUsers
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.String(http.StatusBadRequest, "bad request!")
		return
	}
	resp,token := service.Login(userDTO.Username, userDTO.Pwd)
	if resp.Code == 0 {
		ctx.SetCookie("access-token", token ,0,"/", ctx.Request.Host, false, false)
	}
	ctx.JSON(http.StatusOK, resp)
}

// 登出
func Logout(ctx *gin.Context) {
	store.RemoveUser(ctx)
	ctx.SetCookie("access-token", "", -1, "/", ctx.Request.Host, false, false)
	ctx.JSON(http.StatusOK, utils.Resp{}.Success("退出登录成功！",nil))
}

// 获取当前用户
func GetCurrUser(ctx *gin.Context) {
	currUserId := store.GetCurrUserId(ctx)
	resp := service.GetUserInfo(currUserId)
	ctx.JSON(http.StatusOK, resp)
}

func AddUser(ctx *gin.Context) {
	var userDTO entity.SaasUsers
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.String(http.StatusBadRequest, "bad request!")
		return
	}
	resp := service.AddUser(userDTO)
	ctx.JSON(http.StatusOK, resp)
}

