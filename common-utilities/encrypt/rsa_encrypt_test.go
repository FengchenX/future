//author xinbing
//time 2018/8/25 15:31
package encrypt

import (
	"testing"
	"fmt"
)

func TestRsaEncrypt(t *testing.T) {
	privateKey,publicKey,_ := GenKeyPairs(2048)
	b,_ := RsaEncrypt("三分归元气", publicKey)
	fmt.Println(string(b))
	b,_ = RsaDecrypt(string(b),privateKey)
	fmt.Println(string(b))
}

func TestGenRsaKey(t *testing.T) {
	fmt.Println(GenKeyPairs(2048))
}
