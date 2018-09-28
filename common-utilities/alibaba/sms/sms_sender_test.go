//author xinbing
//time 2018/9/3 16:57
package sms

import (
	"fmt"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	smsClient := GetAliCloudSMSClient("LTAI5TR6wXkOVAUg","bNMQJkDtxXn5O8pth2w0f6h5VliV50")
	resp, err := smsClient.SendMessage(&AliCloudSMSSendReq{
		PhoneNumbers: "18520227050",
		SignName: 	  "立行共享车",
		TemplateCode: "SMS_143714609",
		TemplateParam: map[string]string{
			"code":"123123",
		},
		OutId: "wx233233i",
	})
	fmt.Println(resp)
	fmt.Println(err)
}

func TestSendMessage2(t *testing.T) {
	loc,_ := time.LoadLocation("GMT")
	ti := time.Now()
	fmt.Println(ti.Location())
	fmt.Println(ti.In(loc).Format("2006-01-02 15:04:05"))
	fmt.Println(ti.Format("2006-01-02 15:04:05"))
}