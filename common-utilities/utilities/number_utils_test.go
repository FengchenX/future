//author xinbing
//time 2018/8/28 14:21
package utilities

import (
	"testing"
	"fmt"
	"math"
	"math/rand"
	"time"
	"strconv"
)

func TestRound(t *testing.T) {
	fmt.Println(Round(3.1334, 2))
	fmt.Println(Floor(3.9456, 2))
	fmt.Println(Ceil(3.123, 2))
	fmt.Println(Round(3.9-0.00001, 6))
	fmt.Println(Round(0.999 /10000000,10))
	fmt.Println(math.Floor(3.123))
}

func TestRound2(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for  {
		f1 := r.NormFloat64()
		f1 = math.Abs(f1 / 100.0)
		f1 =  Round(f1 , 6)
		s := strconv.FormatFloat(f1,'f', -1, 64)
		if len(s) > 8 {
			fmt.Println(s)
			break
		}
	}
}

func TestRound3(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		f1 := r.NormFloat64()
		f1 = math.Abs(f1 / 100.0)
		s := strconv.FormatFloat(f1, 'f', -1, 64)
		//ret, _ := strconv.ParseFloat(s, 64)
		if len(s) > 7 {
			fmt.Println(s)
			break
		}
	}
}

func TestCompare(t *testing.T) {
	fmt.Println(CompareWithScale(0.11,0.11,2))
}

