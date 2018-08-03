package manage_handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/golang/glog"
	"net/http"
	"sub_account_service/order_server/utils"
	"encoding/json"
	"sub_account_service/order_server/entity"
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/utils/store"
)

//login
func Login(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln("Login err:", err)
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析请求错误！"))
		return
	}
	userDTO := &entity.Users{}
	err = json.Unmarshal(body, userDTO)
	if err != nil {
		glog.Errorln("Login json unmarshal err:", err, string(body))
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析JSON出错！"))
		return
	}
	var dbUser = &entity.Users{}
	db.DbClient.Client.Where("user_name = ?",userDTO.Username).First(dbUser)
	if dbUser.ID == 0 || utils.EncryptUserPwd(userDTO.Pwd) == dbUser.Pwd {
		ctx.JSON(http.StatusOK, utils.Result.Fail("用户名或密码错误！"))
		return
	}
	ctx.SetCookie("ssr", dbUser.Pwd,-1,"/",string(ctx.Request.Host),true,false)
	store.UserCache.Put(dbUser.Pwd,dbUser)
}