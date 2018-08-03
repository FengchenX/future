package service

import (
	"sub_account_service/bxg_database_scanner/entity"
	"sub_account_service/bxg_database_scanner/db"
	"sub_account_service/bxg_database_scanner/models"
	"fmt"
	"sub_account_service/bxg_database_scanner/config"
	"github.com/golang/glog"
	"sub_account_service/bxg_database_scanner/utils"
	"strings"
	"time"
)

var queryOrderTemp = `SELECT TOP %v * FROM 
	(
		SELECT o01329,min(b00390) b00390, min(b01348) b01348,min(o00781) o00781,min(b01058) b01058, min(b00569) b00569, max(b01264) b01264, max(b00627) b00627, max(O80112) O80112 FROM bbb077 
		WHERE o01329 > '%v'  AND b01058 IN ('A','B')
		GROUP BY o01329
	)z order by o01329`
var queryOrdersByOrderNosTemp = `SELECT * FROM 
	(
		SELECT o01329,min(b00390) b00390, min(b01348) b01348,min(o00781) o00781,min(b01058) b01058, min(b00569) b00569, max(b01264) b01264, max(b00627) b00627, max(O80112) O80112 FROM bbb077 
		WHERE o01329 IN %v  AND b01058 IN ('A','B')
		GROUP BY o01329
	)z order by o01329`
var queryDetailTemp = `SELECT 
	b.o01329,b.o01292,b.b00193,b.o00557,b.b00741,b.b00517,b.b00495,b.b01058,b.o01068,
	a.a00827,a.b01214 from bbb070 b
	LEFT JOIN aaa050 a ON b.o01292 = a.o01292
	WHERE b.O01329 = '%v'`
// QueryOrderDetail 查询订单详情
// latestOrderNo最后的orderNo
func QueryOrderDetail(latestOrderNo string,pageSize int) ([]*models.BillOrderSaveReq, error){
	bbb077List := make([]*entity.BBB077,0)
	err := db.Engine.SQL(fmt.Sprintf(queryOrderTemp,pageSize,latestOrderNo)).Find(&bbb077List)
	if len(bbb077List) == 0 {
		billOrderSaveList := make([]*models.BillOrderSaveReq,0)
		return billOrderSaveList,err
	}
	return buildOrderSaveReq(bbb077List)
}

func QueryFailedOrders(orderNos *[]string) ([]*models.BillOrderSaveReq, error){
	if len(*orderNos) == 0 {
		billOrderSaveList := make([]*models.BillOrderSaveReq,0)
		return billOrderSaveList,nil
	}
	inOrderStr := "('" + strings.Join(*orderNos,"','") + "')"
	bbb077List := make([]*entity.BBB077,0)
	err := db.Engine.SQL(fmt.Sprintf(queryOrdersByOrderNosTemp,inOrderStr)).Find(&bbb077List)
	if err != nil {
		db.InitDb(config.GetInstance().SqlServerUrl)
		time.Sleep(10 * time.Second)
		err = db.Engine.SQL(fmt.Sprintf(queryOrdersByOrderNosTemp,inOrderStr)).Find(&bbb077List)
		fmt.Println("reconnect the database:",err)
	}
	if len(bbb077List) == 0 {
		billOrderSaveList := make([]*models.BillOrderSaveReq,0)
		return billOrderSaveList,err
	}

	return buildOrderSaveReq(bbb077List)
}

func buildOrderSaveReq(bbb077List []*entity.BBB077) ([]*models.BillOrderSaveReq, error) {
	billOrderSaveList := make([]*models.BillOrderSaveReq,len(bbb077List))
	for index,item := range bbb077List {
		orderDetails := make([]*entity.BBB070,0)
		err := db.Engine.SQL(fmt.Sprintf(queryDetailTemp,item.OrderNo)).Find(&orderDetails)
		if err != nil {
			glog.Errorln("查询错误",err)
		}
		for _,item := range orderDetails {
			convertUtf8(item)
		}
		totalOriginPrice := calculateTotalOriginPrice(&orderDetails)
		billOrderSaveList[index] = &models.BillOrderSaveReq{
			ThirdTradeNo: item.OrderNo,// string  //三方交易号
			OrderNo: item.OrderNo,// 	 string //自己平台的交易单号
			SubNumber: config.GetInstance().SubNumbers[item.ScheduleNo],
			SubNumberTag: item.ScheduleNo,
			OrderTime: utils.FormatDateTime(&item.OrderTime),//    string // 创建时间
			Company: config.GetInstance().Company,//     string // 公司名
			BranchShop: config.GetInstance().BranchShop,//   string // 分店名
			OrderDetails: &orderDetails,
			//bills table
			PayWay: convertPayWay(item.PayWay),//       string  // 支付方式
			Price: item.Price,//        float64 // 订单价格
			Discount: totalOriginPrice - item.Price,//    float64 // 优惠价格
			SumPrice: totalOriginPrice,    //float64 // 总价
			OrderState: getOrderState(item.SellType),	 //int //订单状态
			Salesmen: item.Salesman,
			Remarks: item.Remarks,
		}
	}
	return billOrderSaveList,nil
}
//QueryLatestOrderNo string
func QueryLatestOrderNo() (string,error){
	sql := `SELECT TOP 1 o01329 FROM bbb077 ORDER BY o01329 DESC`
	str := ""
	_,err := db.Engine.SQL(sql).Get(&str)
	return str,err
}

func calculateTotalOriginPrice(list *[]*entity.BBB070) (totalOriginPrice float64){
	for _,item := range *list {
		totalOriginPrice = totalOriginPrice + item.OriginPrice * item.Count
	}
	return
}

func getOrderState(sellType string) int{
	if sellType == "B" || sellType == "b" {
		return 2
	} else {
		return 1
	}
}

func convertPayWay(payWay string) string {
	if payWay == "01" {
		return "支付宝"
	} else if payWay == "02" {
		return "微信"
	} else if payWay == "C" {
		return "信用卡"
	} else if payWay == "RMB" {
		return "现金"
	} else {
		return "其它"
	}
}
func convertUtf8(item *entity.BBB070) {
	bs, err := utils.ConverGBKToUtf8([]byte(item.ProductName))
	if err == nil {
		item.ProductName = string(bs)
	}
}