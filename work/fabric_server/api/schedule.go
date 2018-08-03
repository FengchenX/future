/*--------------------------------------------------------------
* package: 排班相关的服务
* time:    2018/04/17
*-------------------------------------------------------------*/

package api

import (
	"github.com/gin-gonic/gin"
)

// SetSchedule 发布分配表服务
func (api *ApiService) SetSchedule(c *gin.Context) {

}

// 查询排班接口
func (api *ApiService) GetSchedule(c *gin.Context) {

}

// NewScheduleId
func (api *ApiService) NewScheduleId(c *gin.Context) {

}

// 2.5.重设按总量分账
func (api *ApiService) ResetQuo(c *gin.Context) {

}

// 2.6.查询按总量分账
func (api *ApiService) GetQuo(c *gin.Context) {

}

// 2.7.发布排班  <--（客户端）v2新增
func (api *ApiService) SetPaiBan(c *gin.Context) {

}

// 2.8.查询排班  <--（客户端）v2新增
func (api *ApiService) GetPaiBan(c *gin.Context) {

}

// 设置每个账户的已分配定额数
func (api *ApiService) SetSubCodeQuotaData(c *gin.Context) {

}

// 修改财务平台的付款账户地址
func (api *ApiService) ChanngeSmartPayer(c *gin.Context) {

}
