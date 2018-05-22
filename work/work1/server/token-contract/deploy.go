package main

import (
	"znfz/server/token-contract/accmanager"
	//"znfz/server/token-contract/addrmanager"
	//"znfz/server/token-contract/subaccount"
	//"strings"
	"log"
	//"math/big"
	"time"
	"znfz/server/config"
)

const (
	key = `{"address":"38e7005d85117f7fc3b3508bb8adcf68ebb3cef9","crypto":{"cipher":"aes-128-ctr","ciphertext":"e99faebbc6a4fa1149e9f1bc23db42b9d640af9367357a14e7840d98feb52166","cipherparams":{"iv":"c41d79cda794cabc8b6ec07b7762c393"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5c9062cb8b675475474537a044552dec7dcf5c3bb929490bc644612f274c88d9"},"mac":"d84469064184122378bea0719fa3673da25d4ee2eb7050ad3fd1c96b7e89a127"},"id":"287e6657-6539-4ed5-8d99-a8b8f24708d0","version":3}`
	operationPhrase = "apple1"
)

var (
	accountAddress  = ""
	addSmartAddress = ""
)

func AddCall() {
	config.ParseToml("/root/work/launch_pf/src/znfz/server/conf/config.toml") // 初始化配置

	//addmanager_samrt_addr, addmanager_samert_name, _ := addrmanager.Deploy("address_manager", "address_test")
	//log.Println("addmanager_samrt_addr:", addmanager_samrt_addr, "samert_name:", addmanager_samert_name)
	//addmanager_samrt_addr := "0xe71872dce008e4d3c4d9bb13174e592f7ee02bda"

	//accmanager_smart_add, accmanager_smart_name, _ := accmanager.Deploy("account_manager", "test_am")
	//log.Println("accmanager_smart_add:", accmanager_smart_add, "accmanager_smart_name:", accmanager_smart_name)

	accmanager_smart_add := "0x0d03613d9ce942a9286bd9d52ba529ed3b11bdf9"

	//hashcode, _ := addrmanager.NewAddressAdd(key,operationPhrase,"", accmanager_smart_add, "account_manager", "test_am")
	//log.Println("address manager new address hashcode:", hashcode)

	tel := "12114561235"
	//newAddress, describe, _ := accmanager.NewAccount("test_acc123")

	newAddress := "0x31c2a129ec54710b28db650c45e3129350491deb"
	describe := `{"address":"31c2a129ec54710b28db650c45e3129350491deb","crypto":{"cipher":"aes-128-ctr","ciphertext":"df26dd552451109962d4066dc2445d8a8009ef25f61bab1cc79eaaee7d8fcbe0","cipherparams":{"iv":"494a1532cb0fa88baef5d88fec19b971"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f3747912acb035cb31de71d71416b8342e29261a21823409d1173875c3614b96"},"mac":"f57d6d05f29a5921925ba1c448f3bf87ad2c62e28d49d7ce69e7b898459c9f0c"},"id":"9ced9220-a2ff-4640-8b0e-9e7db7aa6460","version":3}`

	hashcode, _ := accmanager.NewAccountAdd(describe, "test_acc123", accmanager_smart_add, accmanager.Account{newAddress, "test_acc123", "test_acc123", describe, tel})
	log.Println("account manager new address hashcode:", hashcode)

	tel2 := "12114561236"

	newAddress2 := "0x642b69852c7ac97fbc9e1db06cabde7be313ed76"
	describe2 := `{"address":"642b69852c7ac97fbc9e1db06cabde7be313ed76","crypto":{"cipher":"aes-128-ctr","ciphertext":"355399c0be2cae5b5a2fab174f6021fbf427d763d6f91b1e5465920546edfa88","cipherparams":{"iv":"d9723017654f7f5b7056682d2559f5ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1a7035be42a0eadf5136c1b6795514545c3808329f56ec022434b50065226f98"},"mac":"2d3782038b0607845fc2d168c17f32d17643fc268d3b0bbe1b009fd78414948d"},"id":"6debd6e5-3064-4938-ba41-89d0b8f53c64","version":3}`

	//
	//newAddress2, describe2, _ := accmanager.NewAccount("test_acc124")
	hashcode, _ = accmanager.NewAccountAdd(describe2, "test_acc124", accmanager_smart_add, accmanager.Account{newAddress2, "test_acc124", "test_acc124", describe2, tel2})
	log.Println("account manager new address hashcode:", hashcode)

	//accountAddress = newAddress
	//addSmartAddress = addmanager_samrt_addr
	//
	//subaccount_smart_add, subaccount_smart_name, _ := subaccount.Deploy(key, operationPhrase, "sub_account", "sub_test")
	//log.Println("subaccount_smart_add:", subaccount_smart_add, "subaccount_smart_name:", subaccount_smart_name)
	//
	time.Sleep(60 * time.Second)
	//_, smart_name, smart_sy, _ := addrmanager.GetAddressArray(addmanager_samrt_addr, accmanager_smart_add)
	//addrmanager.GetAddressByKey("", smart_name)
	//addrmanager.GetAddressByKey(addmanager_samrt_addr, smart_sy)
	//addrmanager.GetAddressByKey(addmanager_samrt_addr, smart_name+smart_sy)

	account, _ := accmanager.GetAccountByAddr(accmanager_smart_add, newAddress)
	log.Println("accmanager GetAccountByAddr account:", account)

	account1, _, _ := accmanager.GetAccountByTel(accmanager_smart_add, tel)
	account2, _, _ := accmanager.GetAccountByTel(accmanager_smart_add, tel2)
	log.Println("accmanager 222222 GetAccountByTel account:", account1, "tel:", tel)
	log.Println("accmanager 222222 GetAccountByTel account2:", account2, "tel2:", tel2)

	// 测试白名单的添加删除
	accmanager.AddEmployStaff(describe, "test_acc123", accmanager_smart_add, newAddress2)

	time.Sleep(30 * time.Second)

	accmanager.GetEmployStaffs(accmanager_smart_add, newAddress)
}

func DelTestWhite(){
	config.ParseToml("/root/work/launch_pf/src/znfz/server/conf/config.toml") // 初始化配置

	accmanager_smart_add := "0x0d03613d9ce942a9286bd9d52ba529ed3b11bdf9"
	newAddress := "0x31c2a129ec54710b28db650c45e3129350491deb"
	describe := `{"address":"31c2a129ec54710b28db650c45e3129350491deb","crypto":{"cipher":"aes-128-ctr","ciphertext":"df26dd552451109962d4066dc2445d8a8009ef25f61bab1cc79eaaee7d8fcbe0","cipherparams":{"iv":"494a1532cb0fa88baef5d88fec19b971"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f3747912acb035cb31de71d71416b8342e29261a21823409d1173875c3614b96"},"mac":"f57d6d05f29a5921925ba1c448f3bf87ad2c62e28d49d7ce69e7b898459c9f0c"},"id":"9ced9220-a2ff-4640-8b0e-9e7db7aa6460","version":3}`
	newAddress2 := "0x642b69852c7ac97fbc9e1db06cabde7be313ed76"

	accmanager.GetEmployStaffs(accmanager_smart_add, newAddress)

	// 测试白名单的添加删除
	accmanager.DelEmployStaff(describe, "test_acc123", accmanager_smart_add, newAddress2)

	time.Sleep(30 * time.Second)

	accmanager.GetEmployStaffs(accmanager_smart_add, newAddress)

}

func main() {
	//test_num := big.NewInt(10)
	//log.Println("test_num:",test_num)
	//return

	balance, _ := accmanager.GetBalance("0x38e7005d85117f7fc3b3508bb8adcf68ebb3cef9", "http://172.18.22.34:8546")
	log.Println("0x38e7005d85117f7fc3b3508bb8adcf68ebb3cef9 balance:", balance)

	if balance == "0" {
		return
	}
	DelTestWhite()
}
