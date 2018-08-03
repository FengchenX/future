package manager

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sub_account_service/order_server_zhengdao/client"
	"sub_account_service/order_server_zhengdao/db"
	"sub_account_service/order_server_zhengdao/lib"
	"sub_account_service/order_server_zhengdao/models"
)

//QueryPofile 查询详情
func QueryPofile(c *gin.Context, cli *client.Client) {
	var qp struct {
		CurrentPage int    `json:"currentPage"`
		PageSize    int    `json:"pageSize"`
		OrderID     int    `json:"orderID"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
	}

	if err := parseReq(c, "query_profile", &qp); err != nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}
	//glog.Infoln(qp)

	var resp respQP
	var querystr string
	var params []interface{}

	if qp.OrderID != -1 {
		querystr += " AND order_saves.order_id = ?"
		params = append(params, qp.OrderID)
	}
	if qp.StartDate != "" {
		qp.StartDate += " 00:00:00"
		querystr += " AND order_saves.create_time > ?"
		params = append(params, qp.StartDate)
	}
	if qp.EndDate != "" {
		qp.EndDate += " 23:59:59"
		querystr += " AND order_saves.create_time < ?"
		params = append(params, qp.EndDate)
	}

	if db.DbClient.Client == nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "db.DbClient.Client==nil"))
		glog.Errorln("QueryPofile*******************************db == nil")
		return
	}
	db := db.DbClient.Client

	resp.Pagination.PageSize = qp.PageSize
	resp.Pagination.CurrentPage = qp.CurrentPage
	re := regexp.MustCompile(`(?i:^and)`)
	querystr = re.ReplaceAllString(strings.Trim(querystr, " "), "")
	var Models []ordersvrBill //todo 换成对应模型
	//mydb := db.Where(querystr, params...).Order("create_time desc").Limit(qp.PageSize).Offset((qp.CurrentPage-1)*qp.PageSize).Find(&Models)

	mydb := db.Table("order_saves").Joins("join bills on order_saves.order_id = bills.order_id").
		Select("order_saves.order_id, order_saves.id, order_saves.create_time, order_saves.company,"+
			"order_saves.branch_shop, order_saves.order_content, order_saves.price, bills.trade_status, bills.transfer_info").
		Where(querystr, params...).Order("create_time desc").Limit(qp.PageSize).Offset((qp.CurrentPage - 1) * qp.PageSize).Find(&Models)
	if mydb.Error != nil {
		glog.Errorln("QueryPofile**********************", mydb.Error)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}

	var total int
	mydb = db.Table("order_saves").Joins("join bills on order_saves.order_id = bills.order_id").
		Select("order_saves.order_id, order_saves.id, order_saves.create_time, order_saves.company,"+
			"order_saves.branch_shop, order_saves.order_content, order_saves.price, bills.trade_status, bills.transfer_info").
		Where(querystr, params...).Count(&total)
	if mydb.Error != nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "查询失败"))
		return
	}

	resp.Pagination.Total = total

	var Orders []order
	for _, one := range Models {
		var order order
		order.OrderId = one.OrderId
		order.CreateTime = one.CreateTime
		order.Company = one.Company
		order.BranchShop = one.BranchShop
		oc := make([]models.Dish, 0)
		if err := json.Unmarshal([]byte(strings.Trim(one.OrderContent, " ")), &oc); err != nil {
			glog.Errorln("json*************************解码错误", err)
			return
		}
		order.OrderContent = oc
		order.Price = one.Price
		order.ID = one.ID
		order.TradeStatus = one.TradeStatus
		order.TransferInfo = one.TransferInfo
		Orders = append(Orders, order)
	}
	resp.List = Orders
	glog.Infoln("QueryPofile**********************", resp)

	c.JSON(http.StatusOK, resp)
}

type ordersvrBill struct {
	models.OrderSave
	TradeStatus  int64
	TransferInfo string
}

type order struct {
	ID           uint
	OrderId      int64  // 订单id(订单的唯一id)
	CreateTime   string // 创建时间
	Company      string // 公司名
	BranchShop   string // 分店名
	OrderContent []models.Dish
	Price        float64 // 订单价格
	TradeStatus  int64
	TransferInfo string
}

type respQP struct {
	List       interface{} `json:"list"`
	Pagination pagination  `json:"pagination"`
}

type pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
}

func parseReq(c *gin.Context, funcName string, qb interface{}) error {
	glog.Infoln(funcName + "**************************start")
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln(funcName+"*************************readll 错误", err)
		return err
	}

	err = json.Unmarshal(buf, qb)
	if err != nil {
		glog.Errorln(funcName+"**************************json 解码错误", err)
		return err
	}
	return nil
}
