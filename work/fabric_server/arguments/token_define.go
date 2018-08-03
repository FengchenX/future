package arguments

import "math/big"

type DeployArguments struct {
	TokenName   string
	TokenSymbol string
	SubPayer    string // 分账时的支付账户地址
	Postscript  string // 备注信息
}

type BindSmartArguments struct {
	OperatingAddress string // 操作者要操作的合约地址
	SmartAddress     string
	TokenName        string
	TokenSymbol      string
	StoresNumber     string // 门店编号
}

type DistributionArguments struct {
	SmartAddress string // 程序启动时发布的合约地址，基本固定
	IssueCode    string // 分配编号，唯一标识编号
	SubRoles   [][32]byte // 参与分配的职位
	Rtaios     []*big.Int // 每个账户地址对应的比例
	SubWays    []*big.Int // 分配方式，0表示按比例分配，1表示按定额分配。
	QuotaWays  []*big.Int // 数据置空方式，0不置空，1按日置空，2按月置空。
	ResetTimes []*big.Int // 分配数据重置数据，按日的话，每天的这个时间，按月的话，每月1号的这个时间；1-6之间，表1点到6点之间更新
}

type AccountArguments struct {
	AccountAddr string
	Name        string
	BankCard    string
	WeChat      string
	Alipay      string
	Telephone   string
}

type EachLedgerCxt struct {
	OrderId         string   // 分账的交易信息编号
	Calculate       *big.Int // 这笔交易分到的钱
	Rflag           bool     // 是否更新过打钱后的交易信息，false为未更新，true为已更新。
	TransferDetails string   // 打钱后的交易信息详情，Rflag为true，该字段有值。
}
