package main

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"znfz/server/protocol"
	"time"
)

//var addr string = "localhost:8888"

var managerAddr = "0x3a32e7d31b1418b0468505d2c9b1892053159bc4"
var managerPass = "88881"
var managerDesp = "{\"address\":\"3a32e7d31b1418b0468505d2c9b1892053159bc4\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"f73b035569cdc8d49c8606a132a5a76d245c89c44e83b159da7d306429165e87\",\"cipherparams\":{\"iv\":\"153e8872a5221791c9ec0aede3016269\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"a0a6ff7c97d4f34f5002f61f62c69f9a89f0810d958f107f0f9fbf94bb693895\"},\"mac\":\"273b9ed6f3a7ce2d4dfecf30181c1c03409ae8db341f0df52376838fcd7bd2ce\"},\"id\":\"27f79b80-e4ac-4ab1-87ee-45167229bf6e\",\"version\":3}"

var cookerAddr = "0x07f4370Fb847bc72Fcb9a131EE9daC2956E169F4"
var cookerPass = "88882"
var cookerDesp = "{\"address\":\"07f4370fb847bc72fcb9a131ee9dac2956e169f4\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"83f6fbff8fd14b878faf9f2d5b4d989d4ce3ab99d714554242437315da437ace\",\"cipherparams\":{\"iv\":\"fa65eb4763ce1e4df73caa37faa8f8b4\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"5365a632d01c0fb1635a3f71676be2014a1661f6c9eab5cb3c13c750fdbcfbc1\"},\"mac\":\"260d29ec47ab7d179b55117b3f7696018b969f31d7337ead6364c8babdb5407e\"},\"id\":\"205aa261-1721-4edc-ac8a-4253b2012f5e\",\"version\":3}"

var huangAddr = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"
var huangPass = "15219438281"
var huangDesp = "{\"address\":\"759d4c2e15587fae036f183202f36ca3c667ccbd\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"e9e86c32d3c130072d78cd66a58dd9e7dd84c7066652ecb8db983193b759bebb\",\"cipherparams\":{\"iv\":\"ba6e9792b52b22ae3055cb8d2fe8d90b\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"aac7b6c6ab10e0e57446a97f36485800dd5520c159be2be9e8610074c05a6061\"},\"mac\":\"6d103da9f03a32e5d12f79489f1de03fc3ac858c2b9b2bb3d8d3aa2be2fcc498\"},\"id\":\"962d9f5f-91e6-49a5-b6a0-70e3a9b2136e\",\"version\":3}"

var ta = "0xbAF2e8f1Aa66E97F5c11712945a335B2dC4C4336"
var tp = "9999"
var td = "{\"address\":\"87b67fcf385ce4919884fb982dd63be36ed3ef38\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"598ef34d426fed459fdfe47354ff105e2052496f0853306c45dac2cb37173be9\",\"cipherparams\":{\"iv\":\"bc6d86dec9eee329f07208e279c6dce8\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"02dbc0761bb3220b27a3a14de8d21f87e210312156e3efdcd43a33930800baf4\"},\"mac\":\"b121af4e14390fdb2b29ff823f747c7a7589456a4daaaacef87e5dad9ee2e11a\"},\"id\":\"9aec95aa-35e4-4ff6-b0b8-d39acfa89c18\",\"version\":3}"

var ja = "569ef95c3c40d7bfadf53d02edda3afe0c9bb17a"
var jp = "123456"
var jd = `{"address":"569ef95c3c40d7bfadf53d02edda3afe0c9bb17a","id":"6f996145-6b4b-455d-bf99-16089b7f6509","version":3,"crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"21d3b415546d67309b02799fb488f98c"},"ciphertext":"fc10d4bc320a91742e99c85fc041a211db9bf40dd407b6611d7f57ff156b545b","kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"4cb53365095e37064ed94c788fcb5ffceb78e84dad92e64d3e1b29cf1b0f2adb"},"mac":"5a7bfbe500a30cd5315cc34f14b1d6460018f1ff5eab92ce07b9598839b17fa3"}}`

var ts = "1524901459"
var JobAddress = "0x4c9ffe043bad9a8f466641851b05d82f804faf17"

var payAccount string = "0x759D4c2E15587Fae036f183202F36CA3C667ccbD"

// 试一下grpc效率
func run() {
	glog.Infoln("starting client")
	//addr:="39.108.80.66:8899"
	addr := "192.168.83.200:8899"
	//addr := "120.78.195.103:8899"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		glog.Errorln("Can't connect: "+addr, err)
		return
	}
	client := &Client{}
	client.C = protocol.NewApiServiceClient(conn)
	for j := 0; j < 1; j++ {
		//client.Check()
		//client.SetRegister("15219438281")
		//client.SetRegister("88882")
		//client.SetRegister("88883")
		//client.band("15219438281", "黄博士", string(huangDesp), huangAddr,"15219438281")
		//client.band("88882", "assur_cooker_1", cookerDesp, cookerAddr,"86222222")

		//client.GetRegister(managerAddr)

		//client.Login("86222222")
		//client.Login("15219438281")

		//client.Login("11111")
		//client.Login("22222")
		//client.Login("33333")
		//client.Login("55555")
		//client.GetEthBalance(managerAddr)

		//client.GetRegister(huangAddr)
		//client.GetRegister(cookerAddr)

		client.SetSchedule(huangAddr, payAccount,
			huangPass, huangDesp, "dr.huang", "麦当基", "厨王", time.Now().Unix()) // 设置排班
		//client.SetSchedule(ja, jp, jd, "junluo", "麦当基", "厨王", time.Now().Unix()) // 设置排班

		//client.GetLastDeploy("麦当基", "南山分店", JobAddress)

		//client.GetSchedule(huangAddr,"正稻", time.Now().String()) // 查询排班
		//client.GetCanapplyJob("麦当当", cookerAddr)
		//client.GetFindJob(cookerAddr, JobAddress)

		//client.FindJob(huangAddr, cookerAddr, JobAddress, cookerPass, cookerDesp, "厨王", 3) // 申请工作

		//client.GetApply(cookerAddr,JobAddress)

		//client.Order("3","牛肉饭：50块",JobAddress,8,500)
		//client.Order("4","牛肉饭：50块",JobAddress,9,500)
		//client.Order("2","鸡肉饭：40块",JobAddress)

		//client.GetOrder("1",JobAddress)
		//client.GetOrder("3",JobAddress)

		//client.GetAllOrder(huangAddr, "0x55c9F8542E5B1c946a1CC336607444AdBD8C72D8")

		//client.Pay(1, uint64(100000), managerAddr, managerPass, managerDesp, JobAddress)
		//client.Pay(1, uint64(100000), huangAddr, huangPass, huangDesp, JobAddress)
		//client.Pay(uint64(10000),managerAddr,JobAddress)

		//client.GetAllMoney( huangAddr,"麦当基")
		//client.GetAllMoney( cookerAddr,"麦当基")
		//client.GetAllOrder(cookerAddr, JobAddress)
		//client.GetAllOrder(cookerAddr, JobAddress)

		//client.HistoryJoin(cookerAddr,"launch_a")
		//client.GetAllIncome(huangAddr,"麦当基")
		//client.GetAllIncome(cookerAddr,"麦当基")
	}
	conn.Close()
}

func change(in string) {
	glog.Infoln(string(in))
}
