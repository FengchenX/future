package contracts

import (
	"math/big"
	"sub_account_service/blockchain_server/arguments"
	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/lib/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
)

var UseNonce = true

/*
	title:发布比例分配表
*/
func SetDistributionRatio(key_string string, pargs arguments.DistributionArguments) (string, error) {

	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(pargs.SmartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.IssueSubCxt(auth, pargs.IssueCode, pargs.SubRoles, pargs.Rtaios, pargs.SubWays, pargs.QuotaWays, pargs.ResetTimes)
	//tx, err := token.IssueSubCxt(auth, pargs.IssueCode, subAccounts, pargs.Rtaios)
	myeth.UpdateNonce(auth, conn, err, "SetDistributionRatio", pargs.IssueCode)
	if err != nil {
		glog.Errorln("Failed to request sub account PublishScheduleing: ", err.Error())
		return "", err
	}
	glog.Infoln("sub account PublishScheduleing success pending:", tx.Hash().String(), "smartAddress:", pargs.SmartAddress, "pargs:", pargs)
	return tx.Hash().String(), nil
}

/*
	title:判断当前要发布的排版编号是否已有过
	ps: true 表示已使用，false 表示未使用；
*/
func CheckSubCodeIsOk(operationAddress string, subCode string) (bool, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return false, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return false, err
	}

	keys_len, err := token.GetSubAccountKeysLen(nil)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return false, err
	}
	glog.Infoln("CheckSubCodeIsOk keys_len:", keys_len)

	for kdx := 0; kdx < int(keys_len.Int64()); kdx++ {
		s_code, err := token.SubAccountKeys(nil, big.NewInt(int64(kdx)))
		if err != nil {
			glog.Errorln("Failed to request SubAccountKeys", err.Error())
			continue
		}
		if s_code == subCode {
			return true, nil
		}
	}
	return false, nil
}

/*
	title:查询自己发布的所有分配比例编号
*/
func GetbindSubCode(operationAddress string, faddr string) ([]string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
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
	var accounts = []string{}
	account_addrs, err := token.GetbindIssueToAddr(nil, common.HexToAddress(faddr))
	if err != nil {
		glog.Errorln("Failed to request GetbindSubCode", err.Error())
		return []string{}, err
	}

	for idx := 0; idx < len(account_addrs); idx++ {
		accounts = append(accounts, lib.Byte32ToStr(account_addrs[idx]))
	}

	glog.Infoln("GetbindSubCode operationAddress:", operationAddress, "account:", accounts)

	return accounts, nil
}

/*
	title:根据编号查询分配比例
*/
func GetDistributionRatioByCode(operationAddress string, subCode string) ([][32]byte, []*big.Int, []*big.Int, []*big.Int, []*big.Int, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return [][32]byte{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return [][32]byte{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, err
	}

	if err != nil {
		glog.Errorln("Failed to request GetIssueSubCxt", err.Error(), "subCode:", subCode)
		return [][32]byte{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, err
	}

	//le1, addr, le2, err := token.GetIssueSubCxtLen(nil, subCode)
	//glog.Infoln(le1, le2, addr, err)

	roles, ratios, subWays, quotaWays, resetTimes, err := token.GetIssueSubCxt(nil, subCode)
	if err != nil {
		glog.Errorln("Failed to request GetIssueSubCxt", err.Error(), "subCode:", subCode)
		return [][32]byte{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, []*big.Int{}, err
	}

	glog.Infoln(" ---- GetDistributionRatioByCode operationAddress:", operationAddress, "subCode:", subCode, "roles:", roles, "ratios:", ratios, "subWays:", subWays, "quotaWays:", quotaWays, "resetTimes:", resetTimes)

	return roles, ratios, subWays, quotaWays, resetTimes, nil
}

/*
func GetIssueSubCxtLen(operationAddress string, issueCode string) (int64, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return 0, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return 0, err
	}

	sub_len, joinAddr, ratio, err := token.GetIssueSubCxtLen(nil, issueCode)
	if err != nil {
		glog.Errorln("Failed to request GetIssueSubCxtLen", err.Error(), "issueCode:", issueCode)
		return 0, err
	}

	glog.Infoln("GetIssueSubCxtLen operationAddress:", operationAddress, "sub_len:", sub_len, "issueCode:", issueCode, "joinAddr:", joinAddr.String(), "ratio:", ratio)

	return sub_len.Int64(), nil
}
*/

/*
	title:账户信息绑定
	ps：  每次调用覆盖修改。
*/
func BindAccountInfos(key_string, operationAddress string, account arguments.AccountArguments) (string, error) {
	glog.Infoln("config.Opts().IpcDir", config.Opts().IpcDir)

	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return "", err
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	glog.Infoln(auth)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
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

	tx, err := token.AddAccount(auth, common.HexToAddress(account.AccountAddr), account.Name, account.BankCard, account.WeChat, account.Alipay, account.Telephone)
	myeth.UpdateNonce(auth, conn, err, "BindAccountInfos", account)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Fail to add NewAddAccount", err.Error())
	//	return "", err
	//}
	glog.Infoln(lib.Log("account", "", "NewAccountAdd"), "success pending: 0x", tx.Hash().String(), "operationAddress:", operationAddress)
	return tx.Hash().String(), nil
}

/*
	title:获取当前合约协议里所有绑定过的用户钱包地址；
*/
func GetAccountByAddr(operationAddress string) ([]string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
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
	var accounts = []string{}
	account_addrs, err := token.GetAllAccountsAddr(nil)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return []string{}, err
	}

	for idx := 0; idx < len(account_addrs); idx++ {
		accounts = append(accounts, account_addrs[idx].String())
	}

	glog.Infoln(" operationAddress:", operationAddress, "account:", accounts)

	return accounts, nil
}

/*
	title:查询用户绑定的支付信息
*/
func GetAccountCxt(operationAddress string, faddr common.Address) (arguments.AccountArguments, error) {
	var account = arguments.AccountArguments{AccountAddr: faddr.String()}
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return account, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return account, err
	}

	account.Name, account.BankCard, account.WeChat, account.Alipay, account.Telephone, err = token.GetAccount(nil, faddr)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return account, err
	}

	//glog.Infoln(" ---- GetOneLedgerCxt operationAddress:", operationAddress, "ledger:", account)

	return account, nil
}

/*
	title:实时分账，分账后会生成分成账本；
	smartAddress,deploy地址
	issueCode:发布分配表时对应的发布编号；
	totalConsume ： 当前要分配的总金额；
	orderId ：这笔钱进三方的支付信息编号；
*/
func SettleAccounts(key_string string, smartAddress string, issueCode string, totalConsume *big.Int, orderId string) (string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}
	//glog.Infoln("***实时分账 :查询auth.Nonce", transferId, auth.Nonce.Int64())

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.SettleAccounts(auth, issueCode, totalConsume, orderId)
	myeth.UpdateNonce(auth, conn, err, "SettleAccounts", orderId)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Failed to request sub account SettleAccounts: ", err.Error(), "smartAddress:", smartAddress, "issueCode:", issueCode, "transferId:", transferId, "totalConsume:", totalConsume.Int64())
	//	return "", err
	//}
	return tx.Hash().String(), nil
}

/*
	title: 财务平台查询分账后生产的账本信息
	步骤：1. 先查出transferId（交易单号）对应下的所有参与分账的人；
		2. 遍历所有参与人，查出每个人的账本，针对每个人调用三方支付接口打钱；
		3. 打钱成功后，传入transferId和参与人地址，打钱成功的回包信息，更新这个人的账本状态；
*/

// 1. 先查出transferId（交易单号）对应下的所有参与分账的人；
func GetgetLedgerSubAddrs(operationAddress string, transferId string) ([]string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
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
	var accounts = []string{}
	account_addrs, err := token.GetLedgerSubAddrs(nil, transferId)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return []string{}, err
	}

	for idx := 0; idx < len(account_addrs); idx++ {
		accounts = append(accounts, account_addrs[idx].String())
	}

	glog.Infoln("GetgetLedgerSubAddrs operationAddress:", operationAddress, "account:", accounts)

	return accounts, nil
}

// 2. 查出每个人的账本
func GetOneLedgerCxt(operationAddress string, transferId string, uaddr common.Address) (arguments.EachLedgerCxt, error) {
	var ledger = arguments.EachLedgerCxt{}
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return ledger, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return ledger, err
	}

	ledger.Calculate, ledger.OrderId, ledger.Rflag, ledger.TransferDetails, err = token.GetOneLedgerCxt(nil, uaddr, transferId)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return ledger, err
	}

	glog.Infoln(" ---- GetOneLedgerCxt operationAddress:", operationAddress, "uaddr:", uaddr.String(), "transferId:", transferId, "ledger:", ledger)

	return ledger, nil
}

// 3.更新个人账本状态
func UpdateCalculateLedger(key_string, smartAddress, transferId, transferDetails string, uaddr common.Address) (string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}
	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.UpdateCalulateLedger(auth, uaddr, transferId, transferDetails)
	myeth.UpdateNonce(auth, conn, err, "UpdateCalculateLedger",transferId,"user:addr",uaddr)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Fail SettleAccounts: ", transferId, err.Error())
	//	return "", err
	//}
	return tx.Hash().String(), nil
}

// 4.查询/重置 每个分配比例按定额分配方式的已分配定额数
func GetSubCodeQuotaData(operationAddress string, subCode string, uaddr string) (int64, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return 0, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return 0, err
	}

	quotaNum, err := token.GetSubCodeQuotaData(nil, subCode, common.HexToAddress(uaddr))
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return 0, err
	}

	glog.Infoln(" ---- GetSubCodeQuotaData operationAddress:", operationAddress, "uaddr:", uaddr, "quotaNum:", quotaNum,"subCode:",subCode)

	return quotaNum.Int64(), nil
}

func ResetSubCodeQuotaData(key_string string, operationAddress string, subCode string, uaddr string) (string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return "", err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return "", err
	}

	tx, err := token.ResetSubCodeQuotaData(auth, subCode, common.HexToAddress(uaddr))
	myeth.UpdateNonce(auth, conn, err, "ResetSubCodeQuotaData" ,subCode,"userAddr:",uaddr)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Fail ResetSubCodeQuotaData: ", err.Error())
	//	return "", err
	//}
	return tx.Hash().String(), nil
}

func GetSubCodeBalance(operationAddress string, uaddr string) (int64, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return 0, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return 0, err
	}

	quotaNum, err := token.GetBalance(nil, common.HexToAddress(uaddr))
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return 0, err
	}

	glog.Infoln(" ---- GetSubCodeBalance operationAddress:", operationAddress, "quotaNum:", quotaNum)

	return quotaNum.Int64(), nil
}

/*
	title:获取合约已发布的所有编号
*/
func GetAllSubCodes(operationAddress string) ([]string, error) {
	code_list := []string{}
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return code_list, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return code_list, err
	}

	keys_len, err := token.GetSubAccountKeysLen(nil)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return code_list, err
	}
	glog.Infoln("GetAllSubCodes keys_len:", keys_len)

	for kdx := 0; kdx < int(keys_len.Int64()); kdx++ {
		s_code, err := token.SubAccountKeys(nil, big.NewInt(int64(kdx)))
		if err != nil {
			glog.Errorln("Failed to request SubAccountKeys", err.Error())
			continue
		}
		code_list = append(code_list, s_code)
	}
	glog.Infoln(" ---- GetSubCodeBalance GetAllSubCodes:", operationAddress, "code_list:", code_list)
	return code_list, nil
}

/*
	title: v3 发布排班表
*/
func IssueScheduling(key_string string, smartAddress string, subCode string, roles [][32]byte, joiners []string) (string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}
	glog.Infoln(" IssueScheduling****************", key_string, smartAddress, subCode, roles, joiners)
	//glog.Infoln("***实时分账 :查询auth.Nonce", transferId, auth.Nonce.Int64())

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	joiners_accs := []common.Address{}
	for jdx := 0; jdx < len(joiners); jdx++ {
		joiners_accs = append(joiners_accs, common.HexToAddress(joiners[jdx]))
	}

	tx, err := token.IssueScheduling(auth, subCode, roles, joiners_accs)
	myeth.UpdateNonce(auth, conn, err, "IssueScheduling",subCode)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Failed to request sub account SettleAccounts: ", err.Error())
	//	return "", err
	//}
	glog.Infoln(lib.Log("subacc", "", "IssueScheduling"), "success pending:", tx.Hash().String(), "smartAddress:", smartAddress, "subCode:", subCode)
	return tx.Hash().String(), nil
}

/*
	title: v3 查询排班表
*/
func GetSchedulingCxt(operationAddress string, subCode string) ([]string, []string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return []string{}, []string{}, err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return []string{}, []string{}, err
	}

	roles, joiners, err := token.GetSchedulingCxt(nil, subCode)
	if err != nil {
		glog.Errorln("Failed to request GetAccountByAddr", err.Error())
		return []string{}, []string{}, err
	}

	joiners_accs := []string{}
	role_arrys := []string{}
	for jdx := 0; jdx < len(joiners); jdx++ {
		joiners_accs = append(joiners_accs, joiners[jdx].String())
	}

	for jdx := 0; jdx < len(roles); jdx++ {
		role_arrys = append(role_arrys, lib.Byte32ToStr(roles[jdx]))
	}

	glog.Infoln(" ---- GetSchedulingCxt operationAddress:", operationAddress, "subCode:", subCode, "roles:", role_arrys, "joiners:", joiners_accs)
	return role_arrys, joiners_accs, nil
}

/*
	title: v3 设置已分配定额
*/
func SetSubCodeQuotaData(key_string string, operationAddress string, subCode string, uaddr string, number int64) (string, error) {
	conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client", err.Error())
		return "", err
	}

	if len(operationAddress) == 0 {
		operationAddress = deployAddress
	}
	auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}
	//glog.Infoln("***更新个人账本 :查询auth.Nonce", subCode, auth.Nonce.Int64())

	token, err := NewToken(common.HexToAddress(operationAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract", err.Error())
		return "", err
	}
	tx, err := token.SetSubCodeQuotaData(auth, subCode, common.HexToAddress(uaddr), big.NewInt(number))
	myeth.UpdateNonce(auth, conn, err,"SetSubCodeQuotaData", subCode)
	if err != nil {
		return "",err
	}
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Fail ResetSubCodeQuotaData: ", err.Error())
	//	return "", err
	//}

	glog.Infoln(" ---- SetSubCodeQuotaData operationAddress:", operationAddress, "hash:", tx.Hash().String(), "subCode:", subCode, "uaddr:", uaddr, "number:", number)

	return tx.Hash().String(), nil
}

/*
	title: v3 修改财务平台付款账户地址
*/
func ChangeSamrtPayer(key_string string, operationAddress, caddr string) (string, error) {
	return "", nil
	//conn, err := myeth.GetEthclient(config.Opts().IpcDir)
	//if err != nil {
	//	glog.Errorln("Failed to connect to the Ethereum client", err.Error())
	//	return "", err
	//}
	//
	//if len(operationAddress) == 0 {
	//	operationAddress = deployAddress
	//}
	//auth, err := myeth.ParseEthAuth(key_string, conn, UseNonce)
	//if err != nil {
	//	glog.Errorln("Failed to create authorized transactor: ", err.Error())
	//	return "", err
	//}
	////glog.Infoln("***更新个人账本 :查询auth.Nonce", subCode, auth.Nonce.Int64())
	//
	//token, err := NewToken(common.HexToAddress(operationAddress), conn)
	//if err != nil {
	//	glog.Errorln("Failed to instantiate a Token contract", err.Error())
	//	return "", err
	//}
	//tx, err := token.ChangePayer(auth, common.HexToAddress(caddr))
	//if err != nil {
	//	myeth.NonceMap.Reset(key_string, conn)
	//	glog.Errorln("Fail ResetSubCodeQuotaData: ", err.Error())
	//	return "", err
	//}
	//glog.Infoln(" ---- ChangeSamrtPayer operationAddress:", operationAddress, "hash:", tx.Hash().String(), "caddr:", caddr)
	//
	//return tx.Hash().String(), nil
}
