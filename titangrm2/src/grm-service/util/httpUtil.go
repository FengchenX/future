package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	stdlog "log"
	"mime/multipart"
	"net"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"sort"
	"strings"
)

func GetUrlArgs(args string, r *http.Request) string {
	args_value := r.URL.Query().Get(args)
	if len(args_value) <= 0 {
		args_value = r.URL.Query().Get(strings.ToUpper(args))
		if len(args_value) <= 0 {
			args_value = r.URL.Query().Get(strings.ToLower(args))
		}
	}
	return args_value
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func HttpPost(url string, data []byte) (string, error) {
	body := bytes.NewBuffer(data)
	res, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpGet(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthGet(url string, username, pwd string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthDelete(url string, username, pwd string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthXmlPost(url string, data []byte, username, pwd string) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthPost(url string, data []byte, username, pwd string) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthPutByContentType(url string, data []byte, username, pwd, ct string) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Accept", "*/*")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthGetByContentType(url string, username, pwd, ct, accept string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Accept", accept)
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthPostByContentType(url string, data []byte, username, pwd, ct string) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", ct)
	req.Header.Set("Accept", "*/*")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func HttpAuthPut(url string, data []byte, username, pwd string) (string, error) {
	client := &http.Client{}
	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, pwd)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}

func UploadFile(targetUrl, filename, uploadPath string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer:", err)
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("error io.Copy:", err)
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	uploadPath = url.QueryEscape(uploadPath)
	targetUrl = targetUrl + "?path=" + uploadPath
	_, err = http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("error http.Post:", err)
		return err
	}
	return err
}

func CheckRequestGet(res http.ResponseWriter, req *http.Request, get string) (string, bool) {
	_get := req.URL.Query().Get(get)
	if len(_get) == 0 {
		res.WriteHeader(500)
		res.Write([]byte(get + " is null"))
		stdlog.Println(get + " is null")
		return "", false
	}
	return _get, true
}

func GetInternalIP4s() []string {
	var ip4s []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		stdlog.Println("Oops:", err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsGlobalUnicast() {
			if ipnet.IP.To4() != nil {
				ip4s = append(ip4s, ipnet.IP.String())
			}
		}
	}
	return ip4s
}

func GetHttpAddress1(listen string, defaultPort string) []string {
	var address []string
	s := strings.Split(listen, ",")
	for _, value := range s {
		httpUrl := value
		if strings.HasPrefix(httpUrl, ":") {
			for _, ip := range GetInternalIP4s() {
				httpUrl = "http://" + ip + listen
				address = append(address, httpUrl)
			}
		} else if !strings.HasPrefix(value, "http://") {
			httpUrl = "http://" + value
		} else if !strings.Contains(value, ":") {
			httpUrl = "http://" + value + defaultPort
		}
		address = append(address, httpUrl)
	}
	sort.Strings(address)
	return address
}

func GetHttpAddress(listen string, defaultPort string) []string {
	var address []string
	s := strings.Split(listen, ",")
	for _, value := range s {
		if !strings.Contains(value, ":") {
			value = value + defaultPort
		}
		address = append(address, value)
	}
	return address
}

func GetSingelAddress(listen string, defaultPort string) string {
	if !strings.Contains(listen, ":") {
		listen = listen + defaultPort
	}
	return listen
}
