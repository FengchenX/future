package db


import (
	"time"
	"github.com/golang/glog"
	//mysql 驱动
	_ "github.com/go-sql-driver/mysql"
	//gorm 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var mysql = "root:launch*2018@tcp(172.18.22.48:3306)/old_data?charset=utf8&parseTime=true&loc=Local"

//DB 数据库客户端
var DB *gorm.DB

//InitOrderDB 初始化db
func InitOrderDB() {
	glog.Infoln("InitOrderDB*****************start", mysql)
	db, err := gorm.Open("mysql", mysql)
	if err != nil {
		glog.Errorln("InitOrderDB**************db err:", err)
	}
	err = db.DB().Ping()
	if err != nil {
		glog.Errorln("db ping fail", err)
	} else {
		glog.Infoln("InitOrderDB************success")
	}
	DB = db
}

//GetOrder 拿订单
func GetOrder(pages int) ([]NumberOrder, int) {
	var temps []NumberOrder
	mydb := DB.Select("third_trade_no, order_no, sub_account_no, company, order_type, payment_type, transfer_amount, transfer_info, auto_transfer, create_time, update_time").
			Table("orders").
			Where("sub_account_no = ?", "AC:a8f9037a3cdb6d42fa4c0eae1e727e62").
			Limit(Size).Offset((pages-1)*Size).
			Find(&temps)
	if mydb.Error != nil {
		glog.Errorln("GetOrder**************db err:", mydb.Error)
		return nil, pages
	}
	return temps, pages
}
//Size 每次拿的大小
const Size = 100

//GetOrder2 拿订单
func GetOrder2(pages int) ([]Order, int) {
	var temps []Order
	mydb := DB.Select("expenses_bills.sub_account_no, expenses_bills.trade_no, user_bills.address, user_bills.alipay, user_bills.name, user_bills.money, user_bills.radio, user_bills.sub_way, orders.create_time").
			Table("expenses_bills").
			Joins("join user_bills on user_bills.bill_id = expenses_bills.id").
			Joins("join orders on expenses_bills.sub_account_no = orders.sub_account_no and expenses_bills.trade_no = orders.third_trade_no").
			Where("expenses_bills.sub_account_no = ?", "AC:a8f9037a3cdb6d42fa4c0eae1e727e62").
			Find(&temps)
	if mydb.Error != nil {
		glog.Errorln("GetOrder***************db err:", mydb.Error)
		return nil, pages
	}
	return temps, pages
}

//Order 订单
type Order struct {
	SubAccountNo string
	TradeNo string
	Address string
	Alipay string
	Name string
	Money float64
	Radio float64
	SubWay int64
	CreateTime time.Time
}
