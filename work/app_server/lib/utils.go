package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang/glog"
	"math/big"
	"time"
)

var Multiple float64 = 100

func MoneyIn(in float64) float64 {
	return Multiple * in
}

func MoneyOut(in float64) float64 {
	return in / Multiple
}

func Decimal(num float64) float64 {
	return float64(int(num*100)) / 100
}

func Decimal20(num float64) float64 {
	return float64(int(num*1000)) / 1000
}

func Int64BigIntArr(in []int64) []*big.Int {
	re := make([]*big.Int, 0)
	for _, v := range in {
		re = append(re, big.NewInt(v))
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

func CipherStr(key, value string) string {
	h := md5.New()
	h.Write([]byte(value + key)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	glog.Infoln("MD5 key = ", key, " value = ", value, cipherStr)
	return hex.EncodeToString(cipherStr)
}

func XOClockTomorrow(hour int) time.Time {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, 0, 0, 0, time.Now().Location()).Add(24 * time.Hour)

	return t
}

func XOClockToDay(hour int) time.Time {
	t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, 0, 0, 0, time.Now().Location())

	return t
}

func Today(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Month(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func SetDate(y, d, m int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Now().Location())
}

func MoneyP(in float64) float64 {
	n := in * 1000
	num := int64(n)
	if num%10 == (9) {
		num += 1
	}
	num /= 10
	fmt.Println(float64(num) / 100)

	return float64(num) / 100
}
