//author xinbing
//time 2018/8/25 15:31
//base64工具,ecoding可通过GenerateBase64Encoder接口生成的encoder来创建
package encrypt

import (
	"encoding/base64"
	"math/rand"
	"time"
)
func Base64Encode(encoding *base64.Encoding, key string) string{
	return encoding.EncodeToString([]byte(key))
}

func Base64Decode(encoding *base64.Encoding, base64Str string) (string,error){
	resultBytes, err := encoding.DecodeString(base64Str)
	return string(resultBytes), err
}

// generate base64 encoder string
func GenerateBase64Encoder() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	raw := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	newRawBytes := make([]byte, len(raw))
	for i:=0; i<len(newRawBytes); i++ {
		index := r.Intn(len(raw))
		newRawBytes[i] = raw[index]
		raw = raw[0:index] + raw[index + 1:]
	}
	newRaw := string(newRawBytes) + "+/"
	return newRaw
}
