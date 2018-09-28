//author xinbing
//time 2018/8/25 15:31
package encrypt

import (
	"testing"
	"fmt"
	"encoding/base64"
)

func TestDecodeKey(t *testing.T) {
	for i := 0; i<10; i++ {
		base64Encoding := base64.NewEncoding(GenerateBase64Encoder())
		privateKey, publicKey, _ := GenKeyPairs(2048)
		encode := Base64Encode(base64Encoding,privateKey)
		decode, _ := Base64Decode(base64Encoding,encode)
		flag1 := decode == privateKey
		encode2 := Base64Encode(base64Encoding,publicKey)
		decode2, _ := Base64Decode(base64Encoding,encode2)
		flag2 := decode2 == publicKey
		if !flag1 || !flag2 {
			fmt.Println("failed")
		}
	}
}


func TestGenBase64Salt(t *testing.T){
	fmt.Println(GenerateBase64Encoder())
}