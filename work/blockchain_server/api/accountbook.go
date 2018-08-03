package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"

	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/contracts"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/model"
)

// 订单查询服务
func GetAccountBook(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln(lib.Loger("account", "read request body error"))
		ctx.JSON(http.StatusOK,
			model.RespGetAccountBook{
				StatusCode: uint32(model.Status_GetAccountBookFail),
				Msg:        "read request body error",
			})
		return
	}

	req := model.ReqGetAccountBook{}
	if err := json.Unmarshal(body, &req); err != nil {
		glog.Errorln(lib.Loger("account", "unmarshal body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetAccountBook{
				StatusCode: uint32(model.Status_GetAccountBookFail),
				Msg:        "unmarshal body error",
			})
		return
	}
	glog.Info("beginning GetAccountBook:","userAddr:",req.UserAddress,"orderId:",req.OrderId)
	l, err := contracts.GetOneLedgerCxt(config.Opts().DeployAddress, req.OrderId, common.HexToAddress(req.UserAddress))
	if err != nil {
		glog.Infoln(lib.Loger("api", "GetAccountBook Error"), err)
		ctx.JSON(http.StatusOK,
			model.RespGetAccountBook{
				StatusCode: uint32(model.Status_GetAccountBookFail),
				Msg:        "30017 GetAccountBookFail" + err.Error(),
			})
		return
	}

	resp := &model.RespGetAccountBook{
		StatusCode:      uint32(model.Status_Success),
		OrderId:         l.OrderId,                                  // 分账的交易信息编号
		Money:           lib.MoneyOut(float64(l.Calculate.Int64())), // 这笔交易分到的钱
		Rflag:           l.Rflag,                                    // 是否更新过打钱后的交易信息，false为未更新，true为已更新。
		TransferDetails: l.TransferDetails,
	}
	glog.Info("GetAccountBook success:","userAddr:",req.UserAddress,"orderId:",req.OrderId,"resp:",resp)
	ctx.JSON(http.StatusOK, resp)
}

// 根据orderid查询账本
func GetABById(ctx *gin.Context) {
	//glog.Infoln(lib.Loger("api", "根据orderid查询账本 req"), req)
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln(lib.Loger("account", "read request body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetAccountBook{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "read request body error",
			})
		return
	}

	req := model.ReqGetABById{}
	if err := json.Unmarshal(body, &req); err != nil {
		glog.Errorln(lib.Loger("account", "unmarshal body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetAccountBook{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "unmarshal body error",
			})
		return
	}

	// 获取排班
	_, radio, subWays, _, _, err := contracts.GetDistributionRatioByCode(config.Opts().DeployAddress, req.ScheduleName)
	if err != nil {
		glog.Errorln("[api err] 根据orderid查询账本", err)
		ctx.JSON(http.StatusOK,
			model.RespGetABById{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "30001 SetSchedueFail" + err.Error(),
			})
		return
	}

	//newAcco := lib.ParseStrArr(accos)
	//newAcco, _ := contracts.GetgetLedgerSubAddrs(config.Opts().DeployAddress, req.OrderId)

	_,newAcco,err := contracts.GetSchedulingCxt(config.Opts().DeployAddress,req.ScheduleName)
	if err != nil {
		glog.Errorln("[api err] GetSchedulingCxt fail:", err)
		ctx.JSON(http.StatusOK,
			model.RespGetABById{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "30001 GetSchedueFail" + err.Error(),
			})
		return
	}

	accMap := GetAccountsBySchedule(newAcco)
	reabs := getAccountBook(newAcco, radio, subWays, req.OrderId, accMap)

	if len(reabs) == 0 {
		glog.Errorln("[api err] 根据orderid查询账本 :  len(reabs)==0")
		ctx.JSON(http.StatusOK,
			model.RespGetABById{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "30001 SetSchedueFail len(reabs) == 0",
			})
		return
	}

	resp := &model.RespGetABById{
		StatusCode: uint32(model.Status_Success),
		Msg:        "get acount book success",
		Abs:        reabs,
	}
	glog.Infoln(lib.Loger("api", "根据orderid查询账本 resp"), "resp:", resp)
	ctx.JSON(http.StatusOK, resp)
}

// 根据sheduleid查询账本
func GetABBySh(ctx *gin.Context) {
	// 获取排班
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln(lib.Loger("account", "read request body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetABBySh{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "read request body error",
			})
		return
	}

	req := model.ReqGetABBySh{}
	if err := json.Unmarshal(body, &req); err != nil {
		glog.Errorln(lib.Loger("account", "unmarshal body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetABBySh{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "unmarshal body error",
			})
		return
	}

	glog.Infoln(lib.Loger("api", "根据sheduleid查询账本 req"), "req:", req)

	accos, radio, subWays, _, _, err := contracts.GetDistributionRatioByCode(config.Opts().DeployAddress, req.ScheduleName)
	if err != nil {
		glog.Errorln("[api err] 查询排班接口失败", err)
		ctx.JSON(http.StatusOK,
			model.RespGetABBySh{
				StatusCode: uint32(model.Status_SetSchedueFail),
				Msg:        "30001 SetSchedueFail" + err.Error(),
			})
		return
	}

	newAcco := lib.ParseStrArr(accos)

	reOids := make([]*model.OrderDliver, 0)
	accMap := GetAccountsBySchedule(newAcco)
	for _, oid := range req.OrderId {
		reabs := getAccountBook(newAcco, radio, subWays, oid, accMap)
		reOids = append(reOids, &model.OrderDliver{
			OrderId: oid,
			Abs:     reabs,
		})
	}
	resp := &model.RespGetABBySh{
		StatusCode: uint32(model.Status_Success),
		Ods:        reOids,
	}
	glog.Infoln(lib.Loger("api", "根据sheduleid查询账本 resp"), "resp:", resp)
	ctx.JSON(http.StatusOK, resp)
}

func getAccountBook(accos []string, radio []*big.Int, subWays []*big.Int, order_id string,
	accMap map[string]*model.UserAccount) []*model.AccountBook {
	reabs := make([]*model.AccountBook, 0)

	//glog.Infoln("getAccountBook enter order_id:", order_id, "accos:", accos,"radio:", radio)

	for i, addr := range accos {
		ab, err := contracts.GetOneLedgerCxt(config.Opts().DeployAddress, order_id, common.HexToAddress(addr))
		if err != nil {
			glog.Infoln(lib.Log("api", addr, "GetAccountBook Error"), err)
			continue
		}
		if i > len(radio)-1 {
			glog.Infoln(lib.Log("api", addr, "GetAccountBook Error : i >len(radio)-1"))
			continue
		}
		if i > len(subWays)-1 {
			glog.Infoln(lib.Log("api", addr, "GetAccountBook Error : i >len(subWays)-1"))
			continue
		}
		if radio[i] == nil {
			glog.Infoln(lib.Log("api", addr, "GetAccountBook Error : radio[i]==nil"))
			continue
		}
		if ab.OrderId == "" {
			//glog.Infoln(lib.Log("api", addr, "GetAccountBook Error : ab.OrderId == nil"))
			continue
		}
		var acco *model.UserAccount
		if v, exist := accMap[addr]; exist {
			acco = v
		}

		ra := float64(radio[i].Int64())
		if subWays[i].Int64() == int64(1) {
			ra /= 10000
		} else {
			ra = lib.RadioOut(ra)
		}
		acc_book := &model.AccountBook{
			Address:         addr,
			OrderId:         ab.OrderId,                                  // 分账的交易信息编号
			Money:           lib.MoneyOut(float64(ab.Calculate.Int64())), // 这笔交易分到的钱
			Rflag:           ab.Rflag,
			TransferDetails: ab.TransferDetails, // 是否更新过打钱后的交易信息，false为未更新，true为已更新。
			Radio:           ra,
			Acco:            acco,
			SubWay:          subWays[i].Int64(),
		}
		glog.Infoln("getAccountBook for order_id:", order_id, "addr:", addr, "acc_book:", acc_book, "radio:", radio[i].Int64())

		reabs = append(reabs, acc_book)
	}
	return reabs
}
