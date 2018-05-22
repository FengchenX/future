package rpc

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"time"
	"znfz/conf_server/config"
	"znfz/conf_server/db"
	"znfz/conf_server/protocol"
	"znfz/conf_server/lib"
)

type ConfigServer struct{}

func (this *ConfigServer) RegisterConfig() {

}

//  设置配置文件
func (this *ConfigServer) SetConfig(ctx context.Context, req *protocol.ReqSetConfig) (*protocol.RespSetConfig, error) {
	glog.Infoln((lib.Log("conf", "", "SetConfig")), req)
	config.Optional.KeyDir = req.GetKeyDir()
	config.Optional.ManagerPhrase = req.GetManagerPhrase()
	config.Optional.ManagerKey = req.GetManagerKey()
	config.Optional.ServerId = req.GetServerId()
	config.Optional.IpcDir = req.GetIpcDir()
	config.Optional.EthAddress = req.GetEthAddress()
	config.Optional.AccAddress = req.GetAccAddress()
	config.Optional.Port = req.GetPort()
	config.Optional.LocalAddress = req.GetLocalAddress()
	config.Optional.Operate_timeout = int(req.GetOperateTimeout())
	return &protocol.RespSetConfig{
		StaticCode: 0,
	}, nil
}

// 获取配置文件
func (this *ConfigServer) GetConfig(ctx context.Context, req *protocol.ReqGetConfig) (*protocol.RespGetConfig, error) {
	glog.Infoln(lib.Log("conf", req.GetAddress(), "GetConfig"))

	deploy := &db.HistoryDeploy{}
	dberr := db.DbClient.Client.Order("create_time, desc").Find(deploy)
	if dberr.Error != nil {
		glog.Errorln(dberr.Error)
	}

	dberr = db.DbClient.Client.Create(&db.RpcHistory{
		CreateTime: time.Now(),
		Address:    req.GetAddress(),
		Version:    deploy.Version,
		SubVersion: deploy.SubVersion,
	})
	if dberr.Error != nil {
		glog.Errorln(dberr.Error)
	}

	return &protocol.RespGetConfig{
		OperateTimeout: uint32(config.Opts().Operate_timeout),
		LocalAddress:   config.Optional.LocalAddress,
		Port:           config.Optional.Port,
		AccAddress:     config.Optional.AccAddress,
		EthAddress:     config.Optional.EthAddress,
		IpcDir:         config.Optional.IpcDir,
		ServerId:       config.Optional.ServerId,
		ManagerKey:     config.Optional.ManagerKey,
		ManagerPhrase:  config.Optional.ManagerPhrase,
		KeyDir:         config.Optional.KeyDir,
	}, nil
}
