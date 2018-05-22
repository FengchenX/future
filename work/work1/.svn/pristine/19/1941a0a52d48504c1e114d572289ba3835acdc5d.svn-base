package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"strings"
)

/*
	ipcDir 是要连接这个节点的geth.ipc目录，如：/root/work/launch_pf/src/ethbase/geth.ipc
*/
func GetEthclient(ipcDir string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(ipcDir)
	return conn, err
}

/*
	1.accountKey 对应的是要获取的账户地址的私钥信息，在节点的keystore下面。
	2.passphrase是当前操作的账户在创建账户是输入的密码（字符串）。
*/
func GetEthAuth(accountKey string, passphrase string) (*bind.TransactOpts, error) {
	auth, err := bind.NewTransactor(strings.NewReader(accountKey), passphrase)
	return auth, err
}

/*
	以太坊节点的IP地址和端口（rpc端口），rwaurl: http://172.18.22.34:8546
*/
func ConnectToRpc(rwaurl string) (*ethclient.Client, error) {
	client, err := rpc.Dial(rwaurl)
	if err != nil {
		return nil, err
	}

	conn := ethclient.NewClient(client)
	return conn, nil
}
