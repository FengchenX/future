package third_part_pay

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"

	"sub_account_service/finance/blockchain"
	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	"sub_account_service/finance/models"
	"sub_account_service/finance/payment"
	"sub_account_service/finance/payment/alipayModels"
	"sub_account_service/finance/utils"
	svcModels "sub_account_service/number_server/models"
)

type TardingStreamInfo struct {
}

// 保存第三方交易流水
func (this *TardingStreamInfo) SaveTradingStreamInfo(order svcModels.Orders) error {
	//g根据订单编号查询订单信息
	okflag := false

	defer func() {
		if !okflag {
			if err := AddOrderToFailedQueue(order); err != nil {
				glog.Errorln(lib.Log("add order to failed to queue err, trade num: ", order.ThirdTradeNo, "AddOrderToFailedQueue"), "", err)
			}
		}
	}()

	if config.Opts().IsQueryTradeDetails == 0 {
		if err := this.SaveTradingStreamInfoOfNoQueryDetails(order); err != nil {
			glog.Errorln(lib.Log("query ali trade err, trade num: ", order.ThirdTradeNo, "SaveTradingStreamInfoOfNoQueryDetails"), "", err)
			return err
		}
		okflag = true
		return nil
	}
	if order.OrderType == 0 {
		if order.PaymentType == 0 { //支付宝交易类型是支付
			if err := this.SaveTradingStreamInfoOfAliPay(order); err != nil {
				glog.Errorln(lib.Log("query ali trade err, trade num: ", order.ThirdTradeNo, "SaveTradingStreamInfoOfAliPay"), "", err)
				return err
			}
		} else if order.PaymentType == 1 { //支付宝交易类型是转账
			if err := this.SaveTradingStreamInfoOfAliPayTransfer(order); err != nil {
				glog.Errorln(lib.Log("query ali trade err, trade num: ", order.ThirdTradeNo, "QueryAliTrade"), "", err)
				return err
			}
		}

	} else if order.OrderType == 1 {
		//暂不支持其他支付
		return nil
	} else {
		//暂不支持其他支付
		return nil
	}

	okflag = true
	return nil
}

// 保存支付宝交易流水（付款类型： 支付）
func (this *TardingStreamInfo) SaveTradingStreamInfoOfAliPay(order svcModels.Orders) error {
	//g根据订单编号查询订单信息
	var bPayAmount float64
	var iAmount float64
	var pAmount float64
	var rAmount float64
	var tAmount float64
	var Fees float64
	var err error
	var aliTrade *alipayModels.QueryAliTradeResp

	// 共享车服务订单规则
	if strings.Contains(order.ThirdTradeNo, "_") {
		aliTrade, err = payment.QueryAliTrade(strings.Split(order.ThirdTradeNo, "_")[0])
		tAmount, err = utils.KeepTwoDecimalsOfStr(fmt.Sprintf("%f", order.TransferAmount))
		if err != nil {
			return err
		}
	} else {
		aliTrade, err = payment.QueryAliTrade(order.ThirdTradeNo)
		if aliTrade.TotalAmount != "" {
			tAmount, err = utils.KeepTwoDecimalsOfStr(aliTrade.TotalAmount)
			if err != nil {
				return err
			}

		}
	}
	if Fees, err = GetTradingStreamFees(tAmount, order.OrderType); err != nil {
		return err
	}

	if aliTrade.BuyerPayAmount != "" {
		bPayAmount, err = strconv.ParseFloat(aliTrade.BuyerPayAmount, 64)
		if err != nil {
			return err
		}
	}

	if aliTrade.InvoiceAmount != "" {
		iAmount, err = strconv.ParseFloat(aliTrade.InvoiceAmount, 64)
		if err != nil {
			return err
		}
	}

	if aliTrade.PointAmount != "" {
		pAmount, err = strconv.ParseFloat(aliTrade.PointAmount, 64)
		if err != nil {
			return err
		}
	}

	loc, _ := time.LoadLocation("Local")                                                 //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", aliTrade.SendPayDate, loc) //使用模板在对应时区转化为time.time类型
	sPayDate := theTime.Unix()

	if aliTrade.ReceiptAmount != "" {
		rAmount, err = strconv.ParseFloat(aliTrade.ReceiptAmount, 64)
		if err != nil {
			return err
		}
	}

	if aliTrade.TotalAmount != "" {
		tAmount, err = utils.KeepTwoDecimalsOfStr(aliTrade.TotalAmount)
		if err != nil {
			return err
		}

		if Fees, err = GetTradingStreamFees(tAmount, order.OrderType); err != nil {
			return err
		}
	}

	iStatement := models.IncomeStatement{
		OrderType:      int64(order.OrderType),
		AlipaySeller:   config.ZfbConfig().AlipaySeller,
		BuyerLogonID:   aliTrade.BuyerLogonId,
		BuyerPayAmount: bPayAmount,
		BuyerUserID:    aliTrade.BuyerUserId,
		BuyerUserType:  aliTrade.BuyerUserType,
		InvoiceAmount:  iAmount,
		OutTradeNo:     aliTrade.OutTradeNo,
		PointAmount:    pAmount,
		ReceiptAmount:  rAmount,
		SendPayDate:    sPayDate,
		TotalAmount:    tAmount,
		Fees:           Fees,
		PayerName:      order.Company,
		TradeNo:        order.ThirdTradeNo,
		SubAccountNo:   order.SubAccountNo,
		AutoTransfer:   uint8(utils.BoolToInt(!(tAmount > config.Opts().AutoTransferLimit))),
		OnchainNum:     0,
	}

	if res, err := db.RedisClient.Exists(iStatement.TradeNo).Result(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	} else if res != 0 {
		return nil
	}

	out, err := json.Marshal(iStatement)
	if err != nil {
		glog.Errorln(lib.Log("json marshal tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}

	if err := db.RedisClient.Set(iStatement.TradeNo, out, 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}
	if err := db.RedisClient.Set(iStatement.TradeNo+"-chainnotok"+config.Opts().FinanceOrderSvcAppId, "true", 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}
	return nil
}

// 保存支付宝交易流水（付款类型：转账）
func (this *TardingStreamInfo) SaveTradingStreamInfoOfAliPayTransfer(order svcModels.Orders) error {
	//g根据订单编号查询订单信息
	var bPayAmount float64
	var iAmount float64
	var pAmount float64
	var rAmount float64
	var tAmount float64
	var Fees float64
	var err error

	aliTransfer := alipayModels.QueryTransferToAccountResp{}
	err = json.Unmarshal([]byte(order.TransferInfo), &aliTransfer)
	if err != nil {
		glog.Errorln(lib.Log("query ali trade err, trade num: ", order.ThirdTradeNo, "QueryAliTrade"), "", err)
		return err
	}

	//loc, _ := time.LoadLocation("Local")                                            //重要：获取时区
	//theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", order.OrderTime, loc) //使用模板在对应时区转化为time.time类型
	sPayDate := order.OrderTime.Unix()

	tAmount, err = utils.KeepTwoDecimalsOfStr(fmt.Sprintf("%f", order.TransferAmount))
	if err != nil {
		return err
	}
	if Fees, err = GetTradingStreamFees(tAmount, order.OrderType); err != nil {
		return err
	}

	iStatement := models.IncomeStatement{
		OrderType:      int64(order.OrderType),
		AlipaySeller:   config.ZfbConfig().AlipaySeller,
		BuyerPayAmount: bPayAmount,
		InvoiceAmount:  iAmount,
		PointAmount:    pAmount,
		ReceiptAmount:  rAmount,
		SendPayDate:    sPayDate,
		TotalAmount:    tAmount,
		Fees:           Fees,
		PayerName:      order.Company,
		TradeNo:        order.ThirdTradeNo,
		SubAccountNo:   order.SubAccountNo,
		AutoTransfer:   uint8(utils.BoolToInt(!(utils.Uint8ToBool(order.AutoTransfer) && tAmount > config.Opts().AutoTransferLimit))),
		OnchainNum:     0,
	}

	//检查此流水是否已保存
	if res, err := db.RedisClient.Exists(iStatement.TradeNo).Result(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	} else if res != 0 {
		return nil
	}

	out, err := json.Marshal(iStatement)
	if err != nil {
		glog.Errorln(lib.Log("json marshal tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}

	if err := db.RedisClient.Set(iStatement.TradeNo, out, 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}
	if err := db.RedisClient.Set(iStatement.TradeNo+"-chainnotok"+config.Opts().FinanceOrderSvcAppId, "true", 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}

	return nil
}

// 保存交易流水（不获取交易详情）
func (this *TardingStreamInfo) SaveTradingStreamInfoOfNoQueryDetails(order svcModels.Orders) error {
	//g根据订单编号查询订单信息
	var bPayAmount float64
	var iAmount float64
	var pAmount float64
	var rAmount float64
	var tAmount float64
	var Fees float64
	var err error

	//loc, _ := time.LoadLocation("Local")                                            //重要：获取时区
	//theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", order.OrderTime, loc) //使用模板在对应时区转化为time.time类型
	sPayDate := order.OrderTime.Unix()

	//获取手续费
	tAmount, err = utils.KeepTwoDecimalsOfStr(fmt.Sprintf("%f", order.TransferAmount))
	if err != nil {
		return err
	}
	if Fees, err = GetTradingStreamFees(tAmount, order.OrderType); err != nil {
		return err
	}

	iStatement := models.IncomeStatement{
		OrderType:      int64(order.OrderType),
		AlipaySeller:   config.ZfbConfig().AlipaySeller,
		BuyerPayAmount: bPayAmount,
		InvoiceAmount:  iAmount,
		PointAmount:    pAmount,
		ReceiptAmount:  rAmount,
		SendPayDate:    sPayDate,
		TotalAmount:    tAmount,
		Fees:           Fees,
		PayerName:      order.Company,
		TradeNo:        order.ThirdTradeNo,
		SubAccountNo:   order.SubAccountNo,
		AutoTransfer:   uint8(utils.BoolToInt(!(utils.Uint8ToBool(order.AutoTransfer) && tAmount > config.Opts().AutoTransferLimit))),
		OnchainNum:     0,
	}

	//检查此流水是否已保存
	if res, err := db.RedisClient.Exists(iStatement.TradeNo).Result(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	} else if res != 0 {
		return nil
	}

	out, err := json.Marshal(iStatement)
	if err != nil {
		glog.Errorln(lib.Log("json marshal tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}

	if err := db.RedisClient.Set(iStatement.TradeNo, out, 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}
	if err := db.RedisClient.Set(iStatement.TradeNo+"-chainnotok"+config.Opts().FinanceOrderSvcAppId, "true", 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", iStatement.TradeNo, "SaveTradingStreamInfo"), "", err)
	}

	return nil
}

//将处理订单失败的流水加入失败队列中
func AddOrderToFailedQueue(order svcModels.Orders) error {
	out, err := json.Marshal(order)
	if err != nil {
		glog.Errorln(lib.Log("json marshal tarding stream err", order.ThirdTradeNo, "SaveTradingStreamInfo"), "", err)
	}

	if err := db.RedisClient.Set(order.ThirdTradeNo+"-ordernotok"+config.Opts().FinanceOrderSvcAppId, out, 0).Err(); err != nil {
		glog.Errorln(lib.Log("save tarding stream err", order.ThirdTradeNo, "SaveTradingStreamInfo"), "", err)
	}

	return nil
}

//获取交易手续费
func GetTradingStreamFees(amount float64, orderType uint8) (float64, error) {
	var Fees float64
	var err error
	if config.Opts().FeesSwitch == 1 && orderType == 0 { //支付宝转账手续费
		if amount*0.006 > 0.1 {
			Fees, err = strconv.ParseFloat(fmt.Sprintf("%.2f", amount*0.006), 64)
			if err != nil {
				return 0, err
			}
		} else {
			Fees = 0.1
		}
	} else if config.Opts().FeesSwitch == 1 && orderType == 1 { //微信转账手续费
		if amount*0.008 > 0.1 {
			Fees, err = strconv.ParseFloat(fmt.Sprintf("%.2f", amount*0.008), 64)
			if err != nil {
				return 0, err
			}
		} else {
			Fees = 0.1
		}
	} else {

	}

	return Fees, nil
}

// 从第三方支付获取交易流水并保存
func (this *TardingStreamInfo) GetTradingStreamFromThirdPartPay() error {
	orders, err := GetOrderListFromServer()
	if err != nil {
		glog.Errorln(lib.Log("get order list err", "", "GetTradingStreamFromThirdPartPay"), "", err)
		return err
	}

	if len(orders) != 0 {
		glog.Infoln("GetTradingStreamFromThirdPartPay===========================交易流水数量=", len(orders))
	}

	for _, v := range orders {
		if v.OrderState == 1 { // 1为正常交易流水
			if err := this.SaveTradingStreamInfo(v); err != nil {
				glog.Errorln(lib.Log("get order list err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				return err
			}
		} else if v.OrderState == 2 { // 2为退款流水
			if err := CustomerRefundTrade(v); err != nil {
				glog.Errorln(lib.Log("get order list err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				return err
			}
		}

	}

	return nil
}

// 检查获取处理订单失败的流水，重新获取账单详情
func (this *TardingStreamInfo) GetTradingStreamDetailsFromThirdPartPay() error {
	var keys []string
	var err error

	keys, err = db.RedisClient.Keys("*ordernotok" + config.Opts().FinanceOrderSvcAppId).Result()
	if err != nil {
		glog.Errorln(lib.Log("get data from redis err", "", "TradingStreamWriteToBlockchain"), "", err)
		return err
	}
	for _, key := range keys {
		val, err := db.RedisClient.Get(key).Result()
		if err != nil {
			glog.Errorln(lib.Log("get data from redis err", fmt.Sprintf("%v", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}
		order := svcModels.Orders{}
		if err := json.Unmarshal([]byte(val), &order); err != nil {
			glog.Errorln(lib.Log("unmarshal order info err", fmt.Sprintf("%v", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}
		//重新处理订单
		if order.OrderState == 1 { // 1为正常交易流水
			if err := this.SaveTradingStreamInfo(order); err != nil {
				glog.Errorln(lib.Log("get order list err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				continue
			}
		} else if order.OrderState == 2 { // 2为退款流水
			if err := CustomerRefundTrade(order); err != nil {
				glog.Errorln(lib.Log("get order list err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				continue
			}
		}
		if err := db.RedisClient.Del(key).Err(); err != nil {
			glog.Errorln(lib.Log("del redis record err", fmt.Sprintf("%v", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}
	}

	return nil
}

// 检查未上链的流水，进行上链操作并修改状态
func (this *TardingStreamInfo) TradingStreamWriteToBlockchain() error {
	var keys []string
	var key, val string
	var err error

	keys, err = db.RedisClient.Keys("*chainnotok" + config.Opts().FinanceOrderSvcAppId).Result()
	if err != nil {
		glog.Errorln(lib.Log("get tarding stream without write chaincode err", "", "TradingStreamWriteToBlockchain"), "", err)
		return err
	}

	for _, v := range keys {
		key = strings.Split(v, "-chainnotok")[0]
		val, err = db.RedisClient.Get(key).Result()
		if err != nil {
			glog.Errorln(lib.Log("get data from redis err", fmt.Sprintf("%v", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}

		if err := blockchain.WriteBlockchain(val); err != nil {
			glog.Errorln(lib.Log("write blockchain err", key, "TradingStreamWriteToBlockchain"), "", err)
			continue
		}

		if err := db.RedisClient.Del(v).Err(); err != nil {
			glog.Errorln(lib.Log("del tarding stream writing block no ok err", fmt.Sprintf("%v", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}

		if err := db.RedisClient.Set(key+"-chainok"+config.Opts().FinanceOrderSvcAppId, "true", 0).Err(); err != nil {
			glog.Errorln(lib.Log("add tarding stream writing block ok err", fmt.Sprintf("%v-chainok", key), "TradingStreamWriteToBlockchain"), "", err)
			continue
		}
		break
	}

	return nil
}

// 获取订单列表
func GetOrderListFromServer() ([]svcModels.Orders, error) {

	//获取最新版本号
	vUrl := fmt.Sprintf("http://%s:%s/getLatestVersion?appId=%v", config.Opts().FinanceOrderSvcAddress, config.Opts().FinanceOrderSvcPort, config.Opts().FinanceOrderSvcAppId)
	body, err := utils.SendHttpRequest("GET", vUrl, nil, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get order version err", "", "GetOrderListFromServer"), "", err)
		return []svcModels.Orders{}, err
	}

	var vOut map[string]interface{}
	if err := json.Unmarshal(body, &vOut); err != nil {
		glog.Errorln(lib.Log("unmarshal order version err", "", "GetOrderListFromServer"), "", err)
		return []svcModels.Orders{}, err
	}

	if vOut["code"] == nil {
		glog.Errorln(lib.Log("order version code is nil", "", "GetOrderListFromServer"), "", nil)
		return []svcModels.Orders{}, err
	}
	if vOut["code"].(float64) != 0 {
		glog.Errorln(lib.Log("get order version code err", "", "GetOrderListFromServer"), "", fmt.Errorf("code: %v", vOut["code"].(float64)))
		return []svcModels.Orders{}, err
	}

	latestVer := vOut["data"].(map[string]interface{})["latestVersion"].(string)
	if latestVer == "" {
		glog.Infoln("[GetOrderListFromServer]:  latest version is empty")
		return []svcModels.Orders{}, nil
	}

	//根据版本号获取订单列表
	//orderUrl := fmt.Sprintf("http://%s:%s/orders/batch?appId=%v&version=%s", config.Opts().FinanceOrderSvcAddress, config.Opts().FinanceOrderSvcPort, config.Opts().FinanceOrderSvcAppId, "v2")
	orderUrl := fmt.Sprintf("http://%s:%s/orders/batch?appId=%v&version=%s", config.Opts().FinanceOrderSvcAddress, config.Opts().FinanceOrderSvcPort, config.Opts().FinanceOrderSvcAppId, latestVer)
	body, err = utils.SendHttpRequest("GET", orderUrl, nil, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get order list err", "", "GetOrderListFromServer"), "", err)
		return []svcModels.Orders{}, err
	}

	var oOut map[string]interface{}
	if err := json.Unmarshal(body, &oOut); err != nil {
		glog.Errorln(lib.Log("unmarshal order list err", "", "GetOrderListFromServer"), "", err)
		return []svcModels.Orders{}, err
	}

	if oOut["code"].(float64) != 0 {
		glog.Errorln(lib.Log("get order list code err", "", "GetOrderListFromServer"), "", fmt.Errorf("get order status code=%v", oOut["code"].(float64)))
		return []svcModels.Orders{}, err
	}
	res := []svcModels.Orders{}

	for _, v := range oOut["data"].(map[string]interface{})["orders"].([]interface{}) {
		out, _ := json.Marshal(v.(map[string]interface{}))
		order := svcModels.Orders{}
		json.Unmarshal(out, &order)
		res = append(res, order)
		//res = append(res, v.(svcModels.Orders))
	}

	return res, nil
}
