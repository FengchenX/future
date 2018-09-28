package payment

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay"
	"strconv"
	"common-utilities/alibaba/payment/alipayModels"
	"common-utilities/utilities"
	"time"
)

const (
	ALIPAYLOGONID = "ALIPAY_LOGONID"
	ALIPAYUERID   = "ALIPAY_USERID"
)

//查询接口，使用的支付宝的流水单号，这个接口不能够处查询转账给用户的交易流水号
func QueryAliTrade(tradeNo string, alipayClient *AlipayClient) (*alipayModels.QueryAliTradeResp, error) {
	reqBody := alipay.AliPayTradeQuery{
		TradeNo: tradeNo,
	}
	client := alipay.AliPay(*alipayClient)
	resp, err := client.TradeQuery(reqBody)
	log.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Error("TransToAccount Error", resp)
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
func TransferToAccount(model alipayModels.TransferToAccountReq, alipayClient *AlipayClient) (*alipayModels.TransferToAccountResp, error) {
	if model.PayeeType != ALIPAYUERID {
		model.PayeeType = ALIPAYLOGONID
	}
	reqBody := alipay.AliPayFundTransToAccountTransfer{}
	reqBody.OutBizNo = model.OrderId
	reqBody.Amount = strconv.FormatFloat(model.Amount, 'f', 2, 64)
	reqBody.PayeeType = model.PayeeType
	reqBody.PayerShowName = model.PayerShowName
	reqBody.PayeeRealName = model.PayeeRealName
	reqBody.PayeeAccount = model.PayeeAccount
	reqBody.Remark = model.Remark
	client := alipay.AliPay(*alipayClient)
	resp, err := client.FundTransToAccountTransfer(reqBody)

	log.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.Body.SubCode)
	}
	return &alipayModels.TransferToAccountResp{
		OrderId: resp.Body.OutBizNo,
		TradeNo: resp.Body.OrderId,
		PayDate: utilities.ParseDateTimeStrWithDefault(resp.Body.PayDate, time.Now()),
	}, nil
}

//查询转账给用户的交易，使用的是支付宝的交易流水号，非商家的
func QueryTransferToAccount(tradeNo string, alipayClient *AlipayClient) (*alipayModels.QueryTransferToAccountResp, error) {
	reqBody := alipay.AliPayFundTransOrderQuery{
		OrderId: tradeNo,
	}
	client := alipay.AliPay(*alipayClient)
	resp, err := client.FundTransOrderQuery(reqBody)
	log.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		log.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.Body.SubMsg)
	}
	OrderFee, _ := strconv.ParseFloat(resp.Body.OrderFree, 64)

	return &alipayModels.QueryTransferToAccountResp{
		OrderId:        resp.Body.OutBizNo,
		TradeNo:        resp.Body.OrderId,
		Status:         resp.Body.Status,
		PayDate:        utilities.ParseDateTimeStrWithDefault(resp.Body.PayDate, time.Now()),
		ArrivalTimeEnd: utilities.ParseDateTimeStrWithDefault(resp.Body.ArrivalTimeEnd, time.Now()),
		OrderFee:       OrderFee,
	}, nil
}

type AlipayClient alipay.AliPay
func GetAliPayClient(appId,privateKey,publicKey string) *AlipayClient {
	client := alipay.New(
		appId,
		"",
		publicKey,
		privateKey,
		true,
	)
	alipayClient := AlipayClient(*client)
	return &alipayClient
}