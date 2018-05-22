package main

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

func Pay() {
	client, err := utils.ConnectToRpc(config.Opts().EthAddress)
	if err != nil {
		glog.Infoln(lib.Loger("pay", "请~转账付费"))
		return
	}

	tx := &types.Transaction{}
	client.SendTransaction(context.Background(), tx)
}
