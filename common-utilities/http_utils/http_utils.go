//author xinbing
//time 2018/9/4 13:47
//http请求工具
package http_utils

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)
//FIXME 待调整
func Get(url string, header map[string]string) ([]byte, error){
	return sendHttpRequest("GET", url, nil, header)
}

func Post(url string, body io.Reader, header map[string]string) ([]byte, error){
	return sendHttpRequest("POST", url, body, header)
}

func PostForm(url string, valuePairs, header map[string]string) ([]byte, error) {
	return []byte{}, errors.New("un implements method!")
}

// 待修改
func sendHttpRequest(method, url string, body io.Reader, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	//设置请求头
	setHeaders(req, &header)

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http response status code: %v", resp.StatusCode)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}

func setHeaders(req *http.Request, header *map[string]string) {
	if header == nil {
		req.Header.Add("Content-Type", "application/json")
		return
	}
	xh := *header
	if xh["Content-Type"] == "" {
		req.Header.Add("Content-Type", "application/json")
	}
	for key,value := range xh {
		req.Header.Add(key, value)
	}
}


