package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func getAesKey() []byte {
	strKey := "!@#$%^&*()123456"
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("res key 长度不能小于16")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		//取前32个字节
		return arrKey[:32]
	}
	if keyLen >= 24 {
		//取前24个字节
		return arrKey[:24]
	}
	//取前16个字节
	return arrKey[:16]
}

//加密字符串
func AesEncrypt(strMesg string) (string, error) {
	key := getAesKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))

	str := base64.StdEncoding.EncodeToString(encrypted)
	return str, nil
}

//解密字符串
func AesDecrypt(pwd string) (strDesc string, err error) {
	defer func() {
		//错误处理
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	pwdData, err := base64.StdEncoding.DecodeString(pwd)
	key := getAesKey()
	var iv = []byte(key)[:aes.BlockSize]
	decrypted := make([]byte, len(pwdData))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(decrypted, pwdData)
	return string(decrypted), nil
}
