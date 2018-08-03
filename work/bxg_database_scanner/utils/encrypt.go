package utils

import (
	"crypto/md5"
	"encoding/hex"
)
func MD5(value string) string {
	h := md5.New()
	h.Write([]byte(value)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return  hex.EncodeToString(cipherStr)
}
