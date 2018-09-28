//author xinbing
//time 2018/9/4 17:17
package db

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var RedisClient *redis.Client

func InitRedis(redisConfig *RedisConfig) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.RedisAddr,
		Password: redisConfig.RedisPwd, 				 // no password set
		DB:       redisConfig.RedisDB,                         // use default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		logrus.WithField("addr",redisConfig.RedisAddr).Errorln("ping redis error!")
	}
	go redisTimer(redisConfig)
}

func redisTimer(redisConfig *RedisConfig) {
	redisTicker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-redisTicker.C:
			_, err := RedisClient.Ping().Result()
			if err != nil {
				logrus.Errorln("redis connect fail,err:", err)
				InitRedis(redisConfig)
			}
		}
	}
}