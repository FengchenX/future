package dto

import (
	"sub_account_service/order_server/entity"
	"time"
)

type OrderServer struct {
	ThirdTradeNo string //第三方交易号
	OrderNo string //商户订单编号
	SubAccountNo string //分账编号
	MchId	string  //商户ID
	Company string //所属公司,
	BranchShop string //所属分店
	OrderType uint8 //订单类型，0:支付宝，1:微信
	PaymentType uint8 //付款类型， 0:支付，1:转账
	TransferAmount float64 // 转账金额
	OrderTime	time.Time //下单时间
	//TransferInfo string // 转账信息
	OrderState entity.BillState //订单状态
	AutoTransfer uint8 //是否自动打款 0否，1是，默认是1
}
