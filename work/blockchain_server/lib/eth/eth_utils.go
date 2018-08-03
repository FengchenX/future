package myeth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/glog"
	"sub_account_service/blockchain_server/lib"
	"time"
)

var (
	// connecting flag : if the rpc connection with eth node is down
	// the flag must be false else is true
	CF        bool
	ethClient *ethclient.Client
)

// ipcDir build a http-rpc connection to node's dir geth.ipc，
// the dir is just like: /root/work/launch_pf/src/ethbase/geth.ipc
func GetEthclient(ipcDir string) (*ethclient.Client, error) {

	if ethClient == nil {
		conn, err := Dial(ipcDir)
		return conn, err
	}

	return ethClient, nil
}

func Dial(ipcDir string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(ipcDir)
	ethClient = conn

	//NonceMap.ResetAll(conn)
	return conn, err
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

func GT(dir, hash string) bool {
	client, err := GetEthclient(dir)
	if err != nil {
		glog.Errorln(lib.Loger("GetTrans", hash), err)
		return false
	}
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	_, ispending, _ := client.TransactionByHash(ctx, common.HexToHash(hash))
	cancel()
	return ispending
}

func GetBalance(address string, rwaurl string) (string, error) {
	client, err := ConnectToRpc(rwaurl)
	if err != nil {
		glog.Errorln(lib.Log("ERR ACCOUNT", "", "GetBalance"), "account GetBalance balance ConnectToRpc fail", err.Error())
		return "0", err
	}

	balance, err := client.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	if err != nil {
		glog.Errorln(lib.Log("ERR account", "", "GetBalance"), address, err.Error())
		return "0", err
	}
	glog.Infoln(lib.Log("account", "", "GetBalance"), " balance:", balance.String())
	return balance.String(), nil
}

func ParseUserName(addr string) string {
	if len(addr) < 2 {
		return ""
	}

	header := addr[:2]

	if string(header) == "0x" || string(header) == "0X" {
		return addr
	}

	addr = "0x" + addr
	return addr
}
