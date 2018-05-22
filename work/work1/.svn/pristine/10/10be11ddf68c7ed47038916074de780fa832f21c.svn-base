package main

import (
	"znfz/server/config"
	"znfz/server/client/addrmanager"
	"github.com/golang/glog"
	"znfz/server/client/accmanager"
)

// 重新发布协议
func DeployInit() (addr string, acco string) {
	serId := config.Opts().ServerId
	// 部署地址
	address, address_name, err := addrmanager.Deploy("address_Init "+serId, "address_Init_s "+serId)
	if address == "" || address_name == "" || err != nil {
		glog.Errorln("[init error]: address is nil ", address, address_name, err)
		return "", ""
	}
	glog.Infoln("[initing]: 发布地址:", address, address_name)

	// 部署账户
	account, account_name, _ := accmanager.Deploy("account_Init "+serId, "s "+serId)
	if account == "" || account_name == "" {
		glog.Errorln("[init error]: accmanager_addr is nil ", account, account_name)
		return "", ""
	}
	glog.Infoln("[initing]: 发布账户:", account, "accmanager_smart_name:", account_name)

	return address, account
}

