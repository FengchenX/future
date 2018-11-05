package office

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"grm-service/util"
)

type OnlyOffice struct {
	Endpoints string
	UserName  string
	Password  string

	token string
}

func (office *OnlyOffice) Login() error {
	url := fmt.Sprintf(`%s/api/2.0/authentication.json`, office.Endpoints)

	req := login{
		UserName: office.UserName,
		Password: office.Password,
	}
	args, err := json.Marshal(req)
	if err != nil {
		return err
	}
	ret, err := util.HttpPost(url, args)
	if err != nil {
		return err
	}
	var res loginRes
	if err := json.Unmarshal([]byte(ret), &res); err != nil {
		return err
	}
	office.token = res.Response.Token
	return nil
}

func (office *OnlyOffice) Upload(path, id string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	body := bytes.NewBufferString("")
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", id+"_"+filepath.Base(path))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}
	file.Close()
	// request
	url := fmt.Sprintf(`%s/api/2.0/files/@my/upload`, office.Endpoints)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, io.MultiReader(body))
	if err != nil {
		return "", err
	}
	content := writer.FormDataContentType()
	writer.Close()
	req.Header.Set("Content-Type", content)
	req.Header.Set("Authorization", office.token)
	req.Header.Set("Connection", "Keep-Alive")
	req.ContentLength = int64(body.Len())
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := uploadRes{}
	if err := json.Unmarshal(resBody, &res); err != nil {
		return "", err
	}
	return strconv.Itoa(res.Response.Id), nil
}
func (office *OnlyOffice) Share(id string) (string, error) {
	url := fmt.Sprintf(`%s/api/2.0/files/%s/sharedlink`, office.Endpoints, id)
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(`{"share":"Read"}`))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", office.token)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := shareRes{}
	if err := json.Unmarshal(resBody, &res); err != nil {
		return "", err
	}
	return res.Response, nil
}
