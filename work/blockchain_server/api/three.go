package api

import (
	"fmt"
	//"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"

	//"golang.org/x/net/context"
	"math/big"
	"sub_account_service/blockchain_server/contracts"
	"sub_account_service/blockchain_server/lib"
	myeth "sub_account_service/blockchain_server/lib/eth"

	//"sub_account_service/blockchain/protocol"
	"net/http"
	"sub_account_service/blockchain_server/model"

	"github.com/gin-gonic/gin"
)

//上链
func ThreeSetBill(c *gin.Context) {
	//defer this.RunMiddleWare()
	var req model.ReqThreeSetBill
	if err := lib.ParseReq(c, "ThreeSetBill", &req); err != nil {
		return
	}
	glog.Infoln("beginning ThreeSetBill:", "orderId:",req.OrderId, "subCode:",req.ScheduleName,
		"userAddr:",req.UserAddress, "money:",req.Money)
	// 分账
	flag, err := contracts.CheckSubCodeIsOk(DeployAddress, req.ScheduleName)
	if !flag || err != nil {
		glog.Error(fmt.Sprintf("DeployAddress:%v not contain subCode:%v", DeployAddress, req.ScheduleName))
		c.JSON(http.StatusOK, model.RespThreeSetBill{
			StatusCode: uint32(1),
			Msg:        "不存在该排班地址",
		})
		return
	}

	//不存在，那么自己生成
	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.SettleAccounts(req.KeyString, DeployAddress,
		req.ScheduleName, big.NewInt(int64(lib.MoneyIn(req.Money))), req.OrderId)
	if hash == "" {
		glog.Infoln(lib.Log("api", req.UserAddress, "ThreeSetBill Error"), err)
		c.JSON(http.StatusOK, model.RespThreeSetBill{
			StatusCode: uint32(model.Status_ThreeSetBillFail),
			Msg:        "30020 ThreeSetBillFail!",
		})
		return
	}

	if err != nil {
		glog.Infoln(lib.Log("api", req.UserAddress, "ThreeSetBill Error"), err)
		c.JSON(http.StatusOK, model.RespThreeSetBill{
			StatusCode: uint32(model.Status_ThreeSetBillFail),
			Msg:        "30020 ThreeSetBillFail!" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.RespThreeSetBill{
		StatusCode: uint32(model.Status_Success),
		Hash:       hash,
	})
	glog.Infoln("ThreeSetBill success:", "orderId:",req.OrderId, "subCode:",req.ScheduleName,
		"userAddr:",req.UserAddress, "money:",req.Money,"hash:",hash)
}

func ThreeConfirm(c *gin.Context) {
	var req model.ReqConfirm
	var resp model.RespConfirm
	if err := lib.ParseReq(c, "ThreeConfirm", &req); err != nil {
		return
	}
	glog.Infoln("begin ThreeConfirm,orderId:", req.OrderId,"userAddr:",req.UserAddress,"transferDetails:",req.TransferDetails)

	//不存在，那么自己生成
	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.UpdateCalculateLedger(req.KeyString, DeployAddress, req.OrderId,
		req.TransferDetails, common.HexToAddress(req.UserAddress))
	if hash == "" {
		resp = model.RespConfirm{
			StatusCode: uint32(model.Status_ThreeConfirmFail),
			Msg:        "30018 Status_ThreeConfirmFail",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if err != nil {
		glog.Errorln("ThreeConfirm error:",err,"userAddr:",req.UserAddress)
		resp = model.RespConfirm{
			StatusCode: uint32(model.Status_ThreeConfirmFail),
			Msg:        "30018 Status_ThreeConfirmFail!" + err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = model.RespConfirm{
		StatusCode: uint32(model.Status_Success),
		Hash:       hash,
	}
	glog.Infoln("ThreeConfirm success:","userAddr:",req.UserAddress,"hash:",hash)
	c.JSON(http.StatusOK, resp)
}
