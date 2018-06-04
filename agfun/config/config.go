package config

import (
	"errors"
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
)
const file = "./conf/config.toml"
func init() {
	if err := ParseToml(file); err != nil {
		glog.Fatalln(err)
	}
}

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	LocalAddress    string // 本机地址
	ApiAddress      string // api地址
	ThreePort       string // 端口
	WebPort         string // 端口
	ApiPort         string // 端口
	ConfPort        string // 端口
	AccAddress      string // 账户地址
	EthAddress      string
	IpcDir          string
	ServerId        string
	ManagerKey      string
	ManagerPhrase   string
	KeyDir          string
	ConfAddress     string
	MysqlStr        string
}

// 默认配置
var conf Config

//Conf  获取配置
func Conf() Config {
	return conf
}

//ParseToml 解析配置文件
func ParseToml(file string) error {
	glog.Infoln("ParseToml********parse start")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		//配置文件不存在
		var b []byte
		buf := bytes.NewBuffer(b)
		if err := toml.NewEncoder(buf).Encode(conf); err != nil {
			glog.Infoln("ParseToml********编码错误 ")
			return err
		}
		ioutil.WriteFile(file, buf.Bytes(), os.ModeType)
		return errors.New("配置文件不存在")
	}	
	var temp Config
	if _, err := toml.DecodeFile(file, &temp); err != nil {
		glog.Errorln("ParseToml******解析配置文件错误", err)
		return err
	}
	conf = temp
	glog.Infoln("ParseToml**********success")
	return nil
}
