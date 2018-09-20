package test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestFilePath(t *testing.T) {
	var path = "/root/sub_account_service/order_server/start"
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Base(path))
	var str = "tttlog"
	str = str[0:strings.LastIndex(str, ".")] + "20180806" + str[strings.LastIndex(str, "."):]
	fmt.Println(str)
}
