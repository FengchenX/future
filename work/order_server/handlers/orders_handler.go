package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sub_account_service/order_server/config"
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/dto"
	"sub_account_service/order_server/entity"
	"sub_account_service/order_server/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"sub_account_service/order_server/service"
	"fmt"
)

//保存订单
func SaveOrder(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln("SaveOrder err:", err)
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析请求错误！"))
		return
	}
	fmt.Println(string(body))
	orderSaveDTO := &dto.BillOrderSaveDTO{}
	err = json.Unmarshal(body, orderSaveDTO)
	if err != nil {
		glog.Errorln("SaveOrder json unmarshal err:", err, string(body))
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析JSON出错！"))
		return
	}
	if orderSaveDTO.OrderDetails == nil || len(*orderSaveDTO.OrderDetails) == 0 {
		glog.Errorln("don't have orderDetails")
		ctx.JSON(http.StatusOK, utils.Result.Fail("没有传输orderDetails"))
		return
	}

	glog.Infoln(utils.Loger("[three] SaveOrder", "order_msg"),orderSaveDTO.ThirdTradeNo, orderSaveDTO.Company)
	tx := db.DbClient.Client.Begin()
	code, msg := 0, ""
	defer func() {
		if err := recover(); err != nil || code != 0 { //有错误，或者code不等于0
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	code, msg = addOrder(tx, orderSaveDTO)
	ctx.JSON(http.StatusOK, utils.Result.New(code, msg, nil))
}

//添加订单
func addOrder(tx *gorm.DB, orderSaveDTO *dto.BillOrderSaveDTO) (code int, msg string) {
	company := service.GetCompanyByNameAndBranchName(orderSaveDTO.Company,orderSaveDTO.BranchShop,true)
	if company == nil {
		return -1, "不存在对应的公司！"
	}
	if company.ThirdTradeNoPrefix != "" {
		orderSaveDTO.ThirdTradeNo = company.ThirdTradeNoPrefix + orderSaveDTO.ThirdTradeNo
	}
	count := 0
	tx.Model(&entity.Order{}).Where(" third_trade_no = ?",
		orderSaveDTO.ThirdTradeNo).Count(&count)
	if count > 0 {
		return 10001, "订单已存在，不允许重复添加"
	}
	orderTime, err := utils.ParseDateTimeStr(orderSaveDTO.OrderTime)
	if err != nil {
		return -1, "订单时间格式不正确"
	}

	if orderSaveDTO.SubNumber == "" { //没有传分账编号
		orderSaveDTO.SubNumber = company.DefaultSubNumber
	}
	if orderSaveDTO.SubNumber == "" {
		return -1, "排班编号不允许为空！"
	}
	cont, _ := json.Marshal(orderSaveDTO.OrderDetails)
	sonStr := bytes.NewBuffer(cont).String()
	order := &entity.Order{
		ThirdTradeNo: orderSaveDTO.ThirdTradeNo,
		OrderNo:      orderSaveDTO.OrderNo,
		SubNumber:	  orderSaveDTO.SubNumber,
		AppId:        company.AppId,
		CompanyId:    company.ID,
		CompanyName:  company.Name,
		BranchName:   company.BranchName,
		OrderTime:    orderTime,
		PayWay:       orderSaveDTO.PayWay,
		Discount:     orderSaveDTO.Discount,
		SumPrice:     orderSaveDTO.SumPrice,
		Price:        orderSaveDTO.Price,
		BillState:    orderSaveDTO.OrderState,
		Complete:     0,
		OrderContent: sonStr,
		Salesmen:     orderSaveDTO.Salesmen,
		Remarks:      orderSaveDTO.Remarks,
	}

	err = tx.Create(order).Error
	if err != nil {
		glog.Errorln("SaveOrder create db err:", err, orderSaveDTO)
		return -1, "保存订单出错！"
	}
	orderDetails :=  orderSaveDTO.OrderDetails
	for _,item := range *orderDetails {
		item.OrderId = order.ID
		err = tx.Create(item).Error
		if err != nil {
			glog.Errorln("save order detail failed,err:", err, "thirdTradeNo:",order.ThirdTradeNo)
			return -1,"保存订单明细失败"
		}
	}
	go SendToNumberServer(order) //往编号服务添加
	return 0, "添加订单成功！"
}

// GetLatestOrderNo
func GetLatestOrderNo(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln("GetLatestOrderNo err:", err)
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析请求错误！"))
		return
	}
	orderSaveDTO := &dto.BillOrderSaveDTO{}
	err = json.Unmarshal(body, orderSaveDTO)
	if err != nil {
		glog.Errorln("GetLatestOrderNo json unmarshal err:", err, string(body))
		ctx.JSON(http.StatusOK, utils.Result.Fail("解析JSON出错！"))
		return
	}
	company := service.GetCompanyByNameAndBranchName(orderSaveDTO.Company,orderSaveDTO.BranchShop,true)
	if company == nil {
		ctx.JSON(http.StatusOK, utils.Result.Fail("不存在对应的公司"))
		return
	}
	var order = &entity.Order{}
	err = db.DbClient.Client.Table("orders").Select("orders.order_no").Where("company_id = ?",company.ID).Order("order_no DESC").First(order).Error
	if err != nil {
		glog.Error("get latest order no err:",err,",companyId:",company.ID)
		ctx.JSON(http.StatusOK,utils.Result.Fail("查询最后的订单编号失败！"))
		return
	}
	ctx.JSON(http.StatusOK, utils.Result.Success("查询成功！",order.OrderNo))
}

func SendToNumberServer(order *entity.Order) {
	orderSvc := dto.OrderServer{
		SubAccountNo:   order.SubNumber,
		ThirdTradeNo:   order.ThirdTradeNo,
		OrderNo:        order.OrderNo,
		MchId:			order.AppId,
		Company:        order.CompanyName,
		BranchShop:     order.BranchName,
		OrderType:      0,
		PaymentType:    0,
		AutoTransfer:   0,
		OrderTime:      order.OrderTime,
		TransferAmount: order.Price,     // 转账金额
		OrderState:     order.BillState, //订单状态
	}

	out, err := json.Marshal(orderSvc)
	if err != nil {
		glog.Errorln("parse order err:", orderSvc)
		return
	}
	_, err = utils.Post(config.GetOrderNoConfig().AddOrder, bytes.NewReader(out), nil, nil)
	if err != nil {
		glog.Errorln("SendToNumberServer err", err)
		return
	}
	db.DbClient.Client.Model(order).Update("complete", 1)
}
