package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"sub_account_service/blockchain_server/arguments"
	"sub_account_service/blockchain_server/config"
	"sub_account_service/blockchain_server/contracts"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/model"
	"sub_account_service/blockchain_server/lib/eth"
)

func SetAccount(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	err_result := ""
	if err != nil {
		err_result = "read request body error:" + err.Error()
		glog.Errorln(lib.Loger("account",err_result ))
		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        err_result,
			})
		return
	}
	fmt.Println("1111111111111111111111111111111111111111111111111", string(body))
	req := model.ReqSetAccount{}
	if err := json.Unmarshal(body, &req); err != nil {
		err_result = "unmarshal body error:" + err.Error()
		glog.Errorln(lib.Loger("account", err_result))

		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        err_result,
			})
		return
	}

	glog.Infoln(lib.Loger("api", "SetAccount"), req)

	if req.UserAccount == nil {
		glog.Errorln(lib.Loger("account", "SetAccount error req.GetUserAccount() == nil"))

		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        "10001 SetAccountFail! req.GetUserAccount() == nil",
			})
		return
	}

	if req.UserAccount.Name == "" {
		glog.Errorln(lib.Loger("account", "SetAccount error req.GetUserAccount().Name == "))

		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        "10002 SetAccountFail! req.GetUserAccount() == nil",
			})
		return
	}

	if req.UserAccount.Alipay == "" {
		glog.Errorln(lib.Loger("account", "SetAccount error req.GetUserAccount().Alipay == "))

		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        "10003 SetAccountFail! req.GetUserAccount().Alipay == nil",
			})
		return
	}
	fmt.Println("2222222222222222222222222222", req.KeyString, config.Opts().DeployAddress, *req.UserAccount)
	if req.KeyString == "" && req.UserKeyStore != "" {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}
	hash, err := contracts.BindAccountInfos(req.KeyString,
		config.Opts().DeployAddress,
		arguments.AccountArguments{
			AccountAddr: req.UserAccount.Address,
			Name:        req.UserAccount.Name,
			BankCard:    req.UserAccount.BankCard,
			WeChat:      req.UserAccount.WeChat,
			Alipay:      req.UserAccount.Alipay,
			Telephone:   req.UserAccount.Telephone,
		})
	if err != nil || hash == "" {
		glog.Errorln(lib.Log("account", "", "SetAccount error"), err)
		ctx.JSON(http.StatusOK,
			model.RespSetAccount{
				StatusCode: ACCOUNT_APPLY_NIL,
				Msg:        "10005 SetAccountFail!" + err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		model.RespSetAccount{
			StatusCode: uint32(model.Status_Success),
			Hash:		hash,
			Msg:        "Success",
		})
}

func GetAccount(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		glog.Errorln(lib.Loger("account", "read request body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetAccount{
				StatusCode: ACCOUNT_GET_ERR,
				Msg:        "read request body error",
			})
		return
	}

	req := model.ReqGetAccount{}
	if err := json.Unmarshal(body, &req); err != nil {
		glog.Errorln(lib.Loger("account", "unmarshal body error"))

		ctx.JSON(http.StatusOK,
			model.RespGetAccount{
				StatusCode: ACCOUNT_GET_ERR,
				Msg:        "unmarshal body error",
			})
		return
	}

	glog.Infoln(lib.Log("api", req.UserAddress, "GetAccount"), req)

	acc, err := contracts.GetAccountCxt(config.Opts().DeployAddress, common.HexToAddress(req.UserAddress))
	if err != nil {
		glog.Errorln(lib.Log("account", req.UserAddress, "SetAccount error"), err)

		ctx.JSON(http.StatusOK,
			model.RespGetAccount{
				StatusCode: ACCOUNT_GET_ERR,
				Msg:        "30015 SetAccountFail!" + err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		model.RespGetAccount{
			StatusCode: uint32(model.Status_Success),
			UserAccount: &model.UserAccount{
				Address:   req.UserAddress,
				Name:      acc.Name,
				BankCard:  acc.BankCard,
				WeChat:    acc.WeChat,
				Alipay:    acc.Alipay,
				Telephone: acc.Telephone},
			Msg: "Success",
		})
}
