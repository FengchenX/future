package entity

import (
	"time"
	"github.com/jinzhu/gorm"
)

//OrderDetail
type OrderDetail struct {
	ID uint `gorm:"primary key;auto_increment"`
	OrderId uint `gorm:"index:order_id_idx"`
	ProductCode string //产品编码
	ProductName string //产品名称
	OriginPrice float64 `gorm:"decimal(10,4)"` //原价
	Price float64 `gorm:"decimal(10,4)"`//实际价格
	Count float64 `gorm:"decimal(10,4)"`//数量
	Subtotal float64 `gorm:"decimal(10,4)"`//小计
	Unit string //单位
}
// message
type Order struct {
	gorm.Model
	ThirdTradeNo string `gorm:"index:third_trade_no_idx"` //三方交易号
	OrderNo 	 string `gorm:"index:order_no_idx"` //自己平台的交易单号
	SubNumber		 string // 排班编号
	AppId		 string //appId
	CompanyId	 int `gorm:"index:company_id_idx"`  // 所属公司ID
	CompanyName	 string //所属公司
	BranchName	 string //公司门店
	OrderTime    time.Time  // 支付时间
	PayWay       string  // 支付方式，支付宝，微信，员工卡
	Discount    float64 `sql:"type:decimal(8,2)"`// 优惠价格
	SumPrice     float64 `sql:"type:decimal(8,2)"` // 总价
	Price        float64 `sql:"type:decimal(8,2)"`// 支付价格
	BillState	 BillState //订单状态
	Complete	 int //是否添加到编号服务器
	Salesmen     string
	OrderContent string `sql:"type:text"` // 订单明细（json string）
	Remarks string //备注
}

type BillState int
const (
	Normal BillState = 1 //正常
	Refund BillState = 2 //退款
)
