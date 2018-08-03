package api

import (
	"github.com/golang/glog"
	"sub_account_service/app_server_v2/lib"
	//"sub_account_service/app_server_v2/model"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

//mediator 中介者
func mediator(req, resp interface{}, c *gin.Context, funcName, url string) {
	if err := lib.ParseReq(c, funcName, req); err != nil {
		return
	}
	reqval := reflect.ValueOf(req)
	respval := reflect.ValueOf(resp)
	glog.Infoln(funcName+"**************************req: ", reqval.Elem().Interface())
	if err := doPost(url, reqval.Elem().Interface(), resp); err != nil {
		glog.Errorln(funcName+"*******************doPost err: ", err)
		respval.Elem().FieldByName("Msg").SetString(err.Error())
		c.JSON(http.StatusOK, respval.Elem().Interface())
	}
	glog.Infoln(funcName+"************************resp: ", respval.Elem().Interface())

	c.JSON(http.StatusOK, respval.Elem().Interface())
}

// //GetAccountBook 订单查询服务
// func GetAccountBook(c *gin.Context) {
// 	var req model.ReqGetAccountBook
// 	var resp model.RespGetAccountBook
	
// 	mediator(&req, &resp, c, "GetAccountBook", "/getaccountbook")
// }

// //GetABById 根据orderid查询账本
// func GetABById(c *gin.Context) {
// 	var req model.ReqGetABById
// 	var resp model.RespGetABById
// 	mediator(&req, &resp, c, "GetABById", "/getabbyid")
// }

