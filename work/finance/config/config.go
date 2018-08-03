package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"

	"sub_account_service/finance/lib"
)

// Config 配置类型
type Config struct {
	Operate_timeout        int    // 超时时间设置
	LocalAddress           string // 本机地址
	ApiAddress             string // api地址
	EthAddress             string // Eth地址
	FinanceOrderSvcAddress string // 三方支付地址

	ThreePort           string // 端口
	WebPort             string // 端口
	ApiPort             string // 端口
	FinanceOrderSvcPort string // 端口

	FinanceOrderSvcAppId string // 获取订单需要AppId

	PaySwitch int64 //是否转账开关

	AutoTransferLimit float64 // 自动转账金额限制

	OnchainTimeLimit int64 //重试上链时间

	OnchainNumber int64 //重试上链次数

	FeesSwitch int64 //是否收手续费

	TransferUserLimitMoneyTime int // 定点将累积未转账的money进行转账

	TransferUserLimitMoney float64 // 给余额累加用户转账额度

	IsQueryTradeDetails int //是否查询交易详情

	UserAddress     string
	FinanceAddress  string
	FinanceKeyStore string
	FinancePhrase   string
	ConfAddress     string
	MysqlStr        string

	RedisAddr   string // Redis地址
	RedisPasswd string //Redis密码

	TimeTicker1 time.Duration // 定时从三方支付获取流水信息，存储redis
	TimeTicker2 time.Duration // 定时查询获取详情失败的账单，重新获取账单详情
	TimeTicker3 time.Duration // 定时获取未做上链操作的流水，进行上链操作
	TimeTicker4 time.Duration // 定时获取之前未上链完成的流水，看当前是否已上链
	TimeTicker5 time.Duration // 定时从数据库获取未分账信息，开始分账
	TimeTicker6 time.Duration // 定时查询正在上链的分账信息，是否上链成功
	TimeTicker7 time.Duration // 定时查询上链失败或者未上链的分账信息，重新上链
	TimeTicker8 time.Duration // 定点获取用户累积未转账的money，准备转账
}

// 默认配置
var Optional = Config{}

// Opts 获取配置
func Opts() Config {
	return Optional
}

// ParseToml 解析配置文件
func ParseToml(file string) error {
	glog.Infoln(lib.Log("initing", "", "finding config ..."), file)
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(Opts()); err != nil {
			glog.Infoln("如果配置文件不存在 ...")
			return err
		}
		glog.Infoln("没有找到配置文件，创建新文件 ...")
		return ioutil.WriteFile(file, buf.Bytes(), 0644)
	}
	var conf Config
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		glog.Infoln("DecodeFile Error ...", err)
		return err
	}
	Optional = conf
	glog.Infoln(lib.Log("initing", "", "config.Opts()"), Optional)

	return nil
}
