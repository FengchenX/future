package lib

import (
	"crypto/md5"
	"github.com/golang/glog"
	"encoding/hex"
)

func CipherStr(key, value string) string {
	h := md5.New()
	h.Write([]byte(value + key)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	glog.Infoln("MD5 key = ", key, " value = ", value, cipherStr)
	return  hex.EncodeToString(cipherStr)
}
