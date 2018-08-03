package model

// 客户端 查询排班
type ReqGetPaiBan struct {
	OwnerAddress string
	SubCode      string
}

// 服务端 查询排班
type RespGetPaiBan struct {
	StatusCode uint32
	Msg          string
	Roles        []string
	AddressArray []string
}

// 客户端 发布排班
type ReqSetPaiBan struct {
	OwnerAddress string
	UserKeyStore string
	UserParse    string
	KeyString    string
	UserAddress  string
	SubCode      string
	PaiBans      []*PaiBan
}

// 服务端 发布排班
type RespSetPaiBan struct {
	StatusCode uint32
	Hash	   string
	Msg        string
}

// 客户端 获取所有钱
type ReqGetMoney struct {
	UserAddress string
	StartTime   int64
	EndTime     int64
	Page        uint32
}

// 服务端 获取所有钱
type RespGetMoney struct {
	StatusCode uint32
	AllMoney   float64
	Month      float64
	Date       float64
	Bills      []*Bill
	PageCount  uint32
	Msg        string
}

// 客户端 获取以太坊的余额
type ReqGetEthBalance struct {
	UserAddress string
}

// 服务端 获取以太坊的余额
type RespGetEthBalance struct {
	StatusCode uint32
	Balance    string
	Msg        string
}

// 客户端 发布比例分配表
type ReqSchedule struct {
	UserAddress  string
	UserKeyStore string
	UserParse    string
	KeyString    string
	ScheduleName string
	Rss          []*Rs
	Message      string
}

// 服务端 发布比例分配表
type RespSchedule struct {
	StatusCode   uint32
	Hash         string
	ScheduleName string
	Msg          string
}

// 客户端 查询排班
type ReqGetSchedue struct {
	UserAddress  string
	ScheduleName string
}

// 服务端 查询排班
type RespGetSchedue struct {
	StatusCode uint32
	Accounts   []*UserAccount
	Schedules  []*Rs
	Msg        string
}

// 客户端 绑定支付账户
type ReqSetAccount struct {
	UserKeyStore string
	UserParse    string
	KeyString    string
	UserAccount  *UserAccount
}

// 服务端 绑定支付账户
type RespSetAccount struct {
	StatusCode uint32
	Hash	   string
	Msg        string
}

// 客户端 查询支付账号
type ReqGetAccount struct {
	UserAddress string
}

// 服务端 查询支付账号
type RespGetAccount struct {
	StatusCode  uint32
	UserAccount *UserAccount
	Msg         string
}

type ReqGetAllSchedule struct {
	UserAddress string
	Pages       uint64
}

// 服务端 查询排班
type RespGetAllSchedule struct {
	StatusCode uint32
	Schedules  []*Sd
	Pages      uint64
	PagesCount uint64
	Msg        string
}

type ReqGetQuo struct {
	UserAddr   string
	SubCode string
}

type RespGetQuo struct {
	StatusCode uint32
	Money      float64
	Msg        string
}

type ReqResetQuo struct {
	UserAddr   string
	KeyStore   string
	ScheduleId string
}

type RespResetQuo struct {
	Flag bool
	Hash string
	Msg  string
}

type ReqGetTrans struct {
	Hash string
}

type RespGetTrans struct {
	Flag bool
	Msg  string
}

type ReqNewScheduleId struct {
	Index string
}

// 服务端 根据sheduleid查询账本
type RespNewScheduleId struct {
	StatusCode   uint32
	ScheduleName string
	Msg          string
}

// -------------------------------------------------------
// 客户端 根据sheduleid查询账本
type ReqGetABBySh struct {
	OrderId      []string
	ScheduleName string
}

// 服务端 根据sheduleid查询账本
type RespGetABBySh struct {
	StatusCode uint32
	Ods        []*OrderDliver
	Msg        string
}

// 客户端 根据orderid查询账本
type ReqGetABById struct {
	OrderId      string
	ScheduleName string
}

// 服务端 根据orderid查询账本
type RespGetABById struct {
	StatusCode uint32
	Abs        []*AccountBook
	Msg        string
}

// 客户端 三方支付
type ReqThreeSetBill struct {
	UserAddress  string   //用户地址
	UserKeyStore string
	UserParse    string
	KeyString    string   //秘钥
	ScheduleName string   //排班编号
	Money        float64
	OrderId      string
}

// 服务端 三方支付
type RespThreeSetBill struct {
	StatusCode uint32
	Hash       string
	Msg        string
}

// 客户端 确认交易（三方平台打钱成功后回调）
type ReqConfirm struct {
	UserAddress     string
	UserKeyStore    string
	UserParse       string
	KeyString       string
	OrderId         string
	TransferDetails string
}

// 服务端 确认交易（三方平台打钱成功后回调）
type RespConfirm struct {
	StatusCode uint32
	Hash       string
	Msg        string
}

// 客户端 查询用户账本
type ReqGetAccountBook struct {
	UserAddress string
	OrderId     string
}

// 服务端 查询用户账本
type RespGetAccountBook struct {
	StatusCode      uint32
	OrderId         string
	Money           float64
	Rflag           bool
	TransferDetails string
	Msg             string
}

// 客户端 获取用户收入明细
type ReqGetAllIncome struct {
	CompanyName string
	UserAddress string
}

// 服务端 获取用户收入明细
type RespGetAllIncome struct {
	StatusCode uint32
	Msg        string
}

// 客户端 根据哈希值获取上链状态
type ReqGetByHash struct {
	UserAddress string
	Hash        string
}

// 服务端 根据哈希值获取上链状态
type RespGetByHash struct {
	StatusCode uint32
	Msg        string
}

type CReloadConfig struct {
	OperateTimeout uint32
	LocalAddress   string
	Port           string
	AccAddress     string
	EthAddress     string
	IpcDir         string
	ServerId       string
	ManagerKey     string
	ManagerPhrase  string
	KeyDir         string
}

type SReloadConfig struct {
}

// 设置已每个账户已分配的定额数
type ReqSetQuota struct {
	OwnerAddress string
	UserKeyStore string
	UserParse    string
	KeyString    string
	UserAddress  string
	SetNumber    int64
	SubCode      string
}

// 服务端 发布排班
type RespCommonSet struct {
	StatusCode uint32
	Msg        string
}

// 修改财务平台的付款账户地址
type ReqChangePayer struct {
	OwnerAddress string
	UserKeyStore string
	UserParse    string
	KeyString    string
	PayerAddress string
}
