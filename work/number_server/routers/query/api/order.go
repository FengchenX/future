package api

import (
	"net/http"
	"strconv"
	"sub_account_service/finance/lib"
	"sub_account_service/number_server/models"

	"time"

	"github.com/gin-gonic/gin"
)

//财务服拉取已经成交的支付订单，可根据status来限制商家的拉取，截断
// appId
// version
// shopId
func BatchGetOrderList(c *gin.Context) {
	v := c.Query("version")
	appId := c.Query("appId")
	//shopId:=c.Query("shopId")

	if !models.LegalAppId(appId) {
		c.JSON(http.StatusOK, lib.Result.Fail(80001, "非法的appId"))
		return
	}
	current := models.GetVersionByAppIdAndVersion(appId, v)

	if current.ID == 0 { //该app_id不存在该版本号，那么直接查询数据库中最后的那个版本号

		latest := models.GetLatestVersion(appId)
		beginOrderId := 0
		currentVersionNum := "v1"
		if latest.ID != 0 { //该appId在数据库中有数据
			beginOrderId = latest.EndOrderId
			lastIndex, _ := strconv.Atoi(latest.Version[1:])
			currentVersionNum = "v" + strconv.Itoa(lastIndex+1)
		}
		nextVersionNum := currentVersionNum
		var orders = models.FindOrderLargerThan(beginOrderId,appId)
		if orders != nil && len(orders) > 0 { //有获取到订单，那么插入一条query version，并且nextVersionNum + 1

			current.AppId = appId
			current.BeginOrderId = beginOrderId
			current.EndOrderId = orders[len(orders)-1].ID
			current.Version = currentVersionNum
			current.CreateTime = time.Now()

			if models.IsVersionExist(appId, current.Version) {
				c.JSON(http.StatusOK, lib.Result.Fail(90001, "该version已被使用，请重新获取最新版本号！"))
				return
			}

			if !models.CreateVersion(current) {
				c.JSON(http.StatusOK, lib.Result.Fail(-1, "插入query version失败！"))
			}

			lastIndex, _ := strconv.Atoi(currentVersionNum[1:])
			nextVersionNum = "v" + strconv.Itoa(lastIndex+1)
		}
		c.JSON(http.StatusOK, lib.Result.Success("查询订单成功！", map[string]interface{}{
			"orders":      orders,
			"nextVersion": nextVersionNum,
		}))
	} else { //已经存在该版本
		//获取订单
		orders := models.FindOrderByRange(current.BeginOrderId, current.EndOrderId,appId)

		//获取下一个版本号
		nextVersion := models.GetVersionsByAppIdAndID(current.ID, appId)

		nextVersionNum := nextVersion.Version
		if nextVersion.ID == 0 {
			lastIndex, _ := strconv.Atoi(current.Version[1:])
			nextVersionNum = "v" + strconv.Itoa(lastIndex+1)
		}

		resp := map[string]interface{}{
			"orders":      orders,
			"nextVersion": nextVersionNum,
		}
		c.JSON(http.StatusOK, lib.Result.Success("查询订单成功！", &resp))
	}

}

func GetLatestVersion(c *gin.Context) {

	appId := c.DefaultQuery("appId", "10000")

	if !models.LegalAppId(appId) {
		c.JSON(http.StatusOK, lib.Result.Fail(80001, "非法的appId"))
		return
	}
	latestVersionNum := "v1"

	latestVersion := models.GetLatestVersion(appId)

	if latestVersion.ID != 0 {
		lastIndex, _ := strconv.Atoi(latestVersion.Version[1:])
		latestVersionNum = "v" + strconv.Itoa(lastIndex+1)
	}

	c.JSON(http.StatusOK, lib.Result.Success("获取最新版本号成功！", map[string]interface{}{
		"latestVersion": latestVersionNum,
	}))
}
