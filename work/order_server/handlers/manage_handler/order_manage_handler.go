package manage_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/entity"
	"github.com/golang/glog"
	"regexp"
	"strings"
	"net/http"
	"sub_account_service/order_server/utils"
	"fmt"
)
//type queryOrderModel struct {
//	ThirdTradeNo string `form:"thirdTradeNo" json:"thirdTradeNo"`
//	OrderNo string `form:"orderNo" json:"orderNo"`
//	BeginDate string `form:"beginDate" json:"beginDate"`
//	EndDate string `form:"endDate" json:"endDate"`
//	CurrentPage int `form:"currentPage" json:"currentPage"`
//	PageSize string `form:"pageSize" json:"pageSize"`
//}
//QueryOrders query orders
func QueryOrders(ctx *gin.Context) {
	thirdNo,orderNo,beginDateStr,endDateStr := ctx.Query("thirdTradeNo"), ctx.Query("orderNo"), ctx.Query("beginDate"), ctx.Query("endDate")
	billStateStr := ctx.DefaultQuery("billState","-1")
	companyId,_ := strconv.Atoi(ctx.DefaultQuery("companyId","-1"))
	currentPageStr, pageSizeStr := ctx.DefaultQuery("currentPage","1"), ctx.DefaultQuery("pageSize","10")
	currentPage, _ := strconv.Atoi(currentPageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	billState, _ := strconv.Atoi(billStateStr)
	whereSql := ""
	whereParams := make([]interface{},0)
	if companyId != -1 {
		whereSql += " AND company_id = ?"
		whereParams = append(whereParams, companyId)
	}
	if thirdNo != "" {
		whereSql += " AND third_trade_no LIKE ?"
		whereParams = append(whereParams, "%"+thirdNo+"%")
	}
	if orderNo != "" {
		whereSql += " AND order_no LIKE ?"
		whereParams = append(whereParams, "%"+orderNo+"%")
	}
	if billState > 0 {
		whereSql += " AND bill_state = ?"
		whereParams = append(whereParams, billState)
	}
	if beginDateStr != "" {
		whereSql += " AND DateDiff(order_time,?) >= 0"
		whereParams = append(whereParams, beginDateStr)
	}
	if endDateStr != "" {
		whereSql += " AND DateDiff(order_time,?) <= 0"
		whereParams = append(whereParams, endDateStr)
	}
	if len(whereParams) > 0 {
		re := regexp.MustCompile(`(?i:^and)`)
		whereSql = re.ReplaceAllString(strings.Trim(whereSql, " "), "")
	}
	fmt.Println(whereSql,whereParams)
	pageDb := db.DbClient.Client.Table("orders").Select("orders.*").Where(whereSql,whereParams...)
	totalCount := 0
	if err := pageDb.Count(&totalCount).Error; err != nil {
		glog.Error("QueryOrders queryCount err",err)
		ctx.JSON(http.StatusOK, utils.Result.Fail("查询数量失败！"))
		return
	}
	pagination := &Pagination{
		CurrentPage:currentPage,
		PageSize:pageSize,
	}
	pagination.Total = totalCount
	offset := (currentPage - 1) * pageSize
	if totalCount < offset {
		ctx.JSON(http.StatusOK, utils.Result.Success("查询成功", &PageResult{
			List: make([]interface{},0),
			Pagination:pagination,
		}))
		return
	}
	var orders []*entity.Order
	if err := pageDb.Order("order_time DESC",).Limit(pageSize).Offset(offset).Find(&orders).Error; err != nil {
		glog.Errorln("QueryOrders error:", err)
		ctx.JSON(http.StatusOK, utils.Result.Fail("查询失败!"))
		return
	}
	ctx.JSON(http.StatusOK, utils.Result.Success("查询成功！", &PageResult{
		List: orders,
		Pagination: pagination,
	}))
}

//GetOrderDetails by Order id
func GetOrderDetails(ctx *gin.Context) {
	orderId,err := strconv.Atoi(ctx.Query("orderId"))
	if err != nil {
		ctx.JSON(http.StatusOK, utils.Result.Fail("订单ID不正确！"))
		return
	}
	var orderDetails []*entity.OrderDetail
	err = db.DbClient.Client.Where("order_id = ?", orderId).Find(&orderDetails).Error
	if err != nil {
		glog.Error("orderId：",orderId,"GetOrderDetails Failed",err)
		ctx.JSON(http.StatusOK,utils.Result.Fail("获取订单失败！"))
		return
	}
	ctx.JSON(http.StatusOK, utils.Result.Success("查找成功！",orderDetails))
}
type PageResult struct {
	List interface{} `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize int `json:"pageSize"`
	Total int `json:"total"`
}