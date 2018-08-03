package config

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"sub_account_service/app_server_v2/lib"

	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

func init() {
	ConfInst()
}

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	LocalAddress    string // 本机地址
	LocalPort       string // 端口
	ApiAddress      string
	ApiPort         string
	Mysql           string
	KeyStore        string
	Phrase          string
	CheckLimit      int
	UpdateDate      int
	OrderAddress    string
	AppId           string
	DumpNum         int
	OrderNum        int
	GetBookTime     int
	ResetQuoTime    int
}

//Optional 默认配置
var Optional *Config

//ConfInst 配置单例
func ConfInst() *Config {
	if Optional == nil {
		var path string
		flag.StringVar(&path, "config", "./conf/config.toml", "config path")
		flag.Parse()
		ParseToml(path) // 初始化配置
		return Optional
	}
	return Optional
}

// Opts 获取配置 废弃的方法推荐使用ConfInstance
func Opts() *Config {
	return ConfInst()
}

// ParseToml 解析配置文件
func ParseToml(file string) error {
	InitToml(file)
	//GetFromConfig()
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
