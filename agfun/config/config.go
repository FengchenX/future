package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"znfz/three_server/lib"
)

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	LocalAddress    string // 本机地址
	ApiAddress      string // api地址

	ThreePort string // 端口
	WebPort   string // 端口
	ApiPort   string // 端口
	ConfPort  string // 端口

	AccAddress    string // 账户地址
	EthAddress    string
	IpcDir        string
	ServerId      string
	ManagerKey    string
	ManagerPhrase string
	KeyDir        string
	ConfAddress   string
	MysqlStr      string
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
