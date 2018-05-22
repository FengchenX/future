package web

import (
	"bytes"
	"github.com/golang/glog"
	"net/http"
	"strings"
	conf "znfz/conf_server/config"
	contract "znfz/conf_server/contract"
	"znfz/conf_server/rpc"
	"znfz/conf_server/lib"
)

func WebInit() {
	var addr string = conf.Opts().LocalAddress + conf.Opts().WebPort
	glog.Infoln(lib.Log("web err", "", "initing web server"), addr)

	http.HandleFunc("/reloaddeploy", deploy)
	http.HandleFunc("/reloadconfig", config)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		glog.Errorln(lib.Log("web err", "", "ListenAndServe"), err)
	}
}

// 发布初始化合约
func deploy(w http.ResponseWriter, r *http.Request) {
	glog.Infoln(lib.Log("web", "", "deploy"))
	url := r.URL.RequestURI()
	glog.Infoln(url)
	strs := strings.Split(url, "?")
	if len(strs) < 2 {
		glog.Errorln(lib.Log("web err", "", "deploy"))
		w.Write(bytes.NewBufferString("[Deply Error] Not enough args!").Bytes())
		return
	}
	contents := strings.Split(strs[1], "&")

	addr, acco := contract.DeployInit(contents[0], contents[1], r.RemoteAddr)
	w.Write(bytes.NewBufferString("[Deply Success] address:" + addr + " account:" + acco).Bytes())

	rpc.ReloadDeploy(addr, acco)
}

// 重载配置文件
func config(w http.ResponseWriter, r *http.Request) {
	glog.Infoln(lib.Log("web", "", "config"))
	rpc.ReloadConfig()
	w.Write(bytes.NewBufferString("[ReloadConfig Success] ").Bytes())
}
