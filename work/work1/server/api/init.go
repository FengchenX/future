/*--------------------------------------------------------------
*  package: 初始化服务
*  time:    2018/04/17
*-------------------------------------------------------------*/
package apiserver

import (
	"github.com/golang/glog"
	"znfz/server/arguments"
	c "znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/accmanager"
	"znfz/server/token-contract/addrmanager"
)

var init_key string
// 一期
//var addr string = "0x59f1b27caf3d72cd6edd87d3142991a2a6f35420"
//var acco string = "0xd46966b4b199332a9a03c8e7996c9c6449e426bf"

// 二期
var addr string = "0xC1f42A4722AEA9130328b07E45f61a823520F25C"
var acco string = "0x83dCdd7dc4BAd691607408A68198393398cEf780"

// 初始化合约
func ManageMentInit() (string, string) {
	serId := c.Opts().ServerId

	balance, err := accmanager.GetBalance(acco, c.Opts().EthAddress)
	if balance == "" || err != nil {
		glog.Errorln("[init error]:balance is nil ", balance, err)
		return "", ""
	}
	glog.Infof("%v addr:%v balance:%v", lib.Log("initing", "", "config 1"), acco, balance)

	// 部署地址
	//address, address_name, err := addrmanager.Deploy("address_Init "+serId, "address_Init_s "+serId)
	//if address == "" || address_name == "" || err != nil {
	//	glog.Errorln("[init error]: address is nil ", address, address_name, err)
	//	return "", ""
	//}
	//glog.Infoln("[initing]: 发布地址:", address, address_name)
	//
	// 部署账户
	//account, account_name, _ := accmanager.Deploy("account_Init "+serId, "s "+serId)
	//if account == "" || account_name == "" {
	//	glog.Errorln("[init error]: accmanager_addr is nil ", account, account_name)
	//	return "", ""
	//}
	//glog.Infoln("[initing]: 发布账户:", account, "accmanager_smart_name:", account_name)
	//
	//// 关联账户
	//str, _ := addrmanager.NewAddressAdd(c.Opts().ManagerKey, c.Opts().ManagerPhrase, address, account,
	//	"account_manager "+serId, "account_manager_s "+serId)

	str, _ := addrmanager.NewAddressAdd(arguments.BindSmartArguments{
		OperationKeyStore:     c.Opts().ManagerKey,
		OperationPassWord:  c.Opts().ManagerPhrase,
		OperatingAddress: addr,
		SmartAddress:     acco,
		TokenName:        "account_manager " + serId,
		TokenSymbol:      "account_manager_s" + serId,
		StoresNumber:     "hao,hao",
	})
	if str == "" {
		glog.Errorln("[init error]: associated account fail ", str, addr, acco)
		return "", ""
	}
	glog.Infoln(lib.Log("initing", "", "Associated account"), str)

	//return address, account
	return addr, acco
}
