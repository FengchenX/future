/*--------------------------------------------------------------
*  package: 初始化服务
*  time:    2018/04/17
*-------------------------------------------------------------*/
package api

import (
	"encoding/json"
	"github.com/golang/glog"
	"strconv"
	"sub_account_service/blockchain_server/arguments"
	c "sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/contracts"
	"sub_account_service/blockchain_server/lib"
	myeth "sub_account_service/blockchain_server/lib/eth"
	"time"
)

//DeployAddress 部署地址
var DeployAddress string

var init_key string

// 一期
//var addr string = "0x59f1b27caf3d72cd6edd87d3142991a2a6f35420"
//var acco string = "0xd46966b4b199332a9a03c8e7996c9c6449e426bf"

func Init() {
	glog.Infoln("api init enter")
	DeployAddress = c.ConfInstance().DeployAddress
	ManageMentInit()
}

// 初始化合约
func ManageMentInit() string {

	glog.Infoln(lib.Loger("initing", "print"), c.Opts().AccAddress, c.Opts().DeployAddress)

	if len(c.Opts().DeployAddress) == 0 {
		key := myeth.ParseKeyStore(c.Opts().ManagerKey, c.Opts().ManagerPhrase)
		b, err := json.Marshal(key)
		if err != nil {
			glog.Infoln(err)
			panic(err)
		}

		num := strconv.Itoa(int(time.Now().Unix()))

		glog.Infoln(string(b))

		addr, _, err := contracts.Deploy(string(b),
			arguments.DeployArguments{
				TokenName:   c.Opts().ServerId + num,
				TokenSymbol: num,
				SubPayer:    c.Opts().PayAddress,
				Postscript:  "",
			})

		if addr == "" || err != nil {
			glog.Errorln("[init error]: associated account fail ", addr)
			return ""
		}
		glog.Infoln(lib.Log("initing", "", "Associated account"), addr)
		DeployAddress = addr

		return addr
	}
	return ""
}
