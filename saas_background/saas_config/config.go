//author xinbing
//time 2018/9/11 10:39
//
package saas_config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type config struct {
	Port					string
	LogPath					string
	LogFileName				string
	LogMaxAge 				int
	LogRotationTime			int
	RedisAddr				string
	RedisPwd				string
	RedisDB					int
	MySqlAddr				string
	OrderServerMySqlAddr	string
}
var cf *config

func GetConfigInstance() *config {
	return cf
}

func init() {
	filePath := "./conf/config.toml"
	var c config
	entry := logrus.WithField("filePath", filePath)
	_, err := toml.DecodeFile(filePath, &c)
	if err != nil {
		entry.WithError(err).Errorln("init config error!")
		return
	}
	cf = &c
	entry.WithField("config", c).Infoln("init config success!")
}