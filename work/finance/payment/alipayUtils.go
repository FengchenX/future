package payment

import (
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"github.com/smartwalle/alipay"
	"strconv"
	"sub_account_service/finance/config"
	"sub_account_service/finance/lib"
	"sub_account_service/finance/payment/alipayModels"
)

const (
	ALIPAYLOGONID = "ALIPAY_LOGONID"
	ALIPAYUERID   = "ALIPAY_USERID"
)

//查询接口，使用的支付宝的流水单号，这个接口不能够处查询转账给用户的交易流水号
func QueryAliTrade(tradeNo string) (*alipayModels.QueryAliTradeResp, error) {
	client := getAliPayClient()
	reqBody := alipay.AliPayTradeQuery{
		TradeNo: tradeNo,
	}
	resp, err := client.TradeQuery(reqBody)
	glog.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		glog.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.AliPayTradeQuery.SubMsg)
	}
	bytes, err := json.Marshal(resp.AliPayTradeQuery)
	if err != nil {
		return nil, err
	}
	returnResp := &alipayModels.QueryAliTradeResp{}
	err = json.Unmarshal(bytes, returnResp)
	if err != nil {
		return nil, err
	}
	return returnResp, nil
}

//转账给用户
func TransferToAccount(model alipayModels.TransferToAccountReq) (*alipayModels.TransferToAccountResp, error) {
	if model.PayeeType != ALIPAYUERID {
		model.PayeeType = ALIPAYLOGONID
	}
	client := getAliPayClient()
	reqBody := alipay.AliPayFundTransToAccountTransfer{}
	reqBody.OutBizNo = model.OrderId
	reqBody.Amount = strconv.FormatFloat(model.Amount, 'f', 2, 64)
	reqBody.PayeeType = model.PayeeType
	reqBody.PayerShowName = model.PayerShowName
	reqBody.PayeeRealName = model.PayeeRealName
	reqBody.PayeeAccount = model.PayeeAccount
	reqBody.Remark = model.Remark
	resp, err := client.FundTransToAccountTransfer(reqBody)

	glog.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		glog.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.Body.SubMsg)
	}
	return &alipayModels.TransferToAccountResp{
		OrderId: resp.Body.OutBizNo,
		TradeNo: resp.Body.OrderId,
		PayDate: lib.ParseDateTimeStr(resp.Body.PayDate),
	}, nil
}

//查询转账给用户的交易，使用的是支付宝的交易流水号，非商家的
func QueryTransferToAccount(tradeNo string) (*alipayModels.QueryTransferToAccountResp, error) {
	client := getAliPayClient()
	reqBody := alipay.AliPayFundTransOrderQuery{
		OrderId: tradeNo,
	}
	resp, err := client.FundTransOrderQuery(reqBody)
	glog.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		glog.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.Body.SubMsg)
	}
	OrderFee, _ := strconv.ParseFloat(resp.Body.OrderFree, 64)
	return &alipayModels.QueryTransferToAccountResp{
		OrderId:        resp.Body.OutBizNo,
		TradeNo:        resp.Body.OrderId,
		Status:         resp.Body.Status,
		PayDate:        lib.ParseDateTimeStr(resp.Body.PayDate),
		ArrivalTimeEnd: lib.ParseDateTimeStr(resp.Body.ArrivalTimeEnd),
		OrderFee:       OrderFee,
	}, nil
}

func getAliPayClient() *alipay.AliPay {
	client := alipay.New(
		config.ZfbConfig().AlipayAppID,
		"",
		config.ZfbConfig().AlipayPublicKey,
		config.ZfbConfig().AlipayPrivateKey,
		true,
	)
	return client
}

func isSuccessCode(code string) bool {
	return code == "10000"
}
