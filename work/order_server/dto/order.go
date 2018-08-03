package dto

import (
	"sub_account_service/order_server/entity"
	"errors"
)

type BillOrderSaveDTO struct {
	ThirdTradeNo string  //三方交易号
	OrderNo 	 string //自己平台的交易单号
	OrderTime    string // 创建时间
	SubNumber	 string // 分配表编号
	SubNumberTag string
	Company      string // 公司名
	BranchShop   string // 分店名
	//OrderDetail string  // 订单明细（json string）
	OrderDetails *[]*entity.OrderDetail
	//order table
	PayWay       string  // 支付方式
	Price        float64 // 订单价格
	Discount    float64 // 优惠价格
	SumPrice     float64 // 总价
	OrderState	 entity.BillState //订单状态
	Salesmen	string //售货员
	Remarks string
	//table end
}

func (p *BillOrderSaveDTO) Validate() error{
	if p.OrderNo == "" {
		return errors.New("订单编号不允许为空！")
	} else if p.OrderState != entity.Normal && p.OrderState != entity.Refund {
		return errors.New("订单状态不正确！")
	}
	return nil
}
