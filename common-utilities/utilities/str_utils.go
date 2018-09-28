//author xinbing
//time 2018/8/28 14:18
//字符串工具
package utilities

import (
	"math/rand"
	"time"
)

var randomStrSource = []byte("0123456789abcdefghijklmnopqrstuvwxyz")
//获取随机字符串
func GetRandomStr(length int) string {
	result := make([]byte,length)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63())) //增大随机性
	for i:=0;i<length;i++ {
		result[i] = randomStrSource[r.Intn(len(randomStrSource))]
	}
	return string(result)
}

//生成纯数字的随机字符串
func GetRandomNumStr(length int) string {
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63())) //增大随机性
	for i:=0;i<length;i++ {
		result[i] = byte('0' + r.Intn(10)) //0 - 9
	}
	return string(result)
}
