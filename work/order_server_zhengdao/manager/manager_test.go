package manager

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"
	"znfz/three_server/lib"
	"znfz/three_server/models"
)

type ts struct {
	lock    sync.RWMutex
	count   int
	sumTime int64
}

var huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"

//var ip string = "39.108.80.66:8888"
var ip string = "120.78.195.103:8888"


//var ip string = "localhost:8888"

func TestManager(t *testing.T) {
	glog.Infoln("start testing")

	//tss := &ts{}
	for i := 0; i < 1; i++ {
		time.Sleep(1 * time.Second)
		go func(num int) {
			num += 90
			glog.Infoln(num * 100)
			rand.Seed(time.Now().Unix())
			orderid := int64(rand.Intn(100000000)) + int64(5)
			req1 := order(orderid, 0.01)
			req2 := bill(orderid, 0.01)
			c := http.Client{}
			c.Do(req1)
			c.Do(req2)
		}(i)
	}
	time.Sleep(5 * time.Second)
}

func order(orderid int64, price float64) *http.Request {
	glog.Infoln("start testing order")

	dishs := make([]*models.Dish, 0)
	dishs = append(dishs, &models.Dish{
		DishName:  "x",
		DishCount: "x",
		DishPrice: "x",
	})


	t := time.Now().Unix()
	oo := &models.Order{
		Time:         t,
		Token:        lib.CipherStr(models.ThreeKey, strconv.Itoa(int(t))),
		OrderId:      orderid,
		CreateTime:   time.Now().String(),
		Company:      "正稻",
		BranchShop:   "201",
		Price:        float64(price),
		OrderContent: []*models.Dish{
			&models.Dish{
				DishName:  "x",
				DishCount: "x",
				DishPrice: "x",
			},
		},
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

func bill(orderid int64, price float64) *http.Request {
	glog.Infoln("start testing bill")

	t := time.Now().Unix()
	bb := &models.Bill{
		Time:       t,
		Token:      lib.CipherStr(models.ThreeKey, strconv.Itoa(int(t))),
		OrderId:    orderid,
		CreateTime: time.Now().String(),
		Company:    "正稻",
		BranchShop: "201",
		PayWay:     "昊信支付",
		Price:      float64(price),
	}
	j, _ := json.Marshal(bb)

	glog.Infoln("OrderId ：", bb.OrderId)

	buf := bytes.NewBuffer(j)
	req, err := http.NewRequest("POST", "http://"+ip+"/savebill", buf)
	if err != nil {
		glog.Errorln("err", err)
	}
	return req
}
