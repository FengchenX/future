//author xinbing
//time 2018/8/29 9:20
package eft_encrypt

import (
	"encoding/base64"
	"common-utilities/encrypt"
)

const (
	base64Encoder = "Q427EqRMpVIzutgD0v3Jr6cdsyaNUxjeFW1Z9bSfiBPHX8C5TlwkGAKYnLoOhm+/"
)

var transferKitBase64Encoding = base64.NewEncoding(base64Encoder)

func EftEncrypt(content string) string{
	return byteEncrypt(encrypt.Base64Encode(transferKitBase64Encoding,content))
}

func EftDecrypt(content string) (string, error) {
	return encrypt.Base64Decode(transferKitBase64Encoding, byteDecrypt(content))
}
