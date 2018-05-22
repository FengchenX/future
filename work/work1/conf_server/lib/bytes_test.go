package lib

import (
	"fmt"
	"testing"
)

func TestStrToByte32(t *testing.T) {
	bs := StrToByte32("资深厨师就是我---launch")
	fmt.Println(Byte32ToStr(bs))
}
