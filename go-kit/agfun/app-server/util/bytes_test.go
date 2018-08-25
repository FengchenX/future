package util

import (
	"fmt"
	"math"
	"testing"
)

func TestStrToByte32(t *testing.T) {
	bs := StrToByte32("资深厨师就是我---launch")
	fmt.Println(Byte32ToStr(bs))
}

func TestDecimal(t *testing.T) {
	fmt.Println(float64(10) / float64(3))
	fmt.Println(Decimal(float64(10) / float64(3)))

	NumTest()
}

func NumTest() {
	a := float64(12)
	fmt.Println("print", Round(a, 5)-Round(float64(3.3), 5)-Round(float64(3.3), 5)-Round(float64(3.96), 5))
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
