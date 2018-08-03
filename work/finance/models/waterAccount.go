package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Login 登录
type Login struct {
	UserName string
	Password string
}

//QueryBill 查询流水
type QueryBill struct {
	InBill    bool  //元征账户进账流水
	StartTime int64 //开始时间
	EndTime   int64 //结束时间
}

//IncomeStatement 收入流水
type IncomeStatement struct {
	gorm.Model
	OrderType        int64   `json:"order_type" gorm:"column:order_type;index:income_statement"`          // 订单类型，0:支付宝，1:微信
	AlipaySeller     string  `json:"-" gorm:"column:alipay_seller"`                                       // 商家账户
	BuyerLogonID     string  `json:"buyer_logon_id" gorm:"column:buyer_logon_id"`                         // 买家支付宝账号
	BuyerPayAmount   float64 `json:"buyer_pay_amount" gorm:"column:buyer_pay_amount"`                     // 买家实付金额，单位为元，两位小数。
	BuyerUserID      string  `json:"buyer_user_id" gorm:"column:buyer_user_id"`                           // 买家在支付宝的用户id
	BuyerUserType    string  `json:"buyer_user_type" gorm:"column:buyer_user_type"`                       // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	InvoiceAmount    float64 `json:"invoice_amount" gorm:"column:invoice_amount"`                         // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
	OutTradeNo       string  `json:"out_trade_no" gorm:"column:out_trade_no"`                             // 商家订单号
	PointAmount      float64 `json:"point_amount" gorm:"column:point_amount"`                             // 积分支付的金额，单位为元，两位小数。
	ReceiptAmount    float64 `json:"receipt_amount" gorm:"column:receipt_amount"`                         // 实收金额，单位为元，两位小数
	SendPayDate      int64   `json:"send_pay_date" gorm:"column:send_pay_date"`                           // 本次交易打款给卖家的时间
	TotalAmount      float64 `json:"total_amount" gorm:"column:total_amount"`                             // 交易的订单金额
	Fees             float64 `json:"fees" gorm:"column:fees"`                                             //交易手续费
	Arrears          float64 `json:"arrears" gorm:"column:arrears"`                                       // 欠款详情
	TradeNo          string  `json:"trade_no" gorm:"column:trade_no;index:income_statement;unique_index"` // 支付宝交易号
	SubAccountNo     string  `json:"sub_account_no" gorm:"column:sub_account_no"`                         //分账信息编号
	PayerName        string  `json:"payer_name" gorm:"column:payer_name"`                                 // 付款方
	AutoTransfer     uint8   `json:"auto_transfer" gorm:"column:auto_transfer"`                           //是否自动打款 0否，1是，默认是1
	Hash             string  `json:"hash" gorm:"column:hash"`                                             //上链Hash
	BlockchainStatus uint8   `json:"blockchain_status" gorm:"column:blockchain_status"`                   //交易流水是否上链
	OnchainTime      int64   `json:"on_chain_time" gorm:"column:on_chain_time"`                           //交易流水上链时间
	OnchainNum       int64   `json:"on_chain_num" gorm:"column:on_chain_num"`                             // 重试上链次数
	//AuthTradePayMode string // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
}

//ExpensesBill 支出流水
type ExpensesBill struct {
	gorm.Model
	OrderType      int64   `json:"order_type" gorm:"column:order_type;index:expenses_bill"`          // 订单类型，0:支付宝，1:微信
	AlipaySeller   string  `json:"-" gorm:"column:alipay_seller"`                                    // 商家账户
	BuyerLogonID   string  `json:"buyer_logon_id" gorm:"column:buyer_logon_id"`                      // 买家支付宝账号
	BuyerPayAmount float64 `json:"buyer_pay_amount" gorm:"column:buyer_pay_amount"`                  // 买家实付金额，单位为元，两位小数。
	BuyerUserID    string  `json:"buyer_user_id" gorm:"column:buyer_user_id"`                        // 买家在支付宝的用户id
	BuyerUserType  string  `json:"buyer_user_type" gorm:"column:buyer_user_type"`                    // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	InvoiceAmount  float64 `json:"invoice_amount" gorm:"column:invoice_amount"`                      // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
	OutTradeNo     string  `json:"out_trade_no" gorm:"column:out_trade_no"`                          // 商家订单号
	PointAmount    float64 `json:"point_amount" gorm:"column:point_amount"`                          // 积分支付的金额，单位为元，两位小数。
	ReceiptAmount  float64 `json:"receipt_amount" gorm:"column:receipt_amount"`                      // 实收金额，单位为元，两位小数
	SendPayDate    int64   `json:"send_pay_date" gorm:"column:send_pay_date"`                        // 本次交易打款给卖家的时间
	TotalAmount    float64 `json:"total_amount" gorm:"column:total_amount"`                          // 交易的订单金额
	Fees           float64 `json:"fees" gorm:"column:fees"`                                          //交易手续费
	Arrears        float64 `json:"arrears" gorm:"column:arrears"`                                    // 欠款详情
	TradeNo        string  `json:"trade_no" gorm:"column:trade_no;index:expenses_bill;unique_index"` // 支付宝交易号
	SubAccountNo   string  `json:"sub_account_no" gorm:"column:sub_account_no"`                      //分账信息编号
	PayerName      string  `json:"payer_name" gorm:"column:payer_name"`                              // 付款方
	AutoTransfer   uint8   `json:"auto_transfer" gorm:"column:auto_transfer"`                        //是否自动打款 0否，1是，默认是1
	Hash           string  `json:"hash" gorm:"column:hash"`                                          //上链Hash
	TradeStatus    uint8   `json:"trade_status" gorm:"column:trade_status"`                          // 交易状态
	//AuthTradePayMode string // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
}

//UserBill 用户分账信息
type UserBill struct {
	gorm.Model
	OrderType       int64     `json:"order_type" gorm:"column:order_type"`             // 订单类型，0:支付宝，1:微信
	BillId          uint      `json:"bill_id" gorm:"column:bill_id"`                   // 支出流水ID
	Name            string    `json:"name,omitempty" gorm:"column:name"`               // 用户姓名
	BankCard        string    `json:"bank_card,omitempty" gorm:"column:bank_card"`     // 用户银行卡
	WeChat          string    `json:"wechat,omitempty" gorm:"column:wechat"`           // 用户微信
	Alipay          string    `json:"alipay,omitempty" gorm:"column:alipay"`           // 用户支付宝
	Telephone       string    `json:"telephone,omitempty" gorm:"column:telephone"`     // 用户电话
	Address         string    `json:"address" gorm:"column:address"`                   // 用户地址
	OrderId         string    `json:"order_id" gorm:"column:order_id"`                 // 支付宝交易流水
	TradeNo         string    `json:"trade_no" gorm:"column:trade_no"`                 // 分账的交易信息编号
	Money           float64   `json:"money" gorm:"column:money"`                       // 这笔交易分到的钱
	Arrears         float64   `json:"arrears" gorm:"column:arrears"`                   // 欠款详情
	Rflag           bool      `json:"rflag" gorm:"column:rflag"`                       // 是否更新过打钱后的交易信息，false为未更新，true为已更新。
	TransferDetails string    `json:"transfer_details" gorm:"column:transfer_details"` // 打钱后的交易信息详情，Rflag为true，该字段有值。
	Radio           float64   `json:"radio" gorm:"column:radio"`                       // 分账比例
	SubWay          int64     `json:"sub_way" gorm:"column:sub_way"`                   // 分账方式：1为定额，0为比例
	PaySwitch       int64     `json:"pay_switch" gorm:"column:pay_switch"`             // 分账开关是否打开
	TradeStatus     uint8     `json:"trade_status" gorm:"column:trade_status"`         // 分账状态
	PayDate         time.Time `json:"pay_date" gorm:"column:pay_date"`                 // 付款完成时间
	OnchainTime     int64     `json:"on_chain_time" gorm:"column:on_chain_time"`       // 交易流水上链时间
	OnchainNum      int64     `json:"on_chain_num" gorm:"column:on_chain_num"`         // 重试上链次数
}

// 交易状态
const (
	InitState           = iota //初始化状态（暂不用）
	NotOnchain                 //未上链
	Onchain                    //已完成
	Transfering                //转账中(流水)
	NotTransferOnChain         //准备转账和上链
	Paying                     //付款中
	PayFailed                  //付款失败
	PayOK                      //付款完成
	TransferOnChaining         //用户转账后正在上链
	TransferOnChainFail        //用户转账后上链失败
	TransferOnChainOK          //用户转账后上链成功
	TransferOK                 //转账完成
	RefundFailed               //退款失败
	RefundOK                   //退款成功
)
