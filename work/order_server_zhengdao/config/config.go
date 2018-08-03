package config

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"sub_account_service/order_server_zhengdao/lib"
	"time"
)

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	LocalAddress    string // 本机地址
	ApiAddress      string // api地址
	OrderSvcAddress string // 编号服务地址

	ThreePort    string // 端口
	WebPort      string // 端口
	ApiPort      string // 端口
	ConfPort     string // 端口
	OrderSvcPort string // 编号服务端口

	AccAddress    string // 账户地址
	EthAddress    string
	IpcDir        string
	ServerId      string
	ManagerKey    string
	ManagerPhrase string
	KeyDir        string
	ConfAddress   string
	MysqlStr      string

	ReTransferTime int // 定时重新转账时间（小时）

	PaySwitch int // 是否转账开关

	SubAccountNoJson string            // 分账编号信息Json
	SubAccountNoInfo map[string]string // 分账编号信息

	AliPayAddress string // 收款方账号
	AliPayName    string // 收款方用户名

	TimeTicker1 time.Duration // 定时获取未转账成功的订单，向企业用户转账
	TimeTicker2 time.Duration // 从数据库获取 添加到编号系统失败 的交易流水，重新添加
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

	if err := json.Unmarshal([]byte(Optional.SubAccountNoJson), &Optional.SubAccountNoInfo); err != nil {
		return err
	}
	glog.Infoln(lib.Log("initing", "", "config.Opts()"), Optional)

	return nil
}
