package api

import (
	"sub_account_service/app_server_v2/model"
	"github.com/gin-gonic/gin"
)

//GetEthBalance 获取以太币余额
func GetEthBalance(c *gin.Context) {
	var req model.ReqGetEthBalance
	var resp model.RespGetEthBalance
	mediator(&req, &resp, c, "GetEthBalance", "/getethbalance")
}
