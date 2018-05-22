package addrmanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"strings"
	"time"
	"znfz/conf_server/config"
	"znfz/server/token-contract/utils"
)

//const key = `{"address":"b87d3eb92a1a3b8f2a0e8ab10f0c516f0546815c","crypto":{"cipher":"aes-128-ctr","ciphertext":"0a5f81d64c07d4f99ccd48a545f09956f8acf678a8b5dae6966da624ebab65e3","cipherparams":{"iv":"e5af21fab199b01ac6602f5a4e309bb7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ffab1ceaed886fc4aedf71dc835bb4a35eaff7c6498f336af45d628dac034efb"},"mac":"005c4d129bf119e087741a8391e99c0638a9a29e322550657aebe15513f5935d"},"id":"1aabe58b-da16-4fba-ad5d-59fb6a11f354","version":3}`
//
//const (
//	passphrase = "apple1"
//)

var (
	deployAddress = ""
)

func Deploy(token_name string, token_symbol string) (string, string, error) {
	glog.Infoln("[address]",config.Opts().IpcDir)
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("[address] Failed to connect to the Ethereum client: ", err.Error())
		return "", "", err
	}
	glog.Errorln("************",config.Opts().ManagerPhrase,config.Opts().ManagerKey)
	auth, err := utils.GetEthAuth(config.Opts().ManagerKey, config.Opts().ManagerPhrase)
	if err != nil {
		glog.Errorln("[address] Failed to create authorized transactor: ", err.Error())
		return "", "", err
	}

	address, tx, token, err := DeployToken(auth, conn, token_name, token_symbol)
	if err != nil {
		glog.Errorln("[address] Failed to deploy new token contract: ", err.Error())
		return "", "", err
	}
	glog.Infoln("[address] token_name:", token_name, "token_symbol:", token_symbol, ",Contract pending deploy:", address.String(), ",Transaction waiting to be mined: ", tx.Hash().String())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(15 * time.Second) // Allow it to be processed by the local node :P

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		glog.Errorln("[address] address:", address.String(), "Failed to retrieve pending name: ", err.Error())
		return "", "", err
	}
	glog.Infoln("[address] address:", address.String(), " Pending name:", name)

	deployAddress = strings.ToLower(address.String())
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
	glog.Infoln(" GetTokenName Pending name:", name)
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
