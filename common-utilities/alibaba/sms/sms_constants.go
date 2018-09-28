//author xinbing
//time 2018/9/4 14:29
package sms
const (
	maxRetryTimes=3 //最多重试3次
	smsProduct="Dysmsapi" //#短信API产品名称（短信产品名固定，无需修改）
	smsProductDomain="https://dysmsapi.aliyuncs.com/" //#短信API产品域名（接口地址固定，无需修改）
	defaultRegion="cn-hangzhou"
	version="2017-05-25"
)
// actions
const (
	sendMessageAction="SendSms"				//发送短信
	queryMessageAction="QuerySendDetails"	//查询短信
)
