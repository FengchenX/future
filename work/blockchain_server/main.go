package main

import (
	"flag"
	API "sub_account_service/blockchain_server/api"
	"sub_account_service/blockchain_server/lib/eth"
	"sub_account_service/blockchain_server/router"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	myeth.NewNonceMap()
	API.Init()
	router.Init()
}
