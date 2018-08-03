package error

var HttpStatus = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	INVALID_PARAMS:       "请求参数错误",
	INVALID_ThirdTradeNo: "第三方交易编号不允许为空！",
	INVALID_OrderNo:      "商户订单编号不允许为空！",
	INVALID_SubAccountNo: "分账编号不允许为空！",
	INVALID_Company:      "所属公司不允许为空！",
	INVALID_OrderType:    "订单类型不合法！",
}

func GetMsg(code int) string {
	msg, ok := HttpStatus[code]
	if ok {
		return msg
	}
	return HttpStatus[ERROR]
}
