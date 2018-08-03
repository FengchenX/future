package lib

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"strings"
)

/*
	ipcDir 是要连接这个节点的geth.ipc目录，如：/root/work/launch_pf/src/ethbase/geth.ipc
*/
func GetEthclient(ipcDir string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(ipcDir)
	return conn, err
}

// todo 这个方法
func ParseKeyStore(accountKey, passphrase string) string {
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
