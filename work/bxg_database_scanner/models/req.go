package models

import "sub_account_service/bxg_database_scanner/entity"

type BillOrderSaveReq struct {
	//order_save property
	ThirdTradeNo string  //三方交易号
	OrderNo 	 string //自己平台的交易单号
	SubNumber	 string //分账编号
	SubNumberTag string //分账编号标识
	OrderTime    string // 创建时间
	Company      string // 公司名
	BranchShop   string // 分店名
	OrderDetails *[]*entity.BBB070  // 订单明细（json string）
	//bills table
	PayWay       string  // 支付方式
	Price        float64 // 订单价格
	Discount    float64 // 优惠价格
	SumPrice     float64 // 总价
	OrderState	 int //订单状态
	Salesmen string //售货员
	Remarks string
	//bill table end
}
