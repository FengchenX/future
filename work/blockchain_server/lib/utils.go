package lib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang/glog"
	"math/big"
)

const (
	Multiple = 10000
	Hun = 100
)

func MoneyIn(in float64) float64 {
	return in * Multiple
}

func MoneyOut(in float64) float64 {
	n := in / 10
	num := int64(n)
	if num%10 == (9) {
		num += 1
	}
	num /= 10
	fmt.Println(float64(num) / 100)

	return float64(num) / 100
}

func RadioIn(in float64) float64 {
	return Hun * in
}

func RadioOut(in float64) float64 {
	in *= 10
	if int(in)%10 == 9 {
		in += 1
	}
	return in / 1000
}


func Int64BigIntArr(in []int64) []*big.Int {
	re := make([]*big.Int, 0)
	for _, v := range in {
		re = append(re, big.NewInt(v))
	}
	return re
}

func Float64BigIntArr(in []float64) []*big.Int {
	re := make([]*big.Int, 0)
	for _, v := range in {
		re = append(re, big.NewInt(int64(v)))
	}
	return re
}

func BigIntInt64Arr(in []*big.Int) []int64 {
	re := make([]int64, 0)
	for _, v := range in {
		if v != nil {
			re = append(re, v.Int64())
		}
	}
	return re
}

const SYNBOL = "LIANGSIHAO"

func CipherStr(key, value string) string {
	h := md5.New()
	h.Write([]byte(value + key + SYNBOL)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	glog.Infoln("MD5 key = ", key, " value = ", value, cipherStr)
	return hex.EncodeToString(cipherStr)
}

func ParseStrArr(in [][32]byte) []string {
	re := make([]string, 0)
	for _, v := range in {
		re = append(re, Byte32ToStr(v))
	}
	return re
}

func ReParseStrArr(in []string) [][32]byte {
	re := make([][32]byte, 0)
	for _, v := range in {
		buf := bytes.NewBufferString(v).Bytes()
		ttl := [32]byte{}
		for i := 0; i < (len(buf)) && i < 32; i++ {
			ttl[i] = buf[i]
		}
		re = append(re, ttl)
	}

	return re
}
