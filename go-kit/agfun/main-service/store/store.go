package store

import (
	// "common-utilities/db"
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"time"
)

var conn redis.Conn

func init() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		panic(err)
	}
	conn = c
	// defer c.Close()
}

func redisUserIdKey(accessToken string) string {
	return "saas:user:" + accessToken
}
func CacheUser(accessToken string, userId uint) {
	// db.RedisClient.Set(redisUserIdKey(accessToken), userId, 30 * time.Minute)
	_, err := conn.Do("SET", redis.Args{}.Add(redisUserIdKey(accessToken)).Add(int(userId)).Add("EX").Add(int(30*time.Minute))...)
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveUser(ctx *gin.Context) {
	// token,err := ctx.Cookie("access-token")
	// if err != nil || token == "" {
	// 	logrus.WithError(err).Errorln("GetCurrUserId get cookie failed!")
	// 	return
	// }
	// db.RedisClient.Del(redisUserIdKey(token))
}

func GetCurrUserId(ctx *gin.Context) uint {
	// fmt.Println(ctx.Request.Cookies())
	// token,err := ctx.Cookie("access-token")
	// if err != nil || token == "" {
	// 	logrus.WithError(err).Errorln("GetCurrUserId get cookie failed!")
	// 	return 0
	// }
	// id, err := db.RedisClient.Get(redisUserIdKey(token)).Uint64()
	// if err != nil {
	// 	logrus.WithError(err).Errorln("GetCurrUserId get from redis failed!")
	// 	return 0
	// }
	// retId := uint(id)
	// CacheUser(token, retId)
	// return retId
	panic("todo")
}

func GetUserId(token string) uint {
	id, err := redis.Int(conn.Do("GET", redisUserIdKey(token)))
	if err != nil {
		return 0
	}
	CacheUser(token, uint(id))
	return uint(id)
}

func CacheUserMenus() {

}
