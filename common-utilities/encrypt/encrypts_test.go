//author xinbing
//time 2018/9/5 14:25
package encrypt

import (
	"fmt"
	"testing"
)

func TestSHA1(t *testing.T) {
	str := "1"
	fmt.Println(MD5(str))
	fmt.Println(SHA1(str))
	fmt.Println(SHA256(str))
	fmt.Println(SHA512(str))
}
