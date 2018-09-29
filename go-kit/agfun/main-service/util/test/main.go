package main

import (
	"fmt"
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

func main() {
	CacheUser("mykey", 16)
}
func redisUserIdKey(accessToken string) string {
	return accessToken
}

func CacheUser(accessToken string, userId uint) {
	// db.RedisClient.Set(redisUserIdKey(accessToken), userId, 30 * time.Minute)
	_, err := conn.Do("SET", redis.Args{}.Add(redisUserIdKey(accessToken)).Add(int(userId)).Add("EX").Add(int(30*time.Minute))...)
	if err != nil {
		fmt.Println(err)
	}
}