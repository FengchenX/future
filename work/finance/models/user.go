package models

import (
	"github.com/jinzhu/gorm"
)

//User 用户数据模型
type User struct {
	ID        uint `gorm:"primary_key"`
	UserName  string
	Password  string
	Authority string
}

// 用户账户详情
type UserAccount struct {
	gorm.Model
	OrderType   int64   `json:"order_type" gorm:"column:order_type"`                           // 订单类型，0:支付宝，1:微信
	BillId      uint    `json:"bill_id" gorm:"column:bill_id"`                                 // 支出流水ID
	Name        string  `json:"name,omitempty" gorm:"column:name"`                             // 用户姓名
	BankCard    string  `json:"bank_card,omitempty" gorm:"column:bank_card"`                   // 用户银行卡
	WeChat      string  `json:"wechat,omitempty" gorm:"column:wechat"`                         // 用户微信
	Alipay      string  `json:"alipay,omitempty" gorm:"column:alipay"`                         // 用户支付宝
	Telephone   string  `json:"telephone,omitempty" gorm:"column:telephone"`                   // 用户电话
	Address     string  `json:"address" gorm:"column:address;index:user_account;unique_index"` // 用户地址
	NoPaySum    float64 `json:"no_pay_sum" gorm:"column:no_pay_sum"`                           // 用户未分账的钱
	Arrears     float64 `json:"arrears" gorm:"column:arrears"`                                 //欠款详情
	TradeStatus uint8   `json:"trade_status" gorm:"column:trade_status"`                       // 分账状态
}

var ScheduleAccountArrears = make(map[string]float64, 100)

func GetScheduleAccount(subAccountNo string) float64 {
	if v, ok := ScheduleAccountArrears[subAccountNo]; ok {
		return v
	}
	return 0
}

// 排班详情
type ScheduleAccount struct {
	gorm.Model
	UserAddress  string  `json:"user_address" gorm:"column:user_address"`                                       //
	UserKeyStore string  `json:"user_keystore" gorm:"column:user_keystore"`                                     //
	UserParse    string  `json:"user_parse" gorm:"column:user_parse"`                                           //
	KeyString    string  `json:"keystring" gorm:"column:keystring"`                                             //
	ScheduleName string  `json:"schedule_name" gorm:"column:schedule_name;index:schedule_account;unique_index"` // 排班号
	Arrears      float64 `json:"arrears" gorm:"column:arrears"`                                                 //欠款详情
}

// 退款信息
type RefundTrade struct {
	gorm.Model
	OrderType      int64   `json:"order_type" gorm:"column:order_type;index:refund_trade"`          // 订单类型，0:支付宝，1:微信
	AlipaySeller   string  `json:"-" gorm:"column:alipay_seller"`                                   // 商家账户
	BuyerLogonID   string  `json:"buyer_logon_id" gorm:"column:buyer_logon_id"`                     // 买家支付宝账号
	BuyerPayAmount float64 `json:"buyer_pay_amount" gorm:"column:buyer_pay_amount"`                 // 买家实付金额，单位为元，两位小数。
	BuyerUserID    string  `json:"buyer_user_id" gorm:"column:buyer_user_id"`                       // 买家在支付宝的用户id
	BuyerUserType  string  `json:"buyer_user_type" gorm:"column:buyer_user_type"`                   // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	InvoiceAmount  float64 `json:"invoice_amount" gorm:"column:invoice_amount"`                     // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
	OutTradeNo     string  `json:"out_trade_no" gorm:"column:out_trade_no"`                         // 商家订单号
	PointAmount    float64 `json:"point_amount" gorm:"column:point_amount"`                         // 积分支付的金额，单位为元，两位小数。
	ReceiptAmount  float64 `json:"receipt_amount" gorm:"column:receipt_amount"`                     // 实收金额，单位为元，两位小数
	SendPayDate    int64   `json:"send_pay_date" gorm:"column:send_pay_date"`                       // 本次交易打款给卖家的时间
	TotalAmount    float64 `json:"total_amount" gorm:"column:total_amount"`                         // 交易的订单金额
	Fees           float64 `json:"fees" gorm:"column:fees"`                                         //交易手续费
	TradeNo        string  `json:"trade_no" gorm:"column:trade_no;index:refund_trade;unique_index"` // 支付宝交易号
	SubAccountNo   string  `json:"sub_account_no" gorm:"column:sub_account_no"`                     //分账信息编号
	PayerName      string  `json:"payer_name" gorm:"column:payer_name"`                             // 付款方
	AutoTransfer   uint8   `json:"auto_transfer" gorm:"column:auto_transfer"`                       //是否自动打款 0否，1是，默认是1
	Hash           string  `json:"hash" gorm:"column:hash"`                                         //上链Hash
	TradeStatus    uint8   `json:"trade_status" gorm:"column:trade_status"`                         // 交易状态
	//AuthTradePayMode string // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
}
