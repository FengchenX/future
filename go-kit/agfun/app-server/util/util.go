package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
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

//MIN 自定义比较精度
const MIN = 0.0001

//IsEqual 两个浮点型是否相等
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}
