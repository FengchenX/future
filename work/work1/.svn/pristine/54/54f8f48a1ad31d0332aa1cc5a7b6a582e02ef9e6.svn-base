package accmanager

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"strings"
	"time"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

var (
	deployAddress = ""
)

func Deploy(token_name string, token_symbol string) (string, string, error) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("[account] Failed to connect to the Ethereum client: ", err.Error())
		return "", "", err
	}
	auth, err := utils.GetEthAuth(config.Opts().ManagerKey, config.Opts().ManagerPhrase)
	if err != nil {
		glog.Errorln("[account] Failed to create authorized transactor: ", err.Error())
		return "", "", err
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := DeployToken(auth, conn, token_name, token_symbol)
	if err != nil {
		glog.Errorln("[account] Failed to deploy new token contract: %v", err, "token_name:", token_name, "token_symbol:", token_symbol)
		return "", "", err
	}
	glog.Infoln("[account] token_name:", token_name, "token_symbol:", token_symbol, ",Contract pending deploy:", address.String(), ",Transaction waiting to be mined: ", tx.Hash().String())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(15 * time.Second) // Allow it to be processed by the local node :P

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		glog.Errorln("[account] address:", address.String(), "Failed to retrieve pending name: ", err.Error())
		return "", "", err
	}
	deployAddress = strings.ToLower(address.String())
	glog.Infoln(lib.Log("account", "", "Deploy"), address.String(), " Pending name:", name)
	return deployAddress, name, nil
}

func GetTokenName(operationAddress string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		glog.Errorln("Failed to retrieve GetTokenName pending name: ", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("account", "", "GetTokenName"), " Pending name:", name)
	return name, nil
}

func NewAccount(c_password string) (string, string, error) {
	glog.Infoln("newaccount", c_password, config.Opts().KeyDir)
	ks := keystore.NewKeyStore(config.Opts().KeyDir, keystore.StandardScryptN, keystore.StandardScryptP)

	address, err := ks.NewAccount(c_password)
	if err != nil {
		glog.Errorln("err:",c_password, err)
		return "", "", err
	}
	account, err := ks.Export(address, c_password, c_password)

	if err != nil {
		glog.Errorln("account NewAccount fail,err:", err)
		return "", "", err
	}
	glog.Infoln(lib.Log("account", "", "NewAccount"), "address:", address.Address.Hex(), "account:", string(account))
	return address.Address.Hex(), string(account), nil
}

func GetBalance(address string, rwaurl string) (string, error) {
	client, err := utils.ConnectToRpc(rwaurl)
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

func GetTokenSymbol(operationAddress string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}
	symbol, err := token.Symbol(&bind.CallOpts{Pending: true})
	if err != nil {
		glog.Errorln("Failed to retrieve GetTokenName pending name: ", err.Error())
		return "", err
	}
	glog.Infoln(" GetTokenSymbol Pending symbol:", symbol, "operationAddress:", operationAddress)
	return symbol, nil
}