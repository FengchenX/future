package alipayModels

import "time"

//给用户转账model
type TransferToAccountReq struct {
	OrderId       string  //这笔转账的id，唯一
	Amount        float64 //转账金额，至少一毛钱，单位为元，只支持2位小数
	PayeeType     string  //收款方账户类型，分为：ALIPAY_USERID（支付宝id），ALIPAY_LOGONID（支付宝账号，如手机号、邮箱），默认值为ALIPAY_LOGONID
	PayeeAccount  string  //收款方账户，根据PayeeType不同而可能不同
	PayerShowName string  //转账人姓名
	PayeeRealName string  //收款人姓名，如果填写，那么会支付宝会校验账号的实名是否与该名称一致，如不一致打款失败
	Remark        string  //转账备注 最多200字符
}

type TransferToAccountResp struct {
	OrderId string    //转账的orderId
	TradeNo string    //支付宝上的这笔交易的交易单号,
	PayDate time.Time //转款时间
}

type QueryTransferToAccountResp struct {
	OrderId        string    // 发起转账来源方定义的转账单据号。 该参数的赋值均以查询结果中 的 out_biz_no 为准。 如果查询失败，不返回该参数
	TradeNo        string    // 支付宝转账单据号，查询失败不返回。
	Status         string    // 转账单据状态
	PayDate        time.Time // 支付时间
	ArrivalTimeEnd time.Time // 预计到账时间，转账到银行卡专用
	OrderFee       float64   // 预计收费金额（元），转账到银行卡专用
}
