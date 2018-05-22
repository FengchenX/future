package manager

import (
	"bytes"
	"github.com/golang/glog"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
	"znfz/web_server/models"
	"encoding/json"
)

type ts struct {
	lock    sync.RWMutex
	count   int
	sumTime int64
}

var huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"

var ip string = "39.108.80.66:8888"

//var ip string = "192.168.83.200:8888"

func TestManager(t *testing.T) {
	glog.Infoln("start testing")

	//req := bind()
	//req := getbind()
	//req := bill()

	//tss := &ts{}
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		go func(num int) {
			orderid := int64(rand.Intn(100000000)) + int64(5)
			req := bill(orderid)
			c := http.Client{}
			c.Do(req)
		}(i)
	}
	time.Sleep(10 * time.Second)
}

func apply() *http.Request {
	glog.Infoln("start testing apply")
	j, _ := json.Marshal(&ReqApplyCompany{
		//	CompanyName:      "正稻",
		//	CompanyBanchName: "正稻001",
		CompanyName: "易迈",
		BranchName:  "测试店",
		UserAddress: "0x759D4c2E15587Fae036f183202F36CA3C667ccbD",
		PassWord:    "15219438281",
		Phone:       "15219438281",
	})
	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://"+ip+"/setapplycompany", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func bind() *http.Request {
	glog.Infoln("start testing bind")
	j, _ := json.Marshal(&ReqBinder{
		CompanyName: "易迈",
		BranchName:  "测试店",
		UserAddress: "0x759D4c2E15587Fae036f183202F36CA3C667ccbD",
		Phone:       "15219438281",
	})
	buf := bytes.NewBuffer(j)
	url := "http://" + ip + "/setbindmsg"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func order(orderid int64) *http.Request {
	//glog.Infoln("start testing order")

	dishs := make([]*models.Dish, 0)
	dishs = append(dishs, &models.Dish{
		DishName:  "麻辣火锅",
		DishCount: "2",
		DishPrice: "42",
	})

	//dishs = append(dishs, &models.Dish{
	//	DishName:  "大鸡腿",
	//	DishCount: "2",
	//	DishPrice: "25",
	//})

	oo := &models.Order{
		OrderId:      orderid,
		CreateTime:   time.Now().String(),
		Company:      "麦当基",
		BranchShop:   "南山分店",
		Price:        float64(50),
		OrderContent: dishs,
	}

	j, _ := json.Marshal(oo)

	buf := bytes.NewBuffer(j)
	url := "http://" + ip + "/saveorder"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}

func bill(orderid int64) *http.Request {
	glog.Infoln("start testing bill")
	bb := &models.Bill{
		OrderId:    orderid,
		CreateTime: time.Now().String(),
		Company:    "麦当基",
		BranchShop: "南山分店",
		PayWay:     "昊信支付",
		Price:      float64(500),
	}
	j, _ := json.Marshal(bb)

	glog.Infoln("OrderId ：", bb.OrderId)

	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://192.168.83.200:8888/savebill", buf)
	//req, err := http.NewRequest("POST", "http://39.108.80.66:8888/savebill", buf)
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
