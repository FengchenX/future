package config

import (
	"os"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

type Config struct {
	AppId string
	DevelopKey string
	Company string
	BranchShop string
	SubNumbers map[string]string
	ScanInterval int64
	AddOrderUrl string
	GetLatestOrderNoUrl string
	SqlServerUrl string
}
var config *Config
func GetInstance() *Config {
	return config
}

func init() {
	var conf = Config{}
	loadToml("./conf/config.toml", &conf)
	loadToml("./conf/http_config.toml", &conf)
	config = &conf
	glog.Infoln("init config success",config)
}

func loadToml(filePath string, pointer interface{}) {
	glog.Info("init config.....")
	if _,err := os.Stat(filePath);os.IsNotExist(err) {
		glog.Errorln("not found such config file,",filePath)
		return
	}
	_,err := toml.DecodeFile(filePath,pointer)
	if err != nil {
		glog.Errorln("toml 解码错误!",err)
		return
	}
}