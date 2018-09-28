//author xinbing
//time 2018/9/11 10:45
//
package main

import (
	"common-utilities/db"
	"common-utilities/logs"
	"ibs_service/saas_background/routes"
	"ibs_service/saas_background/saas_config"
	"time"
)

func main() {
	db.InitRedis(&db.RedisConfig{
		RedisAddr: saas_config.GetConfigInstance().RedisAddr,
		RedisPwd: saas_config.GetConfigInstance().RedisPwd,
		RedisDB: saas_config.GetConfigInstance().RedisDB,
	})
	logs.LogInit(saas_config.GetConfigInstance().LogPath, saas_config.GetConfigInstance().LogFileName, time.Duration(saas_config.GetConfigInstance().LogMaxAge * 24) * time.Hour, time.Duration(saas_config.GetConfigInstance().LogRotationTime) * time.Hour)
	routes.Start()
}