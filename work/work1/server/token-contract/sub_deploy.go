package main

import (
	"log"
	"znfz/server/token-contract/subaccount"
	"znfz/server/token-contract/addrmanager"
	"znfz/server/config"
	"znfz/server/arguments"
	"math/big"
	"time"
)

var (
	newAddress3 = "0x38e7005d85117f7fc3b3508bb8adcf68ebb3cef9"
	describe3 = `{"address":"38e7005d85117f7fc3b3508bb8adcf68ebb3cef9","crypto":{"cipher":"aes-128-ctr","ciphertext":"e99faebbc6a4fa1149e9f1bc23db42b9d640af9367357a14e7840d98feb52166","cipherparams":{"iv":"c41d79cda794cabc8b6ec07b7762c393"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5c9062cb8b675475474537a044552dec7dcf5c3bb929490bc644612f274c88d9"},"mac":"d84469064184122378bea0719fa3673da25d4ee2eb7050ad3fd1c96b7e89a127"},"id":"287e6657-6539-4ed5-8d99-a8b8f24708d0","version":3}`

	addmanager_samrt_addr = "0xC1f42A4722AEA9130328b07E45f61a823520F25C"

	newAddress = "0x31c2a129ec54710b28db650c45e3129350491deb"
	describe = `{"address":"31c2a129ec54710b28db650c45e3129350491deb","crypto":{"cipher":"aes-128-ctr","ciphertext":"df26dd552451109962d4066dc2445d8a8009ef25f61bab1cc79eaaee7d8fcbe0","cipherparams":{"iv":"494a1532cb0fa88baef5d88fec19b971"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f3747912acb035cb31de71d71416b8342e29261a21823409d1173875c3614b96"},"mac":"f57d6d05f29a5921925ba1c448f3bf87ad2c62e28d49d7ce69e7b898459c9f0c"},"id":"9ced9220-a2ff-4640-8b0e-9e7db7aa6460","version":3}`

	newAddress2 = "0x642b69852c7ac97fbc9e1db06cabde7be313ed76"
	describe2 = `{"address":"642b69852c7ac97fbc9e1db06cabde7be313ed76","crypto":{"cipher":"aes-128-ctr","ciphertext":"355399c0be2cae5b5a2fab174f6021fbf427d763d6f91b1e5465920546edfa88","cipherparams":{"iv":"d9723017654f7f5b7056682d2559f5ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1a7035be42a0eadf5136c1b6795514545c3808329f56ec022434b50065226f98"},"mac":"2d3782038b0607845fc2d168c17f32d17643fc268d3b0bbe1b009fd78414948d"},"id":"6debd6e5-3064-4938-ba41-89d0b8f53c64","version":3}`

)

func InitConf() {
	config.ParseToml("/root/work/launch_pf/src/znfz/server/conf/config.toml") // 初始化配置
}

func SubDeploy() {
	//addmanager_samrt_addr, addmanager_samert_name, _ := addrmanager.Deploy("address_manager", "address_test")
	//log.Println("addmanager_samrt_addr:", addmanager_samrt_addr, "samert_name:", addmanager_samert_name)

	sargs := arguments.DeployArguments{describe, "test_acc123", "sub_account", "sub_test_apple1", "0x38e7005d85117f7fc3b3508bb8adcf68ebb3cef9", newAddress3, big.NewInt(12), "appleTest001", "apple test,start 11:00"}
	subaccount_smart_add, subaccount_smart_name, _ := subaccount.Deploy(sargs)
	log.Println("subaccount_smart_add:", subaccount_smart_add, "subaccount_smart_name:", subaccount_smart_name)
	//subaccount_smart_add := "0xa32e6c724c21aee2101ee6e7024e3b526c9850ed"

	bargs := arguments.BindSmartArguments{describe, "test_acc123", addmanager_samrt_addr, subaccount_smart_add, "sub_account", "sub_test_apple1", "appleTest001"}
	hashcode, _ := addrmanager.NewAddressAdd(bargs)
	log.Println("address manager new address hashcode:", hashcode)

	roles := [][32]byte{[32]byte{'a'}, [32]byte{'b'}}
	jobIds := []*big.Int{big.NewInt(1), big.NewInt(2)}
	counts := []*big.Int{big.NewInt(1), big.NewInt(1)}
	ratio := []*big.Int{big.NewInt(15), big.NewInt(25)}
	white := []string{newAddress2, newAddress3}

	pargs := arguments.ScheduleArguments{describe, "test_acc123", subaccount_smart_add,
	newAddress, big.NewInt(88), roles, jobIds, counts, ratio, white}
	subaccount.PublishScheduleing(pargs)

	time.Sleep(60 * time.Second)

	subaccount.GetScheduleingCxt(subaccount_smart_add, jobIds)

	subaccount.ApplicationJob(describe2, "test_acc124", subaccount_smart_add, newAddress2, big.NewInt(1), newAddress)
	subaccount.ApplicationJob(describe3, "apple1", subaccount_smart_add, newAddress3, big.NewInt(2), newAddress)

	time.Sleep(30 * time.Second)

	subaccount.GetScheduleingCxt(subaccount_smart_add, []*big.Int{})

	subaccount.GetApplicationCxt(subaccount_smart_add)

	//symbl, _ := subaccount.GetTokenSymbol(subaccount_smart_add)
	//log.Println("************************:", symbl)

	// 订单存证
	subaccount.SetOrdersContent(describe, "test_acc123", subaccount_smart_add, "0x23888b05804ed02066d07da852ee974b04dc", "sub account frist test")

	// 分账
	subaccount.SettleAccounts(describe3, "apple1", subaccount_smart_add, big.NewInt(100))

	time.Sleep(30 * time.Second)
	subaccount.GetBalance(subaccount_smart_add, newAddress)
	subaccount.GetBalance(subaccount_smart_add, newAddress2)
	subaccount.GetBalance(subaccount_smart_add, newAddress3)

	subaccount.GetOrdersContent(subaccount_smart_add, "0x23888b05804ed02066d07da852ee974b04dc")

	subaccount.GetAllContentHashCxt(subaccount_smart_add)

	_, smart_name, smart_sy, stores_number, _ := addrmanager.GetAddressArray(addmanager_samrt_addr, subaccount_smart_add)
	addrmanager.GetAddressByKey(addmanager_samrt_addr, smart_name)
	addrmanager.GetAddressByKey(addmanager_samrt_addr, smart_sy)
	addrmanager.GetAddressByKey(addmanager_samrt_addr, smart_name+smart_sy)

	addrmanager.GetPaySamrtAddress(addmanager_samrt_addr, stores_number)

	subaccount.GetAllRatio(subaccount_smart_add)
}

func QueryTestSub() {
	subaccount_smart_add := "0x4B740ee6D118a1F55E9b0080E3858348FE1c85Fe"
	//subaccount.GetAllContentHashCxt(subaccount_smart_add)

	jobIds := []*big.Int{}
	subaccount.GetScheduleingCxt(subaccount_smart_add, jobIds)

	subaccount.GetAllRatio(subaccount_smart_add)

	subaccount.GetApplicationCxt(subaccount_smart_add)

	subaccount.GetCompanyCxt(subaccount_smart_add)


	// 分账
	//subaccount.SettleAccounts(describe3, "apple1", subaccount_smart_add, big.NewInt(100))
	//
	//time.Sleep(30 * time.Second)
	subaccount.GetBalance(subaccount_smart_add, "0x569ef95c3c40D7bfADF53d02EdDA3aFe0C9Bb17A")
	subaccount.GetBalance(subaccount_smart_add, "0x07f4370Fb847bc72Fcb9a131EE9daC2956E169F4")
	subaccount.GetBalance(subaccount_smart_add, "0x759D4c2E15587Fae036f183202F36CA3C667ccbD")
}

func queryBalance() {
	subaccount_smart_add := "0x10ba33684cbdb031392a1f2e4503cad5b1df71b2"
	subaccount.GetBalance(subaccount_smart_add, newAddress)
	subaccount.GetBalance(subaccount_smart_add, newAddress2)
	subaccount.GetBalance(subaccount_smart_add, newAddress3)

	stores_number := "麦当基南山分店"
	addrmanager.GetPaySamrtAddress(addmanager_samrt_addr, stores_number)
}

func AppliceTestAB() {
	subaccount_smart_add := "0x6b7b5139c8bf513d5236cab9b85a5f23e2a50fff"
	father := "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"

	var cookerAddr = "0x07f4370Fb847bc72Fcb9a131EE9daC2956E169F4"
	var cookerPass = "88882"
	var cookerDesp = "{\"address\":\"07f4370fb847bc72fcb9a131ee9dac2956e169f4\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"83f6fbff8fd14b878faf9f2d5b4d989d4ce3ab99d714554242437315da437ace\",\"cipherparams\":{\"iv\":\"fa65eb4763ce1e4df73caa37faa8f8b4\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"5365a632d01c0fb1635a3f71676be2014a1661f6c9eab5cb3c13c750fdbcfbc1\"},\"mac\":\"260d29ec47ab7d179b55117b3f7696018b969f31d7337ead6364c8babdb5407e\"},\"id\":\"205aa261-1721-4edc-ac8a-4253b2012f5e\",\"version\":3}"


	subaccount.ApplicationJob(cookerDesp, cookerPass, subaccount_smart_add, cookerAddr, big.NewInt(1), father)
}

func main() {
	InitConf()
	//SubDeploy()
	QueryTestSub()
	//queryBalance()
	//AppliceTestAB()
}
