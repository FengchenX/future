package main

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"sub_account_service/blockchain_server/arguments"
	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/contracts"
	myeth "sub_account_service/blockchain_server/lib/eth"
	"time"
)

var (
	newAddress3 = "0x56a58d378fd5647de22bf10007ab2f49e47d83b7"
	describe3   = `{"address":"56a58d378fd5647de22bf10007ab2f49e47d83b7","crypto":{"cipher":"aes-128-ctr","ciphertext":"01b960ea11abb9baa2d9f5e4f8cec0eaaa2a6165f6a3c63b1fb3a767e86ad729","cipherparams":{"iv":"20f30ca3334be401de850a6fd912aa4f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2e219cc58e1468cc95863f62ce9d1bf5bbdec467f13ed4ecd2bc996f11c56822"},"mac":"6faf19ee181fbfd076a446854e2348020408b34e120824edc69085f2306bfa26"},"id":"44e95601-9851-422e-ae00-c0a481436fe6","version":3}`

	addmanager_samrt_addr = "0x6536dbd4df44fb7b85654866d91a5d9755b0046d"
	//addmanager_samrt_addr = "0xcd146c990330b536a274c4b97fcfd297e02d3a0b"

	newAddress = "0x38e7005d85117f7fc3b3508bb8adcf68ebb3cef9"
	describe   = `{"address":"38e7005d85117f7fc3b3508bb8adcf68ebb3cef9","crypto":{"cipher":"aes-128-ctr","ciphertext":"e99faebbc6a4fa1149e9f1bc23db42b9d640af9367357a14e7840d98feb52166","cipherparams":{"iv":"c41d79cda794cabc8b6ec07b7762c393"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5c9062cb8b675475474537a044552dec7dcf5c3bb929490bc644612f274c88d9"},"mac":"d84469064184122378bea0719fa3673da25d4ee2eb7050ad3fd1c96b7e89a127"},"id":"287e6657-6539-4ed5-8d99-a8b8f24708d0","version":3}`

	newAddress2 = "0x642b69852c7ac97fbc9e1db06cabde7be313ed76"
	describe2   = `{"address":"642b69852c7ac97fbc9e1db06cabde7be313ed76","crypto":{"cipher":"aes-128-ctr","ciphertext":"355399c0be2cae5b5a2fab174f6021fbf427d763d6f91b1e5465920546edfa88","cipherparams":{"iv":"d9723017654f7f5b7056682d2559f5ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1a7035be42a0eadf5136c1b6795514545c3808329f56ec022434b50065226f98"},"mac":"2d3782038b0607845fc2d168c17f32d17643fc268d3b0bbe1b009fd78414948d"},"id":"6debd6e5-3064-4938-ba41-89d0b8f53c64","version":3}`
)

func InitConf() {
	config.ParseToml("/root/work/launch_pf/src/sub_account_service/blockchain_server/conf/config.toml") // 初始化配置
	myeth.NewNonceMap()
	log.Println("init conf")
}

func SubDeploy() {

	sargs := arguments.DeployArguments{"apple1", "sub_account1", newAddress3, "apple test,start1d 11:00"}
	key1 := myeth.ParseKeyStore(describe3, "dev3_apple1")

	key_string, _ := json.Marshal(key1)

	subaccount_smart_add, subaccount_smart_name, _ := contracts.Deploy(string(key_string), sargs)

	log.Println("subaccount_smart_add:", subaccount_smart_add, "subaccount_smart_name:", subaccount_smart_name)
	addmanager_samrt_addr = subaccount_smart_add

	//bindAccountInfos()

	//code := "test:a8f9037a3cdb6d42fa4c0eae1e727e62"
	//QueryTestSub(code)
	//time.Sleep(5 * time.Second)
	//queryBalance()

}

func bindAccountInfos() {
	key1 := myeth.ParseKeyStore(describe3, "dev3_apple1")
	key_string, _ := json.Marshal(key1)

	for idx := 0; idx < 20; idx++ {
		account := arguments.AccountArguments{newAddress3, "testApple", "", "", "", ""}
		contracts.BindAccountInfos(string(key_string), addmanager_samrt_addr, account)
		time.Sleep(1 * time.Second)
	}
}

func QueryTestSub(code string) {
	//key1 := eth.ParseKeyStore(describe3, "dev3_apple1")
	key1 := myeth.ParseKeyStore(describe, "apple1")

	key_string, _ := json.Marshal(key1)

	isok, _ := contracts.CheckSubCodeIsOk(addmanager_samrt_addr, code)
	log.Println("isok ***************************** :", isok)
	//if isok == false {
	ratio := []*big.Int{big.NewInt(3333), big.NewInt(50000), big.NewInt(6667)}
	subWay := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(0)}
	quotaWay := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(0)}
	resetTime := []*big.Int{big.NewInt(0), big.NewInt(2), big.NewInt(0)}
	roles := [][32]byte{[32]byte{'a'}, [32]byte{'b'}, [32]byte{'c'}}

	pargs := arguments.DistributionArguments{addmanager_samrt_addr, code, roles, ratio, subWay, quotaWay, resetTime}
	contracts.SetDistributionRatio(string(key_string), pargs)

	time.Sleep(60 * time.Second)
	//}

	//sub_arr, _ := contracts.GetbindSubCode(addmanager_samrt_addr, newAddress)
	//for sdx := 0; sdx < len(sub_arr); sdx++ {
	//	contracts.GetIssueSubCxtLen(addmanager_samrt_addr, sub_arr[sdx])
	//
	//	contracts.GetDistributionRatioByCode(addmanager_samrt_addr, sub_arr[sdx])
	//}

	//contracts.GetDistributionRatioByCode(addmanager_samrt_addr, code)

	white := []string{newAddress, newAddress2, newAddress3}
	contracts.IssueScheduling(string(key_string), addmanager_samrt_addr, code, roles, white)

	//if len(sub_arr) > 0 {
	//	lshuih := "testaph0xxq120"
	//	queryBalance(code, lshuih)
	//}
	// 订单存证
	//subaccount.SetOrdersContent(describe, "apple1", subaccount_smart_add, "0x23888b05804ed02066d07da852ee974b04dc", "sub account frist test")
	//subaccount.SetOrdersContent(describe, "apple1", subaccount_smart_add, "0x23888b05804ed02066d07da852ee974b04dd", "sub account frist test2222222222222")
}

func queryFun(code string) {
	contracts.GetSchedulingCxt(addmanager_samrt_addr,code)
}

func queryBalance(code string, lshuih string) {
	key1 := myeth.ParseKeyStore(describe3, "dev3_apple1")

	key_string, _ := json.Marshal(key1)

	contracts.SettleAccounts(string(key_string), addmanager_samrt_addr, code, big.NewInt(20000), lshuih)

	time.Sleep(60 * time.Second)

	sub_arr, _ := contracts.GetgetLedgerSubAddrs(addmanager_samrt_addr, lshuih)

	for sdx := 0; sdx < len(sub_arr); sdx++ {
		contracts.GetOneLedgerCxt(addmanager_samrt_addr, lshuih, common.HexToAddress(sub_arr[sdx]))
		transferDetails := "update test by 1 apple"
		contracts.UpdateCalulateLedger(string(key_string), addmanager_samrt_addr, lshuih, transferDetails, common.HexToAddress(sub_arr[sdx]))
	}
	//

	time.Sleep(60 * time.Second)
	for sdx := 0; sdx < len(sub_arr); sdx++ {
		contracts.GetOneLedgerCxt(addmanager_samrt_addr, lshuih, common.HexToAddress(sub_arr[sdx]))
	}

	queryLedger(code, lshuih)

	contracts.GetSubCodeQuotaData(addmanager_samrt_addr, code, newAddress)
	//AppliceTestAB()
}

func queryLedger(code string, lshuih string) {
	contracts.GetDistributionRatioByCode(addmanager_samrt_addr, code)
	sub_arr, _ := contracts.GetgetLedgerSubAddrs(addmanager_samrt_addr, lshuih)

	for sdx := 0; sdx < len(sub_arr); sdx++ {
		contracts.GetOneLedgerCxt(addmanager_samrt_addr, lshuih, common.HexToAddress(sub_arr[sdx]))
		contracts.GetSubCodeQuotaData(addmanager_samrt_addr, code, sub_arr[sdx])
	}
}

func AppliceTestAB() {
	//code := "mjn:7c4eedbe4cce25f3f8f9052772d9e3aa"
	//contracts.GetSubCodeQuotaData(addmanager_samrt_addr, code, newAddress)
	contracts.GetSubCodeBalance(addmanager_samrt_addr, newAddress)
	contracts.GetSubCodeBalance(addmanager_samrt_addr, newAddress2)
	contracts.GetSubCodeBalance(addmanager_samrt_addr, newAddress3)
}

func GetAllLedges() {
	code := "AC:a8f9037a3cdb6d42fa4c0eae1e727e62"
	contracts.GetSubCodeQuotaData(addmanager_samrt_addr, code, "0xe30983c8a9a7e6f899011dae309cbcb1d20e181a")
	contracts.GetSubCodeQuotaData(addmanager_samrt_addr, code, "0x76daf8cb871a6c3f4616475e99e00725773c1b83")

	//codeList,_ := contracts.GetAllSubCodes(addmanager_samrt_addr)
	//
	//for cdx:=0;cdx<len(codeList);cdx++{
	//
	//}
}

func main() {
	InitConf()
	//GetAllLedges()
	//code := "AC:a8f9037a3cdb6d42fa4c0eae1e727e62" // txsd012 3333 3333 3334
	SubDeploy()
	//AppliceTestAB()
	//QueryTestSub(code)
	//queryFun(code)

	//bindAccountInfos()

	//lshuih := "testaph0xxq120"
	//queryBalance(code,lshuih)
	//queryLedger(code,lshuih)
	//lsuh2 := "2018063021001004000597467936"
	//queryLedger(code,lsuh2)
	//AppliceTestAB()
}