package accmanager

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"math/big"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

type Account struct {
	AccountAddr     string
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}

type FindAccount struct {
	Name            string
	Password        string
	AccountDescribe string
	Role            *big.Int
	Telephone       string
}

func NewAccountAdd(operationKey string, operationPhrase string, operationAddress string, account Account) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor", err.Error())
		return "", err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	// 这里的地址是部署好了的智能合约的 address 地址（要确保合约已经部署成功）
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return "", err
	}

	tx, err := token.AddAccount(auth, common.HexToAddress(account.AccountAddr), account.Name, account.Password, account.AccountDescribe, account.Role, account.Telephone)
	if err != nil {
		glog.Errorln("Failed to request acccount manager NewAddAccount", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("account", "", "NewAccountAdd"), "success pending: 0x", tx.Hash().String(), "operationAddress:", operationAddress)
	return tx.Hash().String(), nil
}

func GetAccountByAddr(operationAddress string, address string) (FindAccount, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return FindAccount{}, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return FindAccount{}, err
	}
	var account = FindAccount{}
	account, err = token.GetAccount(nil, common.HexToAddress(address))
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return FindAccount{}, err
	}

	glog.Infoln(lib.Log("account", "", "GetAccountByAddr"), " operationAddress:", operationAddress, "account:", account)

	return account, nil
}

func GetAccountByTel(operationAddress string, telephone string) (FindAccount, string, error) {
	glog.Infoln(lib.Log("operation", "", "GetAccountByTel"), operationAddress, telephone)
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return FindAccount{}, "", err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return FindAccount{}, "", err
	}

	address, err := token.GetOneAddress(nil, telephone)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByTel find address,err:", err.Error())
		return FindAccount{}, "", err
	}
	if len(address.String()) == 0 {
		glog.Errorln("Failed to request GetAccountByTel ,telephone is not found account info.")
		return FindAccount{}, "", errors.New("telephone is not found account info.")
	}

	if address.String() == "0x0000000000000000000000000000000000000000" {
		glog.Errorln("Failed to request GetAccountByTel ,the addr of the account is nil.")
		return FindAccount{}, "", errors.New("the addr of the account is nil.")
	}

	var account = FindAccount{}
	account, err = token.GetAccount(nil, address)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByTel,err:", err.Error())
		return FindAccount{}, "", err
	}

	glog.Infoln(lib.Log("account", "", "GetAccountByTel"), " operationAddress:", operationAddress, "address:", address.String(), "account:", account)
	return account, address.String(), nil
}
