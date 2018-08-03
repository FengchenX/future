package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"sub_account_service/blockchain_server/lib"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"fmt"
	"flag"
)

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	Port            int    // 端口
	AccAddress      string // 账户地址
	EthAddress      string
	IpcDir          string
	ServerId        string
	ManagerKey      string
	ManagerPhrase   string
	ConfAddress     string
	ConfPort        string
	KeyDir          string
	DeployAddress   string
	PayAddress      string
}

//Optional 默认配置
var Optional *Config

// ConfInstance 获取conf单例
func ConfInstance() *Config {
	if Optional == nil {
		fmt.Println("init cofing ...........................................................")
		var path string
		flag.StringVar(&path, "config", "./conf/config.toml", "config path")
		flag.Parse()
		ParseToml(path)
		return Optional
	}
	return Optional
}

func init() {
	ConfInstance()
}
// Opts 获取配置 以前遗留推荐使用ConfInstance
func Opts() *Config {
	return ConfInstance()
}

// ParseToml 解析配置文件
func ParseToml(file string) error {
	InitToml(file)
	return nil
}

//InitToml 读配置文件
func InitToml(file string) error {
	glog.Infoln(lib.Log("initing", "", "finding config ..."))
	// 如果配置文件不存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(Opts()); err != nil {
			return err
		}
		glog.Infoln("没有找到配置文件，创建新文件 ...")
		return ioutil.WriteFile(file, buf.Bytes(), 0644)
	}
	var conf Config
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		return err
	}
	Optional = &conf
	glog.Infoln(lib.Log("initing", "", "config.Opts()"), Optional)

	return nil
}
