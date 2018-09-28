package wechat

import (
	"gopkg.in/chanxuehong/wechat.v2/mch/core"
)

// NewWechatPaySSLClient
func NewWechatPaySSLClient(appId,mchId,key,sslCertPath,sslKeyPath string) (*core.Client,error){
	sslHttpClient,err := core.NewTLSHttpClient(sslCertPath,sslKeyPath)
	if err != nil {
		return nil, err
	}
	client := core.NewClient(appId,mchId,key,sslHttpClient)
	return client, nil
}
