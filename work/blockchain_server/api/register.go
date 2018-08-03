/*--------------------------------------------------------------
*  package: 注册相关的服务
*  time：   2018/04/18
*-------------------------------------------------------------*/
package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/lib/eth"
	"sub_account_service/blockchain_server/model"
)

// 获取以太币余额
func GetEthBalance(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln(lib.Loger("account", "read request body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetEthBalance{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "read request body error",
			})
		return
	}

	req := model.ReqGetEthBalance{}
	if err := json.Unmarshal(body, &req); err != nil {
		glog.Errorln(lib.Loger("account", "unmarshal body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetEthBalance{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "unmarshal body error",
			})
		return
	}

	glog.Infoln(lib.Log("api", req.UserAddress, "获取以太币余额"))
	money, err := myeth.GetBalance(req.UserAddress, config.Opts().EthAddress)
	if err != nil {
		glog.Errorln(lib.Log("api err", req.UserAddress, "获取以太币余额失败"), err)
	}

	resp := &model.RespGetEthBalance{
		StatusCode: uint32(model.Status_Success),
		Balance:    money,
	}

	glog.Infoln(lib.Log("api", req.UserAddress, "获取以太币余额"), resp)
	ctx.JSON(http.StatusOK, resp)
}


func GetAllSchedule(ctx *gin.Context) {

}

func GetTrans(ctx *gin.Context) {
	req := model.ReqGetTrans{}
	ispending := myeth.GT(config.Optional.IpcDir, req.Hash)

	ctx.JSON(http.StatusOK,
		model.RespGetTrans{
			Flag: !ispending,
		})
}

// rpc GetMoney (ReqGetMoney) returns (RespGetMoney) {};
func GetMoney(ctx *gin.Context) {
	//	defer this.RunMiddleWare()

	//	return nil, nil
}
