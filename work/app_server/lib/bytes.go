package lib

import (
	"bytes"
)

func StrToByte32(str string) [32]byte {
	buf := bytes.NewBufferString(str).Bytes()
	ttl := [32]byte{}
	for i := 0; i < (len(buf)) && i < 32; i++ {
		ttl[i] = buf[i]
	}
	return ttl
}

func StrToByte20(str string) [20]byte {
	str = string(str)
	buf := bytes.NewBufferString(str).Bytes()
	ttl := [20]byte{}
	for i := 0; i < (len(buf)) && i < 20; i++ {
		ttl[i] = buf[i]
	}
	return ttl
}

func Byte32ToStr(bs [32]byte) string {
	bss := make([]byte, 32)
	rebs := make([]byte, 0)

	copy(bss, bs[:])
	for _, v := range bss {
		if v != 0 {
			rebs = append(rebs, v)
		}
	}
	str := string(rebs)
	return str
}

func Byte20ToStr(bs [20]byte) string {
	bss := make([]byte, 32)
	rebs := make([]byte, 0)

	copy(bss, bs[:])
	for _, v := range bss {
		if v != 0 {
			rebs = append(rebs, v)
		}
	}
	str := string(rebs)
	return str
}
