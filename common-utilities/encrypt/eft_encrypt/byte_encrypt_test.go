//author xinbing
//time 2018/8/27 14:30
package eft_encrypt

import (
	"testing"
	"fmt"
	"common-utilities/encrypt"
)

func TestByteEncrypt(t *testing.T) {
	tt := "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"
	newTT := byteEncrypt(tt)
	decode := byteDecrypt(newTT)
	fmt.Println(newTT)
	fmt.Println(newTT == "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")
	fmt.Println(decode == tt)
}
func TestByteEncrypt2(t *testing.T) {
	for i:=0; i<100;i++ {
		privateKey,_,_ := encrypt.GenKeyPairs(2048)
		newPrivateKey := byteEncrypt(privateKey)
		decodedPrivateKey := byteDecrypt(newPrivateKey)
		if decodedPrivateKey != privateKey {
			fmt.Println("xxxxxxxxxxxxxxxx")
		}
	}
}

