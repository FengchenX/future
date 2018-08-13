package util

import (
	"io/ioutil"
	"bytes"
	"fmt"
	"encoding/json"
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