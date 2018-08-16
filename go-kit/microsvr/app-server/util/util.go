package util

import (
	"io/ioutil"
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"

	"crypto/md5"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/sirupsen/logrus"
)


//DoPost 发送json post 对象
func DoPost(url string, reqobj interface{}, respobj interface{}) error {
	fmt.Println("URL:>", url)

	buf, err := json.Marshal(reqobj)
	if err != nil {
		fmt.Println("DoPost*******************json.Marshal err", err)
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("DoPost********************client.Do err", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if err != nil {
		fmt.Println("DoPost*******************ioutil.ReadAll", err)
		return err
	}

	if err = json.Unmarshal(body, respobj); err != nil {
		fmt.Println("DoPost****************json.UnMarshal err", err)
		return err
	}
	return nil
}

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
	logrus.Infoln("MD5 key = ", key, " value = ", value, cipherStr)
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