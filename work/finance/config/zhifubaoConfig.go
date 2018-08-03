package config

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"

	"sub_account_service/finance/lib"
)

type zhifubaoConfig struct {
	AlipayAppID      string
	AlipayUrl        string
	AlipayPrivateKey string
	AlipayPublicKey  string
	Format           string
	Charset          string
	SignType         string
	GoodsType        string
	AlipaySeller     string
	RefundUrl        string
}

//private
var zfbConfig = zhifubaoConfig{}

func ZfbConfig() zhifubaoConfig {
	return zfbConfig
}
func init() {
	file := "../conf/zhifubao_config.toml"
	glog.Infoln(lib.Log("initing", "", "finding config ..."), file)
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(Opts()); err != nil {
			glog.Infoln("如果配置文件不存在 ...")
		}
		glog.Infoln("没有找到配置文件，创建新文件 ...")
	}
	var conf zhifubaoConfig
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		glog.Infoln("DecodeFile Error ...", err)
	}
	zfbConfig = conf
	glog.Infoln(lib.Log("initing", "", "config.zhifubaoConfig()"), zfbConfig)
}
