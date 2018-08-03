package test

import (
	"testing"
	"sync"
	"fmt"
	"time"
	"sub_account_service/order_server/entity"
	"sub_account_service/order_server/dto"
	"encoding/json"
	"bytes"
	"sub_account_service/order_server/utils"
)

func TestAddOrder(t *testing.T) {
	var wg sync.WaitGroup
	timeStamp := fmt.Sprintf("%v",time.Now().Unix() * 1000)
	developKey := "a9cd0f73a013f6f2704c9d0fcc72ab8f"
	sign := utils.MD5(timeStamp + developKey)
	url := "http://localhost:7777/orders/save?Sign="+sign +"&AppId=9527&Timestamp="+timeStamp
	beginTime := time.Now().UnixNano()
	count := 1
	for i :=0;i<count;i++ {
		wg.Add(1)
		go func(z int64) {
			defer wg.Done()
			orderState := entity.Normal
			thirdNo := "third" + fmt.Sprintf("%v",time.Now().Unix() * 1000 + z)
			orderNo := fmt.Sprintf("%v",time.Now().Unix() * 1000 + z)
			bill := dto.BillOrderSaveDTO{
				ThirdTradeNo: thirdNo ,
				OrderNo: orderNo,
				SubNumber: "AC:2acea93ab7bfff66e492c4169bf615b1",//AC:aee8a4c271a2572623ed3294706c826a
				Company: "测试",
				BranchShop: "测试公司",
				OrderTime: "2018-12-12 11:11:11",
				PayWay: "支付宝",
				Discount: 300,
				SumPrice: 900,
				Price: 600,
				OrderState:orderState,
				Remarks:"1212",
			}
			out, _ := json.Marshal(bill)
			tt,err := utils.Post(url, bytes.NewReader(out),nil,nil)
			time.Sleep(time.Second)
			fmt.Println(string(tt),err)
		}(int64(i))
	}
	wg.Wait()
	endTime := time.Now().UnixNano()
	diffTime := (endTime - beginTime)/int64(time.Millisecond)
	fmt.Printf("count:%v,duration:%v",count,diffTime)
}
