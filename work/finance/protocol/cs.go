package protocol

// 客户端 查询排班
type ReqGetPaiBan struct {
	OwnerAddress string `json:"OwnerAddress"`
	SubCode      string `json:"SubCode"`
}

// 服务端 查询排班
type RespGetPaiBan struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	//PaiBans    []*PaiBan `protobuf:"bytes,2,rep,name=PaiBans" json:"PaiBans,omitempty"`
	Msg          string
	Roles        []string
	AddressArray []string
}

// 客户端 发布排班
type ReqSetPaiBan struct {
	OwnerAddress string    `protobuf:"bytes,1,opt,name=OwnerAddress" json:"OwnerAddress,omitempty"`
	UserKeyStore string    `protobuf:"bytes,2,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse    string    `protobuf:"bytes,3,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString    string    `protobuf:"bytes,4,opt,name=KeyString" json:"KeyString,omitempty"`
	UserAddress  string    `protobuf:"bytes,5,opt,name=UserAddress" json:"UserAddress,omitempty"`
	SubCode      string    `json:"SubCode,omitempty"`
	PaiBans      []*PaiBan `protobuf:"bytes,6,rep,name=PaiBans" json:"PaiBans,omitempty"`
}

// 服务端 发布排班
type RespSetPaiBan struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Msg        string
}

// 客户端 获取所有钱
type ReqGetMoney struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	StartTime   int64  `protobuf:"varint,2,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime     int64  `protobuf:"varint,3,opt,name=EndTime" json:"EndTime,omitempty"`
	Page        uint32 `protobuf:"varint,4,opt,name=Page" json:"Page,omitempty"`
}

// 服务端 获取所有钱
type RespGetMoney struct {
	StatusCode uint32  `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	AllMoney   float64 `protobuf:"fixed64,2,opt,name=AllMoney" json:"AllMoney,omitempty"`
	Month      float64 `protobuf:"fixed64,3,opt,name=Month" json:"Month,omitempty"`
	Date       float64 `protobuf:"fixed64,4,opt,name=Date" json:"Date,omitempty"`
	Bills      []*Bill `protobuf:"bytes,5,rep,name=Bills" json:"Bills,omitempty"`
	PageCount  uint32  `protobuf:"varint,6,opt,name=PageCount" json:"PageCount,omitempty"`
	Msg        string
}

// 客户端 获取以太坊的余额
type ReqGetEthBalance struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
}

// 服务端 获取以太坊的余额
type RespGetEthBalance struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Balance    string `protobuf:"bytes,2,opt,name=Balance" json:"Balance,omitempty"`
	Msg        string
}

// 客户端 发布比例分配表
type ReqSchedule struct {
	UserAddress  string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	UserKeyStore string `protobuf:"bytes,2,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse    string `protobuf:"bytes,3,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString    string `protobuf:"bytes,4,opt,name=KeyString" json:"KeyString,omitempty"`
	ScheduleName string `protobuf:"bytes,5,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
	Rss          []*Rs  `protobuf:"bytes,6,rep,name=Rss" json:"Rss,omitempty"`
	Message      string `protobuf:"bytes,7,opt,name=Message" json:"Message,omitempty"`
}

// 服务端 发布比例分配表
type RespSchedule struct {
	StatusCode   uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Hash         string `protobuf:"bytes,2,opt,name=Hash" json:"Hash,omitempty"`
	ScheduleName string `protobuf:"bytes,3,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
	Msg          string
}

// 客户端 查询排班
type ReqGetSchedue struct {
	UserAddress  string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	ScheduleName string `protobuf:"bytes,2,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
}

// 服务端 查询排班
type RespGetSchedue struct {
	StatusCode uint32         `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Accounts   []*UserAccount `protobuf:"bytes,2,rep,name=Accounts" json:"Accounts,omitempty"`
	Schedules  []*Rs          `protobuf:"bytes,3,rep,name=Schedules" json:"Schedules,omitempty"`
	Msg        string
}

// 客户端 绑定支付账户
type ReqSetAccount struct {
	UserKeyStore string `protobuf:"bytes,1,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse    string `protobuf:"bytes,2,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString    string `protobuf:"bytes,3,opt,name=KeyString" json:"KeyString,omitempty"`
	UserAccount  *UserAccount
}

// 服务端 绑定支付账户
type RespSetAccount struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Msg        string
}

// 客户端 查询支付账号
type ReqGetAccount struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
}

// 服务端 查询支付账号
type RespGetAccount struct {
	StatusCode  uint32       `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	UserAccount *UserAccount `protobuf:"bytes,2,opt,name=UserAccount" json:"UserAccount,omitempty"`
	Msg         string
}

type ReqGetAllSchedule struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	Pages       uint64 `protobuf:"varint,2,opt,name=Pages" json:"Pages,omitempty"`
}

// 服务端 查询排班
type RespGetAllSchedule struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Schedules  []*Sd  `protobuf:"bytes,2,rep,name=Schedules" json:"Schedules,omitempty"`
	Pages      uint64 `protobuf:"varint,3,opt,name=Pages" json:"Pages,omitempty"`
	PagesCount uint64 `protobuf:"varint,4,opt,name=PagesCount" json:"PagesCount,omitempty"`
	Msg        string
}

type ReqGetQuo struct {
	UserAddr   string `protobuf:"bytes,1,opt,name=UserAddr" json:"UserAddr,omitempty"`
	ScheduleId string `protobuf:"bytes,2,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
}

type RespGetQuo struct {
	StatusCode uint32  `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Money      float64 `protobuf:"fixed64,2,opt,name=Money" json:"Money,omitempty"`
	Msg        string
}

type ReqResetQuo struct {
	UserAddr   string `protobuf:"bytes,1,opt,name=UserAddr" json:"UserAddr,omitempty"`
	KeyStore   string `protobuf:"bytes,2,opt,name=KeyStore" json:"KeyStore,omitempty"`
	ScheduleId string `protobuf:"bytes,3,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
}

type RespResetQuo struct {
	Flag bool `protobuf:"varint,1,opt,name=Flag" json:"Flag,omitempty"`
	Msg  string
}

type ReqGetTrans struct {
	Hash string `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
}

type RespGetTrans struct {
	Flag bool `protobuf:"varint,1,opt,name=Flag" json:"Flag,omitempty"`
	Msg  string
}

type ReqNewScheduleId struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
}

// 服务端 根据sheduleid查询账本
type RespNewScheduleId struct {
	StatusCode   uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	ScheduleName string `protobuf:"bytes,2,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
	Msg          string
}

// -------------------------------------------------------
// 客户端 根据sheduleid查询账本
type ReqGetABBySh struct {
	OrderId      []string `protobuf:"bytes,1,rep,name=OrderId" json:"OrderId,omitempty"`
	ScheduleName string   `protobuf:"bytes,2,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
}

// 服务端 根据sheduleid查询账本
type RespGetABBySh struct {
	StatusCode uint32         `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Ods        []*OrderDliver `protobuf:"bytes,2,rep,name=Ods" json:"Ods,omitempty"`
	Msg        string
}

// 客户端 根据orderid查询账本
type ReqGetABById struct {
	OrderId      string `protobuf:"bytes,1,opt,name=OrderId" json:"OrderId,omitempty"`
	ScheduleName string `protobuf:"bytes,2,opt,name=ScheduleName" json:"ScheduleName,omitempty"`
}

// 服务端 根据orderid查询账本
type RespGetABById struct {
	StatusCode uint32         `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Abs        []*AccountBook `protobuf:"bytes,2,rep,name=Abs" json:"Abs,omitempty"`
	Msg        string
}

// 客户端 三方支付
type ReqThreeSetBill struct {
	UserAddress  string  `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"` //用户地址
	UserKeyStore string  `protobuf:"bytes,2,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse    string  `protobuf:"bytes,3,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString    string  `protobuf:"bytes,2,opt,name=KeyString" json:"KeyString,omitempty"`       //秘钥
	ScheduleName string  `protobuf:"bytes,3,opt,name=ScheduleName" json:"ScheduleName,omitempty"` //排班编号
	Money        float64 `protobuf:"fixed64,5,opt,name=Money" json:"Money,omitempty"`
	OrderId      string  `protobuf:"bytes,6,opt,name=OrderId" json:"OrderId,omitempty"`
}

// 服务端 三方支付
type RespThreeSetBill struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Hash       string `protobuf:"bytes,2,opt,name=Hash" json:"Hash,omitempty"`
	Msg        string
}

// 客户端 确认交易（三方平台打钱成功后回调）
type ReqConfirm struct {
	UserAddress     string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	UserKeyStore    string `protobuf:"bytes,2,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse       string `protobuf:"bytes,3,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString       string `protobuf:"bytes,2,opt,name=KeyString" json:"KeyString,omitempty"`
	OrderId         string `protobuf:"bytes,3,opt,name=OrderId" json:"OrderId,omitempty"`
	TransferDetails string `protobuf:"bytes,4,opt,name=TransferDetails" json:"TransferDetails,omitempty"`
}

// 服务端 确认交易（三方平台打钱成功后回调）
type RespConfirm struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Hash       string `protobuf:"bytes,2,opt,name=Hash" json:"Hash,omitempty"`
	Msg        string
}

// 客户端 查询用户账本
type ReqGetAccountBook struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	OrderId     string `protobuf:"bytes,2,opt,name=OrderId" json:"OrderId,omitempty"`
}

// 服务端 查询用户账本
type RespGetAccountBook struct {
	StatusCode      uint32  `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	OrderId         string  `protobuf:"bytes,2,opt,name=OrderId" json:"OrderId,omitempty"`
	Money           float64 `protobuf:"fixed64,3,opt,name=Money" json:"Money,omitempty"`
	Rflag           bool    `protobuf:"varint,4,opt,name=Rflag" json:"Rflag,omitempty"`
	TransferDetails string  `protobuf:"bytes,5,opt,name=TransferDetails" json:"TransferDetails,omitempty"`
	Msg             string
}

// 客户端 获取用户收入明细
type ReqGetAllIncome struct {
	CompanyName string `protobuf:"bytes,1,opt,name=CompanyName" json:"CompanyName,omitempty"`
	UserAddress string `protobuf:"bytes,2,opt,name=UserAddress" json:"UserAddress,omitempty"`
}

// 服务端 获取用户收入明细
type RespGetAllIncome struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Msg        string
}

// 客户端 根据哈希值获取上链状态
type ReqGetByHash struct {
	UserAddress string `protobuf:"bytes,1,opt,name=UserAddress" json:"UserAddress,omitempty"`
	Hash        string `protobuf:"bytes,2,opt,name=Hash" json:"Hash,omitempty"`
}

// 服务端 根据哈希值获取上链状态
type RespGetByHash struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Msg        string
}

type CReloadConfig struct {
	OperateTimeout uint32 `protobuf:"varint,1,opt,name=Operate_timeout,json=OperateTimeout" json:"Operate_timeout,omitempty"`
	LocalAddress   string `protobuf:"bytes,2,opt,name=LocalAddress" json:"LocalAddress,omitempty"`
	Port           string `protobuf:"bytes,3,opt,name=Port" json:"Port,omitempty"`
	AccAddress     string `protobuf:"bytes,4,opt,name=AccAddress" json:"AccAddress,omitempty"`
	EthAddress     string `protobuf:"bytes,5,opt,name=EthAddress" json:"EthAddress,omitempty"`
	IpcDir         string `protobuf:"bytes,6,opt,name=IpcDir" json:"IpcDir,omitempty"`
	ServerId       string `protobuf:"bytes,7,opt,name=ServerId" json:"ServerId,omitempty"`
	ManagerKey     string `protobuf:"bytes,8,opt,name=ManagerKey" json:"ManagerKey,omitempty"`
	ManagerPhrase  string `protobuf:"bytes,9,opt,name=ManagerPhrase" json:"ManagerPhrase,omitempty"`
	KeyDir         string `protobuf:"bytes,10,opt,name=KeyDir" json:"KeyDir,omitempty"`
}

type SReloadConfig struct {
}

// 设置已每个账户已分配的定额数
type ReqSetQuota struct {
	OwnerAddress string `protobuf:"bytes,1,opt,name=OwnerAddress" json:"OwnerAddress,omitempty"`
	UserKeyStore string `protobuf:"bytes,2,opt,name=UserKeyStore" json:"UserKeyStore,omitempty"`
	UserParse    string `protobuf:"bytes,3,opt,name=UserParse" json:"UserParse,omitempty"`
	KeyString    string `protobuf:"bytes,4,opt,name=KeyString" json:"KeyString,omitempty"`
	UserAddress  string `protobuf:"bytes,5,opt,name=UserAddress" json:"UserAddress,omitempty"`
	SetNumber    int64  `protobuf:"bytes,6,rep,name=SetNumber" json:"SetNumber,omitempty"`
	SubCode      string `protobuf:"bytes,6,rep,name=SubCode" json:"SubCode,omitempty"`
}

// 服务端 发布排班
type RespCommonSet struct {
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Msg        string
}

// 修改财务平台的付款账户地址
type ReqChangePayer struct {
	OwnerAddress string `json:"OwnerAddress,omitempty"`
	UserKeyStore string `json:"UserKeyStore,omitempty"`
	UserParse    string `json:"UserParse,omitempty"`
	KeyString    string `json:"KeyString,omitempty"`
	PayerAddress string `json:"PayerAddress,omitempty"`
}
