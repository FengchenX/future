package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"sub_account_service/finance/config"
)

/*
	ipcDir 是要连接这个节点的geth.ipc目录，如：/root/work/launch_pf/src/ethbase/geth.ipc
*/
func GetEthclient(ipcDir string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(ipcDir)
	return conn, err
}

// todo 这个方法
func ParseKeyStore() string {
	accountKey := config.Opts().FinanceKeyStore
	passphrase := config.Opts().FinancePhrase
	data, err := ioutil.ReadAll(strings.NewReader(accountKey))
	if err != nil {
		return ""
	}
	k, _ := keystore.DecryptKey(data, passphrase)
	if err != nil {
		return ""
	}
	b, err := json.Marshal(k)
	if err != nil {
		return ""
	}
	return string(b)
}

// bool to int64(1,0)
func BoolToInt(data bool) int64 {
	if data {
		return 1
	}
	return 0
}

// bool to int64(1,0)
func Uint8ToBool(data uint8) bool {
	if data > 0 {
		return true
	}
	return false
}

// 保留两位小数的浮点数字符串
func KeepTwoDecimalsOfStr(s string) (float64, error) {
	sTmp := strings.Split(s, ".")
	if len(sTmp) != 2 {
		return 0, fmt.Errorf("it's not float string")
	}

	out, err := strconv.ParseFloat(sTmp[0]+"."+sTmp[1][0:2], 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}

// 发送http请求
func SendHttpRequest(method, url string, body io.Reader, head, fdata map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http response status code: %v", resp.StatusCode)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}
