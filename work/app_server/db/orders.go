package db

import (
	"github.com/golang/glog"
	"strings"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/lib"
	"time"
	"github.com/jinzhu/gorm"
)

//NumberOrder 订单
type NumberOrder struct {
	ID             int       `gorm:"primary_key" json:"-"`
	ThirdTradeNo   string    `json:"third_trade_no"`  //第三方交易号
	OrderNo        string    `json:"order_no"`        //商户订单编号
	SubAccountNo   string    `json:"sub_account_no"`  //分账编号
	Company        string    `json:"company"`         //所属公司,
	OrderType      uint8     `json:"order_type"`      //订单类型，0:支付宝，1:微信
	PaymentType    uint8     `json:"payment_type"`    //付款类型， 0:支付，1:转账
	TransferAmount float64   `json:"transfer_amount"` // 转账金额
	TransferInfo   string    `json:"transfer_info"`   // 转账信息
	AutoTransfer   uint8     `json:"auto_transfer"`   //是否自动打款 0否，1是，默认是1
	OrderTime      time.Time `json:"order_time"`
	CreateTime     time.Time `json:"create_time"`     //创建时间
	UpdateTime     time.Time `json:"update_time"`     //更新时间

	//OrderState     uint8     //订单状态 退单
	Read bool `json:"read"` //是否已读
	DumpNum int `gorm:"not null;default 0"`
}

//AppOrder 账本
type AppOrder struct {
	gorm.Model        
	Name       string    `gorm:"type:text;not null"`
	SubCode    string    `gorm:"type:text;not null"`
	OrderId    string    `gorm:"type:text;not null"`
	UserAddr   string    `gorm:"not null"`
	CreateTime time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	Ratio      float64   `gorm:"not null"`
	Price      float64   `gorm:"not null"`
	SubWay     uint32    `gorm:"not null"`
	PayAccount string    `gorm:"not null"`
}

//InsertOrder 插入账本
func InsertOrder(sn, od, usr, pa, na string, pr, ra float64, sw uint32, orderTime time.Time) {
	a := &AppOrder{}
	DbClient.Client.Find(a, "order_id = ? AND user_addr = ? AND sub_code = ?", od, usr, sn)
	if strings.ToLower(a.SubCode) == strings.ToLower(sn) {
		glog.Infoln("InsertOrder*************db 中已经包含")
		return
	}

	// inserting New one

	db := DbClient.Client.Create(&AppOrder{
		SubCode:    sn,
		OrderId:    od,
		CreateTime: orderTime,
		UserAddr:   usr,
		Price:      pr,
		Ratio:      ra,
		SubWay:     sw,
		PayAccount: pa,
		Name:       na,
	})
	glog.Infoln("InsertOrder*****************appOrder:", sn, od, orderTime, usr, pr, ra, sw, pa, na)
	if db.Error != nil {
		glog.Errorln("InsertOrder err", db.Error)
	}
}

//NotFindOrder true 没有找到 false 找到了
func NotFindOrder(od, sn string) bool {
	// a := &AppOrder{}
	// db := DbClient.Client.Find(a, "order_id = ?", od)
	// if db.Error != nil {
	// 	glog.Infoln("FindOrder*******************not find")
	// 	return true
	// }

	// if strings.ToLower(a.SubCode) == strings.ToLower(sn) {
	// 	return false
	// }

	var numOrders []NumberOrder
	DB := DbClient.Client

	if mydb := DB.Where("third_trade_no = ? AND sub_account_no = ?", od, sn).Find(&numOrders); mydb.Error != nil {
		glog.Errorln("NotFindOrder************db err:", mydb.Error)
		//发生错误当作找到了
		return false
	}
	if len(numOrders) > 0 {
		glog.Infoln("NotFindOrder***********had find", od, sn)
		return false
	}

	return true
}

//GetAllMoney 拿所有的money
func GetAllMoney(userAddr string) (float64, float64, float64) {
	var orders []*AppOrder

	db := DbClient.Client.Where("user_addr = ?", userAddr).Find(&orders)
	glog.Infoln("len", len(orders))
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}

	all := float64(0)
	m := make(map[time.Time]float64)
	mm := make(map[time.Time]float64)

	for i, order := range orders {
		t := lib.Today(order.CreateTime)
		mt := lib.Month(order.CreateTime)
		//glog.Infoln("day:", t, "   month:", mt)
		if _, ok := m[t]; ok {
			m[t] += order.Price
		} else {
			m[t] = order.Price
		}
		if _, ok := mm[mt]; ok {
			mm[mt] += order.Price
		} else {
			mm[mt] = order.Price
		}
		all += order.Price
		glog.Infoln(i, "---------", order.Price)
	}

	var tod float64 = 0
	var mon float64 = 0

	if v, exist := m[lib.Today(time.Now())]; exist {
		tod = v
	}
	if v, exist := mm[lib.Month(time.Now())]; exist {
		mon = v
	}

	return all, tod, mon
}

//GetMoney 去某个人的钱
func GetMoney(userAddr string, b, e time.Time, p int) ([]*AppOrder, int) {

	var count int
	db := DbClient.Client.Model(&AppOrder{}).Where("user_addr = ? AND create_time > ? AND create_time < ?", userAddr, b, e).Count(&count)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
	var orders []*AppOrder
	limit := config.Optional.CheckLimit
	if p == 0 {
		p = 1
	}

	db = DbClient.Client.Where("user_addr = ? AND create_time > ? AND create_time < ?", userAddr, b, e).
		Limit(limit).Offset(limit * (p - 1)).Find(&orders)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}

	glog.Infoln(count)

	pageC := int(count) / int(limit)
	if int(count)%int(limit) > 0 {
		pageC++
	}

	return orders, pageC
}
