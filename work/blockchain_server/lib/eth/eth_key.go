package myeth

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
	"io/ioutil"
	"runtime"
	"strings"
)

/*
	1.accountKey 对应的是要获取的账户地址的私钥信息，在节点的keystore下面。
	2.passphrase是当前操作的账户在创建账户是输入的密码（字符串）。
*/

func ParseEthAuth(key_string string, conn *ethclient.Client, hasNonce bool) (*bind.TransactOpts, error) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 1024)
			glog.Infoln("panic:", runtime.Stack(buf, true), r)
		}
	}()

	var mp map[string]interface{}
	err := json.Unmarshal(bytes.NewBufferString(string(key_string)).Bytes(), &mp)
	if err != nil {
		glog.Infoln("ParseEthAuth err", err)
		return nil, err
	}

	//addr := ""
	if _, exist := mp["address"]; !exist {
		return nil, errors.New("don't have address")
	} else {
		//addr = v.(string)
	}

	if _, exist := mp["id"]; !exist {
		return nil, errors.New("don't have id")
	}

	if _, exist := mp["privatekey"]; !exist {
		return nil, errors.New("don't have privatekey")
	}

	key := &keystore.Key{}
	err = json.Unmarshal(bytes.NewBufferString(key_string).Bytes(), key)
	if err != nil {
		glog.Infoln(err)
		return nil, err
	} else {
	}
	auth := bind.NewKeyedTransactor(key.PrivateKey)

	if hasNonce {
		nonce := getNonce(auth,conn)
		auth.Nonce = &nonce
		//auth.Nonce = big.NewInt(NonceMap.calc(addr, auth, conn))
		glog.Infoln("ParseEthAuth Nonce Calc After:", auth.Nonce, "from:", auth.From.String())
	}

	return auth, nil
}

// build eth key store by using keystore and keyparse
// the func DecryptKey is very fucking slow
func ParseKeyStore(accountKey string, passphrase string) *keystore.Key {
	json, err := ioutil.ReadAll(strings.NewReader(accountKey))
	if err != nil {
		return nil
	}
	k, _ := keystore.DecryptKey(json, passphrase)
	return k
}

func ParseKeyStoreToString(userKeyStore string, passphrase string) string {
	data, err := ioutil.ReadAll(strings.NewReader(userKeyStore))
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
