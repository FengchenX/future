package arguments

import "math/big"

type DeployArguments struct {
	OperationKeyStore string
	OperationPassWord string
	TokenName         string
	TokenSymbol       string
	SubPayer          string   // 分账时的支付账户地址
	ManagerPayee      string   // 管理费的收款地址
	ManagerRatio      *big.Int // 管理费的固定比例
	StoresNumber      string   // 门店编号
	Postscript        string   // 备注信息
}

type BindSmartArguments struct {
	OperationKeyStore string // 操作者password
	OperationPassWord string // 操作者keystore
	OperatingAddress  string // 操作者要操作的合约地址
	SmartAddress      string
	TokenName         string
	TokenSymbol       string
	StoresNumber      string // 门店编号
}

type ScheduleArguments struct {
	OperationKeyStore string
	OperationPassWord string
	SmartAddress      string
	FatherAddress     string   // 发布职位的父节点地址
	IssueRatio        *big.Int // 当前发布者所持有的百分比
	Roles             [][32]byte
	JobIds            []*big.Int
	Counts            []*big.Int
	Ratio             []*big.Int
	Whitelists        []string
}
