//author xinbing
//time 2018/8/29 17:02
//加密方法工具集
package encrypt

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

func MD5(content string) string {
	return encrypt(md5.New(), content)
}

func SHA1(content string) string {
	return encrypt(sha1.New(), content)
}

func SHA256(content string) string {
	return encrypt(sha256.New(), content)
}

func SHA512(content string) string {
	return encrypt(sha512.New(), content)
}

func encrypt(h hash.Hash,content string) string {
	h.Write([]byte(content))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}