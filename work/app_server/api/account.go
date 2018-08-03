package api

import (
	"fmt"
	"github.com/golang/glog"
	"sub_account_service/app_server_v2/lib"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/model"
)

//SetAccount 设置账户
func SetAccount(c *gin.Context) {
	var req model.ReqSetAccount
	var resp model.RespSetAccount
	if err := lib.ParseReq(c, "SetAccount", &req); err != nil {
		return
	}
	glog.Infoln("SetAccount**********************req: ", req)

	keyString := lib.KeyMap.Calc(req.UserKeyStore, req.UserParse)
	req.KeyString = keyString

	err := doPost("/setaccount", req, &resp)
	if err != nil {
		glog.Errorln("SetAccount*******************doPost err", err)
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	glog.Infoln("SetAccount******************resp: ", resp)
	c.JSON(http.StatusOK, resp)
}

//GetAccount 获取账户
func GetAccount(c *gin.Context) {
	var req model.ReqGetAccount
	var resp model.RespGetAccount
	mediator(&req, &resp, c, "GetAccount", "/getaccount")
}

//doPost 发送Post到api
func doPost(url string, reqobj interface{}, respobj interface{}) error {
	url = "http://" + config.ConfInst().ApiAddress + config.ConfInst().ApiPort + url
	fmt.Println("URL:>", url)

	buf, err := json.Marshal(reqobj)
	if err != nil {
		fmt.Println("doPost********************json.Marshal err", err)
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("doPost********************client.Do err", err)
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if err != nil {
		fmt.Println("doPost*************************ioutil.ReadAll", err)
		return err
	}

	if err = json.Unmarshal(body, respobj); err != nil {
		fmt.Println("doPost**********************json.UnmMarshal", err)
		return err
	}
	return nil
}
