package accmanager

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

type Account struct {
	AccountAddr     string
	Name            string
	Password        string
	AccountDescribe string
	Telephone       string
}

type FindAccount struct {
	Name            string
	Password        string
	AccountDescribe string
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

	tx, err := token.AddAccount(auth, common.HexToAddress(account.AccountAddr), account.Name, account.Password, account.AccountDescribe, account.Telephone)
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

func GetAccountByTel(operationAddress string, telephone string) ([]FindAccount, []string, error) {
	glog.Infoln(lib.Log("operation", "", "GetAccountByTel"), operationAddress, telephone)
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return []FindAccount{}, []string{}, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return []FindAccount{}, []string{}, err
	}

	addressList, err := token.GetOneAddress(nil, telephone)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByTel find address,err:", err.Error())
		return []FindAccount{}, []string{}, err
	}
	if len(addressList) == 0 {
		glog.Errorln("Failed to request GetAccountByTel ,telephone is not found account info.")
		return []FindAccount{}, []string{}, errors.New("telephone is not found account info.")
	}

	var account_list = []FindAccount{}
	address_list := []string{}

	for adx := 0; adx < len(addressList); adx++ {
		address := addressList[adx]
		if address.String() == "0x0000000000000000000000000000000000000000" {
			glog.Errorln("Failed to request GetAccountByTel ,the addr of the account is nil.")
			continue
		}

		var account = FindAccount{}
		account, err = token.GetAccount(nil, address)
		if err != nil {
			glog.Errorln("Failed to request GetAccountByTel,err:", err.Error())
			continue
		}
		account_list = append(account_list, account)
		address_list = append(address_list, address.String())
	}

	glog.Infoln(lib.Log("account", "", "GetAccountByTel"), " operationAddress:", operationAddress, "address_list:", address_list, "account_list:", account_list)
	return account_list, address_list, nil
}

// 添加用人白名单
func AddEmployStaff(operationKey string, operationPhrase string, operationAddress string, staffAddr string) (string, error) {
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

	tx, err := token.AddEmployStaff(auth, common.HexToAddress(staffAddr))
	if err != nil {
		glog.Errorln("Failed to request acccount manager AddEmployStaff", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("account", "", "AddEmployStaff"), "hash code:", tx.Hash().String())
	return tx.Hash().String(), nil
}

// 删除用人白名单
func DelEmployStaff(operationKey string, operationPhrase string, operationAddress string, staffAddr string) (string, error) {
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

	tx, err := token.DelEmployStaff(auth, common.HexToAddress(staffAddr))
	if err != nil {
		glog.Errorln("Failed to request acccount manager AddEmployStaff", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("account", "", "DelEmployStaff"), "hash code:", tx.Hash().String())
	return tx.Hash().String(), nil
}

// 获取用人白名单
func GetEmployStaffs(operationAddress string, accountAddress string) ([]FindAccount, []string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return []FindAccount{}, []string{}, err
	}
	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return []FindAccount{}, []string{}, err
	}
	addressList, err := token.GetEmployStaffs(nil, common.HexToAddress(accountAddress))
	var account_list = []FindAccount{}
	address_list := []string{}

	for adx := 0; adx < len(addressList); adx++ {
		address := addressList[adx]
		if address.String() == "0x0000000000000000000000000000000000000000" {
			glog.Errorln("Failed to request GetEmployStaffs ,the addr of the account is nil.")
			continue
		}

		var account = FindAccount{}
		account, err = token.GetAccount(nil, address)
		if err != nil {
			glog.Errorln("Failed to request GetEmployStaffs,err:", err.Error())
			continue
		}
		account_list = append(account_list, account)
		address_list = append(address_list, address.String())
	}

	glog.Infoln(lib.Log("account", "", "GetEmployStaffs"), " operationAddress:", operationAddress, "address_list:", address_list, "account_list:", account_list)
	return account_list, address_list, nil
}

// 检查账户地址是否在白名单里
func CheckEmployStaffs(operationAddress string, accountAddress string) (bool, []string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return false, []string{}, err
	}
	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return false, []string{}, err
	}
	addressList, err := token.GetEmployStaffs(nil, common.HexToAddress(accountAddress))

	address_list := []string{}
	is_ok := false
	for adx := 0; adx < len(addressList); adx++ {
		address := addressList[adx]
		if address.String() == "0x0000000000000000000000000000000000000000" {
			glog.Errorln("Failed to request CheckEmployStaffs ,the addr of the account is nil.")
			continue
		}
		address_list = append(address_list, address.String())
		if address.String() == accountAddress {
			is_ok = true
		}
	}

	glog.Infoln(lib.Log("account", "", "GetEmployStaffs"), " operationAddress:", operationAddress, "address_list:", address_list, "is_ok:", is_ok)
	return is_ok, address_list, nil
}

func GetAllAccountsAddress(operationAddress string) ([]string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return []string{}, err
	}
	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return []string{}, err
	}
	addressList, err := token.GetAllAccountsAddr(nil)
	address_list := []string{}
	for adx := 0; adx < len(addressList); adx++ {
		address := addressList[adx]
		if address.String() == "0x0000000000000000000000000000000000000000" {
			glog.Errorln("Failed to request GetAllAccountsAddress ,the addr of the account is nil.")
			continue
		}
		address_list = append(address_list, address.String())
	}
	return address_list, nil
}
