package subaccount

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"znfz/server/token-contract/utils"
	"github.com/ethereum/go-ethereum/common"
	"znfz/server/config"
	"github.com/golang/glog"
	"znfz/server/arguments"
	"time"
)

var (
	deployAddress = ""
)

func Deploy(targs arguments.DeployArguments) (string, string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", "", err
	}
	auth, err := utils.GetEthAuth(targs.OperationKeyStore, targs.OperationPassWord)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", "", err
	}

	address, tx, token, err := DeployToken(auth, conn, targs.TokenName, targs.TokenSymbol, common.HexToAddress(targs.SubPayer), common.HexToAddress(targs.ManagerPayee), targs.ManagerRatio, targs.StoresNumber, targs.Postscript, 18)
	if err != nil {
		glog.Errorln("Failed to deploy new token contract: ", err.Error(), ",token_name:", targs.TokenName, "token_symbol:", targs.TokenSymbol)
		return "", "", err
	}
	glog.Infoln("token_name:", targs.TokenName, "token_symbol:", targs.TokenSymbol, ",Contract pending deploy: ", address.String(), ",Transaction waiting to be mined:", tx.Hash().String())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(30 * time.Second) // Allow it to be processed by the local node :P

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		glog.Errorln("address:", address.String(), "Failed to retrieve pending name: ", err.Error())
		return "", "", err
	}

	deployAddress = strings.ToLower(address.String())
	glog.Infoln("address:", deployAddress, " Pending name:", name)

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
	//glog.Infoln(" GetTokenName Pending name:", name, "operationAddress:", operationAddress)
	return name, nil
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
