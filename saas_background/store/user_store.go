//author xinbing
//time 2018/9/11 14:11
//
package store

import (
	"common-utilities/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func redisUserIdKey(accessToken string) string {
	return "saas:user:"+accessToken
}
func CacheUser(accessToken string, userId uint) {
	db.RedisClient.Set(redisUserIdKey(accessToken), userId, 30 * time.Minute)
}

func RemoveUser(ctx *gin.Context) {
	token,err := ctx.Cookie("access-token")
	if err != nil || token == "" {
		logrus.WithError(err).Errorln("GetCurrUserId get cookie failed!")
		return
	}
	db.RedisClient.Del(redisUserIdKey(token))
}

func GetCurrUserId(ctx *gin.Context) uint {
	fmt.Println(ctx.Request.Cookies())
	token,err := ctx.Cookie("access-token")
	if err != nil || token == "" {
		logrus.WithError(err).Errorln("GetCurrUserId get cookie failed!")
		return 0
	}
	id, err := db.RedisClient.Get(redisUserIdKey(token)).Uint64()
	if err != nil {
		logrus.WithError(err).Errorln("GetCurrUserId get from redis failed!")
		return 0
	}
	retId := uint(id)
	CacheUser(token, retId)
	return retId
}

func CacheUserMenus() {

}