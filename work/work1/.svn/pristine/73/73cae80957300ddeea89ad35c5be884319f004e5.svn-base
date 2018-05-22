package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	"znfz/server/api"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/protocol"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "../conf/config.toml", "config path")
	flag.Parse()
	config.ParseToml(path)                  // 初始化配置
	addr, acc := apiserver.ManageMentInit() // 初始化合约

	listener, err := net.Listen("tcp", config.Opts().LocalAddress+config.Opts().Port)
	if err != nil {
		glog.Errorln("net err", err)
		return
	}
	glog.Infoln(lib.Log("initing", "", "starting service"), config.Opts().LocalAddress+config.Opts().Port)
	rpcServer := grpc.NewServer()
	myServer := &apiserver.ApiService{
		AddressManager: addr,
		AccountAddress: acc,
	}

	protocol.RegisterApiServiceServer(rpcServer, myServer)
	// 注册rpc服务
	err = rpcServer.Serve(listener)
	if err != nil {
		glog.Errorln("rpc err", err)
		return
	} else {
		glog.Infoln("start grpc process!")
	}

}
