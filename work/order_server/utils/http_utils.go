package utils

import (
	"io"
	"net/http"
	"fmt"
	"io/ioutil"
)

func Get(url string,body io.Reader,head,fdata map[string]string)([]byte,error) {
	return sendHttpRequest("GET",url,body,head,fdata)
}
func Post(url string,body io.Reader,head,fdata map[string]string)([]byte,error) {
	return sendHttpRequest("POST",url,body,head,fdata)
}
func sendHttpRequest(method, url string, body io.Reader, head, fdata map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

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
