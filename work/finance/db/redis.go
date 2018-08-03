package db

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/glog"

	"sub_account_service/finance/config"
)

var RedisClient *redis.Client

func InitRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Opts().RedisAddr,
		Password: config.Opts().RedisPasswd, // no password set
		DB:       0,                         // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		glog.Errorln("db ping fail", err)
	}

	go Redistimer()
}

func Redistimer() {
	redisTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-redisTicker.C:
			_, err := RedisClient.Ping().Result()
			if err != nil {
				glog.Errorln("redis connect fail,err:", err)
				InitRedis()
			}
		}
	}
}
