package models

import "github.com/jinzhu/gorm"

// token key with emai
var ThreeKey = "ehpos"

// order message three_party
type Order struct {
	gorm.Model
	Time          int64 // Time 为当前消息发送的 unix 时间戳
	Token         string // Token为md5后的字符串
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
	OrderId      int64                    // 订单id(订单的唯一id)
	CreateTime   string                   // 创建时间
	Company      string                   // 公司名
	BranchShop   string                   // 分店名
	OrderContent string `sql:"type:text"` // 订单明细（json string）
	Price        float64                  // 订单价格
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
	Time          int64  // Time 为当前消息发送的 unix 时间戳
	Token         string // Token为md5后的字符串
	OrderId      int64   // 订单id（与Order.OrderId对应）
	Company      string  // 公司名
	BranchShop   string  // 分店名
	CreateTime   string  // 支付时间
	PayWay       string  // 支付方式
	Drink        float64 // 酒水费
	ServicePrice float64 // 服务费
	Favorable    float64 // 优惠价格
	SumPrice     float64 // 总价
	Price        float64 // 支付价格
}
