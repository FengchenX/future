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

func EncryptUserPwd(pwd string) string{
	pwd = pwd + "此情可待成追忆"
	return MD5(pwd)
}
func GenerateDevelopKey(appId string) string{
	s := MD5(appId)
	ss := s[0:16] + appId + s[16:]
	return MD5(ss)
}
