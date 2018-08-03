package utils

import "bytes"

func StrToByte32(str string) [32]byte {
	buf := bytes.NewBufferString(str).Bytes()
	ttl := [32]byte{}
	for i := 0; i < (len(buf)) && i < 32; i++ {
		ttl[i] = buf[i]
	}
	return ttl
}

func Byte32ToStr(bs [32]byte) string {
	bss := make([]byte, 32)
	copy(bss, bs[:])
	str:=string(bss)
	return str
}
