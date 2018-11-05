package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5加密
func Md5Encrypt(data string) (string, error) {
	// md5 init
	md5Ctx := md5.New()
	n, err := md5Ctx.Write([]byte(data))
	if n <= 0 || err != nil {
		return "", err
	}
	// hex_digest
	return hex.EncodeToString(md5Ctx.Sum(nil)), nil
}
