//author xinbing
//time 2018/9/4 15:42
package utilities

import (
	"fmt"
	"testing"
)

func TestGetRandomNumStr(t *testing.T) {
	fmt.Println(GetRandomStr(32))
	fmt.Println(GetRandomNumStr(32))
}
