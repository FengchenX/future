package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"sub_account_service/process_monitor/model"
)

type Config struct {
	LocalPort int
	Interval int64
	EmailSender model.EmailSender
	Process map[string]model.Process
}

var configInstance Config

func GetConfigInstance() Config {
	return configInstance
}

func init() {
	var filePath = "./conf/config.toml"
	_,err := toml.DecodeFile(filePath,&configInstance)
	if err != nil {
		logrus.WithField("filePath",filePath).Errorln("load file error!",err)
	}
}

