package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
	"sub_account_service/fabric_server/model"
)

func main() {
	//var ticker=time.NewTicker(time.Millisecond)
	//for _=range ticker.C {
	//	send()
	//}
	send()
}

func send()  {
	var data=model.ReqSetAccount{
		UserKeyStore:"UserKeyStore",
		UserParse:    "UserParse",
		KeyString :   "KeyString",
		UserAccount:  &model.UserAccount{
			Address :  "Address",
			Name    :"Name",
			BankCard : "BankCard",
			WeChat    :"WeChat",
			Alipay  :  "Alipay",
			Telephone: "Telephone",
		},
	}
	fmt.Printf("in:%+v\n",data)
	//压力测试
	send := new(bytes.Buffer)
	json.NewEncoder(send).Encode(&data)
	var request, err =http.NewRequest("POST","http://127.0.0.1:20000/setaccount",send)
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