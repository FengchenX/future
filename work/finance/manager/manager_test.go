package manager

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/golang/glog"
)

type ts struct {
	lock    sync.RWMutex
	count   int
	sumTime int64
}

var huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"

//var ip string = "39.108.80.66:7777"
//var ip string = "localhost:1111"
var ip string = "120.78.195.103:7777"

func TestManager(t *testing.T) {
	glog.Infoln("start testing")

	for i := 0; i < 1; i++ {
		time.Sleep(1 * time.Second)
		go func(num int) {
			req := apply()
			c := http.Client{}
			resp, _ := c.Do(req)
			glog.Infoln(resp)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

func apply() *http.Request {
	glog.Infoln("start testing apply")
	j, _ := json.Marshal(&ReqApplyCompany{
		CompanyName: "正稻",
		BranchName:  "301",
		UserAddress: "569ef95c3c40d7bfadf53d02edda3afe0c9bb17a",
		PassWord:    "123456",
		Phone:       "13148894184",
	})
	buf := bytes.NewBuffer(j)
	glog.Infoln(string(buf.Bytes()))

	url := "http://" + ip + "/setapplycompany"
	glog.Infoln(url)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func bind() *http.Request {
	glog.Infoln("start testing bind")
	j, _ := json.Marshal(&ReqBinder{
		//CompanyName: "正稻",
		//BranchName:  "201",
		//UserAddress: "0x759D4c2E15587Fae036f183202F36CA3C667ccbD",
		//Phone:       "111111",
		CompanyName: "公司名",
		BranchName:  "分店编号（在第三方方平台获取）",
		UserAddress: "申请者的地址",
		Phone:       "申请者的电话号码",
	})
	buf := bytes.NewBuffer(j)
	glog.Infoln(string(buf.Bytes()))
	url := "http://" + ip + "/setbindmsg"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func getapply() *http.Request {
	glog.Infoln("start testing bind")
	j, _ := json.Marshal(&GetApplyCompany{
		//CompanyName: "正稻",
		//BranchName:  "201",
		//UserAddress: "0x759D4c2E15587Fae036f183202F36CA3C667ccbD",
		//Phone:       "111111",
		//UserAddress: "申请者的地址",
	})
	buf := bytes.NewBuffer(j)
	glog.Infoln(string(buf.Bytes()))
	url := "http://" + ip + "/getapplycompany"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func getbind() *http.Request {
	glog.Infoln("start testing getbind")
	j, _ := json.Marshal(&ReqBindMsg{
		UserAddress: "0x1111",
		Phone:       "6666",
	})
	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://localhost:8088/getbindmsg", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}
