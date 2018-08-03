package main

import (
	"strconv"
	"time"
	"sub_account_service/number_server/models"
	"math/rand"
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//var ticker=time.NewTicker(time.Millisecond)
	//for _=range ticker.C {
	//	send()
	//}
	send()
}

func send()  {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	no:=timestamp[:12]
	var data=models.OrderServer{
		ThirdTradeNo:no,
		OrderNo:no,
		SubAccountNo:no,
		Company:"正稻",
		MchId:"10001",
		BranchShop:"zhendao",
		OrderType:0,
		PaymentType:0,
		TransferAmount:rand.Float64()*10,
		OrderTime:time.Now(),
		OrderState:1,
		AutoTransfer:0,
	}
	fmt.Printf("in:%+v\n",data)
	//压力测试
	send := new(bytes.Buffer)
	json.NewEncoder(send).Encode(&data)
	var request, err =http.NewRequest("POST","httpServ://39.108.80.66:9897/addOrder",send)
	if err!=nil{
		fmt.Println(err)
	}
	request.Header.Set("Content-Type","application/json")
	http_client := &http.Client{}
	response, err := http_client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("out:%s\n",body)
}