package utils

import (
	"github.com/golang/glog"
	"os"
	"github.com/BurntSushi/toml"
)
//加载toml配置
func LoadToml(file string, conf interface{}) error{
	glog.Infoln(Log("init", "", "finding config ..."), file)
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		glog.Infoln(file+ "配置文件不存在")
		return err
	}
	_, err := toml.DecodeFile(file, conf)
	if err != nil {
		glog.Infoln("DecodeFile Error ...", err)
		return err
	}
	if err != nil {
		glog.Infoln("DecodeFile Error ...", err)
		return err
	}
	glog.Infoln(Log("init", "", "config.Opts()"), conf)
	return nil
}
