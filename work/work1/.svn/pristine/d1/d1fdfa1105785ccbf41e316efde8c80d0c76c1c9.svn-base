package addrmanager

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"znfz/conf_server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

func NewAddressAdd( operationKey string, operationPhrase string, operationAddress string,smartAddress string, name string, symbol string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client:", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor:", err.Error())
		return "", err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract:", err.Error())
		return "", err
	}

	// 事件监听
	//EventListen()

	tx, err := token.AddAddress(auth, common.HexToAddress(smartAddress), name, symbol)
	if err != nil {
		glog.Errorln("Failed to request NewAddressAdd:", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("address", "", "NewAddressAdd"), "success pending: 0x", tx.Hash().String())

	tx, err = token.AddAddressIndex(auth, common.HexToAddress(smartAddress), name, symbol)
	if err != nil {
		glog.Errorln("Failed to request NewAddressAdd:", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("address", "", "NewAddressAdd"), "success pending:", tx.Hash().String())

	return tx.Hash().String(), nil
}

func GetAddressByKey(operationAddress string, fkey string) ([]string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client:", err.Error())
		return []string{}, err
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract:", err.Error())
		return []string{}, err
	}
	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	address_list, err := token.GetAddressIndex(nil, fkey)
	if err != nil {
		glog.Errorln("Failed to GetAddressByKey:", err.Error())
		return []string{}, err
	}

	f_address := []string{}
	for fdx := 0; fdx < len(address_list); fdx++ {
		f_address = append(f_address, address_list[fdx].String())
	}

	glog.Infoln(lib.Log("address", "", "GetAddressByKey"), "fkey:", fkey, " ,f_address:", f_address)
	return f_address, nil
}

func GetAddressArray(operationAddress string, smartAddress string) (string, string, string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client:", err.Error())
		return "", "", "", err
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract:", err.Error())
		return "", "", "", err
	}
	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}

	smart_address, smart_name, smart_sy, err := token.GetAddressArray(nil, common.HexToAddress(smartAddress))
	if err != nil {
		glog.Errorln("Failed to GetAddressByKey:", err.Error())
		return "", "", "", err
	}

	glog.Infoln(lib.Log("address", "", "GetAddressArray"), "smart_address:", smart_address.String(), "smart_name:", smart_name, "smart_sy:", smart_sy)
	return smart_address.String(), smart_name, smart_sy, nil
}

// 事件监听
//event := token.FilterGloballog()
func EventListen(operationAddress string) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Fatalf("Failed to connect to the Ethereum client:", err.Error())
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Fatalf("Failed to instantiate a Token contract:", err.Error())
	}

	addressEvent, _ := token.FilterGloballog(nil)

	glog.Infoln(lib.Log("address", "", "addressEvent.Event.Logstr"), ":", addressEvent.Event.Logstr)
	// 关闭监听
	addressEvent.Close()
	//addressEvent := token.WatchGloballog()
	//addressEvent.Unsubscribe()
}
