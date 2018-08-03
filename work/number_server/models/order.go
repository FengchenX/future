package models

import   (
	"time"
	"sub_account_service/number_server/pkg/log"
)

type Orders struct {
	ID             int       `json:"id"`
	ThirdTradeNo   string    `json:"third_trade_no"`  //第三方交易号
	OrderNo        string    `json:"order_no"`        //商户订单编号
	SubAccountNo   string    `json:"sub_account_no"`  //分账编号
	MchId	   	   string		     `json:"mch_id"`
	Company        string    `json:"company"`         //所属公司,
	BranchShop	   string	 `json:"branch_shop"`
	OrderType      uint8     `json:"order_type"`      //订单类型，0:支付宝，1:微信
	PaymentType    uint8     `json:"payment_type"`    //付款类型， 0:支付，1:转账
	TransferAmount float64   `json:"transfer_amount"` // 转账金额
	TransferInfo   string    `json:"transfer_info"`   // 转账信息
	AutoTransfer   uint8     `json:"auto_transfer"`   //是否自动打款 0否，1是，默认是1
	OrderTime      time.Time `json:"order_time"`
	OrderState     uint8     `json:"order_state"` //订单状态
	CreateTime     time.Time `json:"create_time"` //创建时间
	UpdateTime     time.Time `json:"update_time"` //更新时间a
}



type OrderServer struct {
	ThirdTradeNo string //第三方交易号
	OrderNo string //商户订单编号
	SubAccountNo string //分账编号
	MchId    string
	BranchShop   string
	Company string //所属公司,
	OrderType uint8 //订单类型，0:支付宝，1:微信
	PaymentType uint8 //付款类型， 0:支付，1:转账
	TransferAmount float64 // 转账金额
	OrderTime	time.Time //下单时间
	TransferInfo string // 转账信息
	OrderState uint8 //订单状态
	AutoTransfer uint8 //是否自动打款 0否，1是，默认是1

}


//商家编号

type Shop_Indexes struct {
	ID     int   `json:"id"`
	ShopNo int   `json:"shop_no"`
	Status uint8 `json:"status"`
}

func AddOrder(data *Orders) error {
	order := Orders{
		ThirdTradeNo:   data.ThirdTradeNo,
		OrderNo:        data.OrderNo,
		SubAccountNo:   data.SubAccountNo,
		MchId:		data.MchId,
		Company:        data.Company,
		BranchShop:     data.BranchShop,
		OrderType:      data.OrderType,
		PaymentType:    data.PaymentType,
		TransferAmount: data.TransferAmount,
		TransferInfo:   data.TransferInfo,
		OrderTime:      data.OrderTime, //下单时间
		OrderState:     data.OrderState,
		AutoTransfer:   data.AutoTransfer,
	}
	now := time.Now()
	order.CreateTime = now
	order.UpdateTime = now
	order.OrderTime = now
	dbase := db.Create(&order)
	if dbase.RowsAffected > 0 {
		log.DbInfo("AddOrder","insert success",order,nil)
	} else {
		log.DbInfo("AddOrder","insert fail",order, dbase.Error)
	}
	return nil
}

func Validate(data *OrderServer) error {

	return nil
}

func FindOrderByRange(begin int, end int,appId string) []*Orders {
	var orders []*Orders
	if appId == "00000"{
		db.Where("id > ? AND id <= ?", begin, end).Find(&orders)
		return orders
	}
	db.Where("mch_id = ? AND id > ? AND id <= ?",appId, begin, end).Find(&orders)
	return orders
}

func FindOrderLargerThan(begin int,appId string) []*Orders {
	var orders []*Orders
	if appId =="00000"{
		db.Where("id > ?",begin).Find(&orders)
		return orders
	}
	db.Where("mch_id =? AND id > ?",appId, begin).Find(&orders)
	return orders
}

//给财务服务器推送最新的订单
func GetOrders(appId int, version string, maps interface{}) {

}
