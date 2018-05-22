package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	"znfz/conf_server/config"
	"znfz/conf_server/db"
	"znfz/conf_server/protocol"
	"znfz/conf_server/rpc"
	"znfz/conf_server/web"
	"znfz/server/lib"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "../conf/config.toml", "config path")
	flag.Parse()

	// 初始化配置
	config.ParseToml(path)

	// 初始化db
	db.InitDb(config.Opts().MysqlStr)

	// 初始化web
	go web.WebInit()

	// 初始化rpc
	rpcInit()
}

func rpcInit() {
	listener, err := net.Listen("tcp", config.Opts().LocalAddress+config.Opts().Port)
	if err != nil {
		glog.Errorln(lib.Loger("rpc initing err", "starting service"), err)
		return
	}

	glog.Infoln(lib.Loger("rpc initing", "starting service"), config.Opts().LocalAddress+config.Opts().Port)
	rpcServer := grpc.NewServer()
	protocol.RegisterConfServerServer(rpcServer, &rpc.ConfigServer{})

	// 注册rpc服务
	err = rpcServer.Serve(listener)
	if err != nil {
		glog.Errorln(lib.Loger("rpc initing err", "starting service"), err)
		return
	} else {
		glog.Infoln(lib.Loger("rpc initing err", "starting service success!"))
	}
}
