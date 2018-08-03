package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// token key with emai
var ThreeKey = "ehpos"

// order message three_party
type Order struct {
	gorm.Model
	Time         int64   // Time 为当前消息发送的 unix 时间戳
	Token        string  // Token为md5后的字符串
	OrderId      int64   // 订单id(订单的唯一id)
	CreateTime   string  // 创建时间
	Company      string  // 公司名
	BranchShop   string  // 分店名
	OrderContent []*Dish // 订单明细（json string）
	Price        float64 // 订单价格
}

// order message
type OrderSave struct {
	gorm.Model
	OrderId      int64   `gorm:"not null;unique"` // 订单id(订单的唯一id)
	CreateTime   string  // 创建时间
	Company      string  // 公司名
	BranchShop   string  // 分店名
	OrderContent string  `sql:"type:text"` // 订单明细（json string）
	Price        float64 // 订单价格
}

// dish message
type Dish struct {
	//gorm.Model
	DishName  string // 菜品名字
	DishCount string // 菜品数量
	DishPrice string // 菜品价格（ = 单价 * 数量）
}

// message
type Bill struct {
	gorm.Model
	Time         int64   // Time 为当前消息发送的 unix 时间戳
	Token        string  // Token为md5后的字符串
	OrderId      int64   `gorm:"not null;unique"` // 订单id（与Order.OrderId对应）
	Company      string  // 公司名
	BranchShop   string  // 分店名
	CreateTime   string  // 支付时间
	PayWay       string  // 支付方式
	Drink        float64 // 酒水费
	ServicePrice float64 // 服务费
	Favorable    float64 // 优惠价格
	SumPrice     float64 // 总价
	Price        float64 // 支付价格
	TradeStatus  int64   // 订单状态
	TransferInfo string  //转账结果信息
}

type OrderServer struct {
	ThirdTradeNo   string    //第三方交易号
	OrderNo        string    //商户订单编号
	SubAccountNo   string    //分账编号
	Company        string    //所属公司,
	OrderType      uint8     //订单类型，0:支付宝，1:微信
	PaymentType    uint8     //付款类型， 0:支付，1:转账
	TransferAmount float64   // 转账金额
	OrderTime      time.Time //下单时间
	TransferInfo   string    // 转账信息
	OrderState     uint8     //订单状态
	AutoTransfer   uint8     //是否自动打款 0否，1是，默认是1
}

const (
	InitState      = iota //初始化状态（暂不用）
	Unpaid                //未付款
	Paying                //付款中
	PayFailed             //付款失败
	PayOK                 //付款完成
	AddTradeFailed        //向编号服务添加交易流水失败
	TransferOK            //转账完成
)
