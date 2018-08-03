package test

import (
	"testing"
	"net/http"
	"sub_account_service/blockchain_server/model"
	"strconv"
	"sync"
	"time"
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

func TestThreeBill(t *testing.T) {
	var wg = sync.WaitGroup{}
	url := "http://localhost:8082/threesetbill"
	count := 128
	beginTime := time.Now().UnixNano()
	for i := 123;i  < count; i++ {
		wg.Add(1)
		go func(z int) {
			defer wg.Done()
			req := model.ReqThreeSetBill{
				UserAddress: "0x56a58d378fd5647de22bf10007ab2f49e47d83b7",
				UserKeyStore:"",
				UserParse:"",
				KeyString:`{"address":"56a58d378fd5647de22bf10007ab2f49e47d83b7","privatekey":"a2debe0a48a45a926850fe6df063365258c965f0ef09fe3746038d92f1b61b01","id":"44e95601-9851-422e-ae00-c0a481436fe6","version":3}`,
				ScheduleName:"AC:a171cc9d89c8359cf794fc0b95bc417d",// string  `protobuf:"bytes,3,opt,name=ScheduleName" json:"ScheduleName,omitempty"` //排班编号
				Money:13.3,        //float64 `protobuf:"fixed64,5,opt,name=Money" json:"Money,omitempty"`
				OrderId: "orderAbc121212" + strconv.Itoa(z),     //string  `protobuf:"bytes,6,opt,name=OrderId" json:"OrderId,omitempty"`
			}
			out, _ := json.Marshal(req)
			resp,err := http.Post(url,"application/json", bytes.NewReader(out))
			if err != nil {
				fmt.Println("post error",z,err)
			}else {
				out,err = ioutil.ReadAll(resp.Body)
				fmt.Println("post success",string(out))
			}
		}(i)
	}
	wg.Wait()
	endTime := time.Now().UnixNano()
	diffTime := (endTime - beginTime)/int64(time.Millisecond)
	fmt.Printf("count:%v,duration:%v",count,diffTime)
}
