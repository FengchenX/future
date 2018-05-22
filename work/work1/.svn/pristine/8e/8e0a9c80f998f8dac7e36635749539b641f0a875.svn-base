package config

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"znfz/server/lib"
	"google.golang.org/grpc"
	"znfz/server/protocol"
	"golang.org/x/net/context"
)

// Config 配置类型
type Config struct {
	Operate_timeout int    // 超时时间设置
	LocalAddress    string // 本机地址
	Port            string // 端口
	AccAddress      string // 账户地址
	EthAddress      string
	IpcDir          string
	ServerId        string
	ManagerKey      string
	ManagerPhrase   string
	ConfAddress     string
	ConfPort        string
	KeyDir          string
}

// 默认配置
var Optional = Config{}

// Opts 获取配置
func Opts() Config {
	return Optional
}

// ParseToml 解析配置文件
func ParseToml(file string) error {
	InitToml(file)
	//GetFromConfig()
	return nil
}

func GetFromConfig(){
	glog.Infoln(lib.Log("config", "", "ParseToml1"), Opts().ConfAddress+Opts().ConfPort)
	conn, err := grpc.Dial(Opts().ConfAddress+Opts().ConfPort, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		glog.Errorln(lib.Log("config err", "", "ParseToml2"), err)
	}
	client := protocol.NewConfServerClient(conn)
	resp, err := client.GetConfig(context.Background(), &protocol.ReqGetConfig{
		Address: Opts().LocalAddress + Opts().Port,
	})

	if err != nil {
		glog.Errorln(lib.Log("config err", "", "ParseToml3"), err)
		return
	}
	if resp.GetStaticCode()!=0{
		glog.Errorln(lib.Log("config err", "", "ParseToml4"), "get config fail")
		return
	}
	Optional.KeyDir = resp.GetKeyDir()
	Optional.ManagerPhrase = resp.GetManagerPhrase()
	Optional.ManagerKey = resp.GetManagerKey()
	Optional.IpcDir = resp.GetIpcDir()
	Optional.EthAddress = resp.GetEthAddress()
	Optional.AccAddress = resp.GetAccAddress()
	Optional.Operate_timeout = int(resp.GetOperateTimeout())
	glog.Infoln(lib.Log("config", "", "ParseToml5"),Opts())
}

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
	Optional = conf
	glog.Infoln(lib.Log("initing", "", "config.Opts()"), Optional)

	return nil
}
