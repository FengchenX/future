//author xinbing
//time 2018/9/4 14:11
package sms

import "time"

type AliCloudSMSSendReq struct {
	PhoneNumbers 	string //必填 发送的电话号码，如有多个以英文逗号分割
	SignName 		string //必填 短信签名，可在阿里云短信控制台里面设置
	TemplateCode 	string //必填 模板ID
	TemplateParam 	map[string]string //可选 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议。
	SmsUpExtendCode string	//上行短信扩展码,无特殊需要此字段的用户请忽略此字段
	OutId			string  //outId为提供给业务方扩展字段,最终在短信回执消息中将此值带回给调用者
}

type AliCloudSMSSendResp struct {
	RequestId		string	//请求ID
	Code			string	//状态码-返回OK代表请求成功,其他错误码详见错误码列表
	Message			string	//状态码的描述
	BizId			string	//发送回执ID,可根据该ID查询具体的发送状态
}

//是否发送成功
func (p *AliCloudSMSSendResp) ISSuccess() bool {
	return p.Code == "OK"
}

type AliCloudSMSQueryReq struct {
	PhoneNumber		string //要查询的号码
	BizId			string //发送短信返回的回执
	SendDate		time.Time //短信发送日期格式yyyyMMdd,支持最近30天记录查询
	PageSize		int		//每页大小，最大50
	CurrentPage		int		//当前页
}

type AliCloudSMSQueryResp struct {
	RequestId 	string  //
	Code 		string 	// OK 	状态码-返回OK代表请求成功,其他错误码详见错误码列表
	Message 	string 	// 请求成功 	状态码的描述
	TotalCount 	int 	//100
	smsSendDetailDTOs 	[]*AliCloudSMSDetail
}

type AliCloudSMSDetail struct {
	PhoneNum 		string
	SendStatus 		int
	ErrCode 		string
	TemplateCode	string
	Content 		string
	SendDate 		string
	ReceiveDate 	string
	OutId 			string
}