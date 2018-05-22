/*--------------------------------------------------------------
* package: 更新配置相关的服务
* time:    2018/05/10
*-------------------------------------------------------------*/
package apiserver

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/protocol"
)

func (this *ApiService) ReloadConfig(ctx context.Context, req *protocol.CReloadConfig) (*protocol.SReloadConfig, error) {
	opt := config.Optional
	opt.KeyDir = req.GetKeyDir()
	opt.ManagerPhrase = req.GetManagerPhrase()
	opt.ManagerKey = req.GetManagerKey()
	opt.IpcDir = req.GetIpcDir()
	opt.EthAddress = req.GetEthAddress()
	opt.AccAddress = req.GetAccAddress()
	opt.Operate_timeout = int(req.GetOperateTimeout())
	glog.Infoln(lib.Log("reload", "", "ReloadConfig"), opt)
	return nil, nil
}

func (this *ApiService) ReloadDeploy(ctx context.Context, req *protocol.CReloadDeploy) (*protocol.SReloadDeploy, error) {
	this.AccountAddress = req.Account
	this.AddressManager = req.Address
	glog.Infoln(lib.Log("reload", "", "ReloadDeploy"), this.AddressManager,this.AccountAddress)
	return nil, nil
}
