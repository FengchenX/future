package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

func init() {
	AppInst()
}

//AppConf 默认配置文件
type AppConf struct {
	APIAddr string
	APIPort string
	MysqlAddr string
}

var std *AppConf

//AppInst 配置单例
func AppInst() *AppConf {
	if std == nil {
		temp := new(AppConf)
		_, err := toml.DecodeFile("config/static/default-conf.toml", temp)
		if err != nil {
			log.Fatal(err)	
		}
		std = temp
		return std
	}
	return std
}