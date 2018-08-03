package test

import (
	"testing"
	"fmt"
	"sub_account_service/order_server/utils"
)

func TestGenerateDevelopKey(t *testing.T) {
	fmt.Println(utils.GenerateDevelopKey("20001"))
}

func TestMD5(t *testing.T) {
	fmt.Printf(utils.MD5("20001" + "61f0920ffbb1b4c531d4f13432e6a1db"))
}