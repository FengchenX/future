package api

import (
	"net/http"
	"sub_account_service/fabric_server/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (api *ApiService) SetAccount(ctx *gin.Context) {
	req := model.ReqSetAccount{}
	resp := model.RespSetAccount{}

	ctx.Bind(&req)
	logrus.Info(req)

	if req.UserAccount == nil {
		resp.StatusCode = ACCOUNT_ACCOUNT_NIL
		resp.Msg = "10001 SetAccountFail! req.GetUserAccount() == nil"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	if req.UserAccount.Name == "" {
		logrus.Info("account", "SetAccount error req.GetUserAccount().Name == ")
		resp.StatusCode = ACCOUNT_NAME_NIL
		resp.Msg = "10002 SetAccountFail! req.GetUserAccount() == nil"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	if req.UserAccount.Alipay == "" {
		logrus.Info("10003 SetAccountFail! req.GetUserAccount().Alipay == nil")
		resp.StatusCode = ACCOUNT_APPLY_NIL
		resp.Msg = "10003 SetAccountFail! req.GetUserAccount().Alipay == nil"
		ctx.JSON(http.StatusOK, resp)
		return
	}
	 if api.Fabric==nil{
	 	logrus.Info("err")
	 	return
	 }
	result,err:=api.Fabric.QueryAccount(req)
	if err!=nil{
		logrus.Error("query error")
	}else {
		logrus.Info("query success :"+result)
	}
	resp.StatusCode=0
	resp.Msg=result
	ctx.JSON(http.StatusOK,resp)

}

func (api *ApiService) GetAccount(ctx *gin.Context) {

}
