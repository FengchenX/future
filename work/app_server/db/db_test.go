package db

import (
	"testing"
	"time"
)

type HistoryDeploy struct {
	Id         uint      "gorm:PRIMARY KEY"
	CreateTime time.Time //创建时间
	Address    string    // 地址
	Account    string    // 账户
	Version    string    // 版本
	SubVersion string    // 子版本號
	RequsetIp  string    // 請求地址
	Menthon    string    // 請求方式
}

type Company struct {
	ID               uint      "gorm:PRIMARY KEY"
	CompanyName      string    `gorm:"not null"`
	CompanyBanchName string    `gorm:"not null"`
	CompanyPassWord  string    `gorm:"not null"`
	CompanyAddress   string    `gorm:"not null"`
	CompanyPhone     string    `gorm:"not null"`
	CompanyDespcrite string    `gorm:"type:text;not null"`
	CreateTime       time.Time `gorm:"not null"`
}

// order message
type OrderSave struct {
	OrderId      int64     // 订单id(订单的唯一id)
	CreateTime   time.Time // 创建时间
	Company      string    // 公司名
	BranchShop   string    // 分店名
	OrderContent string    // 订单明细（json string）
	Price        float64   // 订单价格
}

// dish message
type Dish struct {
	DishName  string // 菜品名字
	DishCount string // 菜品数量
	DishPrice string // 菜品价格（ = 单价 * 数量）
}

// message
type Bill struct {
	OrderId      int64     // 订单id（与Order.OrderId对应）
	Company      string    // 公司名
	BranchShop   string    // 分店名
	CreateTime   time.Time // 支付时间
	PayWay       string    // 支付方式
	Drink        float64   // 酒水费
	ServicePrice float64   // 服务费
	Favorable    float64   // 优惠价格
	SumPrice     float64   // 总价
	Price        float64   // 支付价格
}

type Binder struct {
	ID             uint      "gorm:PRIMARY KEY"
	CompanyId      uint      `gorm:"not null"`
	CompanyName    string    `gorm:"not null"`
	BrenchName     string    `gorm:"not null"`
	UserAddress    string    `gorm:"not null"`
	CompanyAddress string    `gorm:"not null"`
	BindTime       time.Time `gorm:"not null"`
}

var huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"
var huangPass = "15219438281"
var huangDesp = "{\"address\":\"759d4c2e15587fae036f183202f36ca3c667ccbd\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"e9e86c32d3c130072d78cd66a58dd9e7dd84c7066652ecb8db983193b759bebb\",\"cipherparams\":{\"iv\":\"ba6e9792b52b22ae3055cb8d2fe8d90b\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"aac7b6c6ab10e0e57446a97f36485800dd5520c159be2be9e8610074c05a6061\"},\"mac\":\"6d103da9f03a32e5d12f79489f1de03fc3ac858c2b9b2bb3d8d3aa2be2fcc498\"},\"id\":\"962d9f5f-91e6-49a5-b6a0-70e3a9b2136e\",\"version\":3}"

func TestInitDb(t *testing.T) {
	InitDb("launch:launch@tcp(39.108.80.66:3306)/test_order?charset=utf8&parseTime=true&loc=Local")
	//DbClient.Client.CreateTable(&Orders{})
	for i := 0; i < 1; i++ {
		DbClient.Client.CreateTable(&Schedule{})
	}
}
