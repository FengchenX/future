//author xinbing
//time 2018/9/3 16:56
package sms

import (
	"common-utilities/http_utils"
	"common-utilities/utilities"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type AliCloudSMSClient struct {
	AccessKeyId				string //必填
	AccessKeySecret			string //必填
	RegionId                string //区域
	Version					string //
}

func GetAliCloudSMSClient(accessKeyId, accessKeySecret string) *AliCloudSMSClient {
	client := &AliCloudSMSClient{
		AccessKeyId:accessKeyId,
		AccessKeySecret:accessKeySecret,
		RegionId:defaultRegion,
		Version:version,
	}
	return client
}

// un implement
func GetAliCloudSMSClientByProfile() *AliCloudSMSClient {
	return nil
}

func (p *AliCloudSMSClient) SendMessage(req *AliCloudSMSSendReq) (*AliCloudSMSSendResp,error) {
	params := p.getBaseParams()
	content,err := json.Marshal(req.TemplateParam)
	if err != nil {
		return nil, errors.New("parse template param error:"+err.Error())
	}
	params["Action"] = sendMessageAction
	params["PhoneNumbers"] = req.PhoneNumbers
	params["SignName"] = req.SignName
	params["TemplateCode"] = req.TemplateCode
	params["TemplateParam"] = string(content)
	params["SmsUpExtendCode"] = req.SmsUpExtendCode
	params["OutId"] = req.OutId
	queryStr,signature := p.generateQueryStringAndSignature(params)
	reqUrl := smsProductDomain + "?Signature=" + signature + queryStr
	fmt.Println("reqUrl:", reqUrl)
	content, err = http_utils.Get(reqUrl, nil)
	if err != nil {
		return nil, errors.New("req to send message error:"+err.Error())
	}
	//fmt.Println(string(content))
	var resp AliCloudSMSSendResp
	err = json.Unmarshal(content, &resp)
	if err != nil {
		return nil, errors.New("unmarshal resp error:"+err.Error())
	}
	return &resp, nil
}

func (p *AliCloudSMSClient) getBaseParams() map[string]string {
	baseParams := make(map[string]string)
	timeStamp, nonce := getTimestampAndNonce()
	baseParams["SignatureMethod"] = "HMAC-SHA1" //固定值HMAC-SHA1,
	baseParams["SignatureNonce"] = nonce
	baseParams["AccessKeyId"] = p.AccessKeyId
	baseParams["SignatureVersion"] = "1.0"
	baseParams["Timestamp"] = timeStamp
	baseParams["Format"] = "JSON" //xml

	baseParams["Version"] = p.Version
	baseParams["RegionId"] = p.RegionId
	return baseParams
}

func (p *AliCloudSMSClient) generateQueryStringAndSignature(params map[string]string) (string, string) {
	delete(params,"Signature")
	keys := make([]string, 0)
	for key,item := range params {
		if item != "" {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	sortQueryStringTmp := ""
	for _, key := range keys {
		rstKey := specialUrlEncode(key)
		rstVal := specialUrlEncode(params[key])
		sortQueryStringTmp = sortQueryStringTmp + "&" + rstKey + "=" + rstVal
	}

	sortQueryString := strings.Replace(sortQueryStringTmp, "&", "", 1)
	stringToSign := "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortQueryString)
	//fmt.Println("stringToSign:",stringToSign)
	sign := sign(p.AccessKeySecret+"&", stringToSign)
	//fmt.Println("before special:", sign)
	signature := specialUrlEncode(sign)
	//fmt.Println("queryStr:", sortQueryString)
	return sortQueryStringTmp, signature
}

func specialUrlEncode(value string) string {
	rstValue := url.QueryEscape(value)
	rstValue = strings.Replace(rstValue, "+", "%20", -1)
	rstValue = strings.Replace(rstValue, "*", "%2A", -1)
	rstValue = strings.Replace(rstValue, "%7E", "~", -1)
	return rstValue
}

func sign(accessKeySecret, sortQueryStr string) string {
	h := hmac.New(sha1.New, []byte(accessKeySecret))
	h.Write([]byte(sortQueryStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

var logGMT,_ = time.LoadLocation("GMT")
func getTimestampAndNonce() (string,string) {
	nonce := utilities.GetRandomStr(32)
	t := time.Now()
	timeStamp := t.In(logGMT).Format("2006-01-02T15:04:05Z")
	return timeStamp, nonce
}
