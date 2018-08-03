package payment

import (
	"fmt"
	"strconv"
	"sub_account_service/finance/payment/alipayModels"
	"testing"
	"time"
)

func TestTransferToAccount(t *testing.T) {
	resp, err := TransferToAccount(alipayModels.TransferToAccountReq{
		OrderId:       strconv.Itoa(int(time.Now().UnixNano())),
		Amount:        0.1,
		PayeeAccount:  "18520227050",
		PayerShowName: "张家集",
		Remark:        "烦死了发动机可",
	})
	fmt.Println(resp)
	fmt.Println(err)
}

func TestQueryAliTrade(t *testing.T) {
	resp, err := QueryAliTrade("2018060521001004600522607607")
	fmt.Println(resp)
	fmt.Println(err)
}

func TestQueryTransferToAccount(t *testing.T) {
	resp, err := QueryTransferToAccount("2018060521001004600522607607")
	fmt.Println(resp)
	fmt.Println(err)
}
