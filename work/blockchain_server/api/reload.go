/*--------------------------------------------------------------
* package: 更新配置相关的服务
* time:    2018/05/10
*-------------------------------------------------------------*/
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/protocol"
)

func ReloadConfig(ctx *gin.Context) {
	req := protocol.CReloadConfig{}
	opt := config.Optional
	opt.KeyDir = req.GetKeyDir()
	opt.ManagerPhrase = req.GetManagerPhrase()
	opt.ManagerKey = req.GetManagerKey()
	opt.IpcDir = req.GetIpcDir()
	opt.EthAddress = req.GetEthAddress()
	opt.AccAddress = req.GetAccAddress()
	opt.Operate_timeout = int(req.GetOperateTimeout())
	glog.Infoln(lib.Log("reload", "", "ReloadConfig"), opt)

	ctx.JSON(http.StatusOK, "")
}

func (this *ApiService) ReloadDeploy(ctx gin.Context, req *protocol.CReloadDeploy) (*protocol.SReloadDeploy, error) {

	return nil, nil
}
