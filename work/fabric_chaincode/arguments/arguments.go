package arguments



type ReqAccount struct {
	AccountAddr string
	Name string
	Telephone string // 电话号码
	BankCard string // 银行卡，保持只有一个，可更换。
	WeChat string // 微信
	Alipay string // 支付宝
}

type ReqIssueSubCxt struct {
	IssueCode    string     // 分配编号，唯一标识编号
	SubRoles     [][32]byte // 参与分配的职位
	Rtaios       []uint // 每个账户地址对应的比例
	SubWays      []uint // 分配方式，0表示按比例分配，1表示按定额分配。
	QuotaWays    []uint // 数据置空方式，0不置空，1按日置空，2按月置空。
	ResetTimes   []uint // 分配数据重置数据，按日的话，每天的这个时间，按月的话，每月1号的这个时间；1-6之间，表1点到6点之间更新
	AccountAddr	 string    //提案人
}

type ReqIssueScheduling struct {
	IssueCode string
	Roles 	[][32]byte //职位
	Joiners []string  //账户名
}

type ReqSettleAccounts struct {
	IssueCode string
	TotalConsume uint
	TransferId string
}