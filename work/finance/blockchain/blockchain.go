package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	"sub_account_service/finance/models"
	"sub_account_service/finance/protocol"
	"sub_account_service/finance/utils"
)

// 从redis取出的交易流水，进行上链操作
//
func WriteBlockchain(val string) error {
	inState := models.IncomeStatement{}
	if err := json.Unmarshal([]byte(val), &inState); err != nil {
		glog.Errorln(lib.Log("unmarshal incomestatement err", "", "WriteBlockchain"), "", err)
		return err
	}

	//先将交易记录存储数据库，避免重复写区块链操作
	//
	inStatetmp := models.IncomeStatement{TradeNo: inState.TradeNo}
	if dberr := db.DbClient.Client.Where(&inStatetmp).First(&inStatetmp); dberr.Error == nil {
		return nil
	}

	//如果排班表有退款，走带退款的上链流程
	if v, ok := models.ScheduleAccountArrears[inState.SubAccountNo]; ok && v != 0 {
		if err := WriteBlockchainWithArrears(inState); err != nil {
			glog.Errorln(lib.Log("write block chain with arrears err", "", "WriteBlockchainWithArrears"), "", err)
			return err
		}
		return nil
	}

	// 交易流水金额为0时，不再上链，直接保存数据库记录
	if inState.TotalAmount-inState.Fees == float64(0) {
		if err := SaveTradeInfoToDatabase(inState); err != nil {
			glog.Errorln(lib.Log("save trade info to database err", "", "SaveTradeInfoToDatabase"), "", err)
			return err
		}
		return nil
	}

	// 初始化交易秘钥，准备上链
	//
	keyStr := utils.ParseKeyStore()
	if keyStr == "" {
		glog.Errorln(lib.Log("get key store err", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("get key store err: key string is empty")
	}

	// 上链数据结构初始化
	//
	reqg := &protocol.ReqThreeSetBill{
		UserAddress:  config.Opts().FinanceAddress,
		KeyString:    keyStr,
		ScheduleName: inState.SubAccountNo,
		Money:        inState.TotalAmount - inState.Fees,
		OrderId:      inState.TradeNo,
	}

	out, err := json.Marshal(reqg)
	if err != nil {
		glog.Errorln(lib.Log("json marshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("json marshal error: %v", err)
	}

	// 存证上链太慢了，800ms+，简易先回包
	//
	//RespThreeSetBill, err := client.Cli.C.ThreeSetBill(context.Background(), reqg)
	reqBody := bytes.NewReader(out)
	reqUrl := fmt.Sprintf("http://%s:%s/threesetbill", config.Optional.ApiAddress, config.Optional.ApiPort)
	rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("write blockchain err", reqg.OrderId, "WriteBlockchain"), "", err, reqg.ScheduleName)
		return err
	}

	RespThreeSetBill := protocol.RespThreeSetBill{}
	if err := json.Unmarshal(rspBody, &RespThreeSetBill); err != nil {
		glog.Errorln(lib.Log("json unmarshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("json unmarshal error: %v", err)
	}

	if RespThreeSetBill.Hash == "" || RespThreeSetBill.StatusCode != uint32(0) {
		glog.Errorln(lib.Log("RespThreeSetBill.GetHash == nil or StatusCode!=uint32(0)", reqg.OrderId, "WriteBlockchain"), RespThreeSetBill.StatusCode, RespThreeSetBill.Msg, reqg.ScheduleName)
		return fmt.Errorf("RespThreeSetBill.GetHash == nil or StatusCode!=uint32(0)")
	}

	// 交易流水记录存储sql
	//
	inState.BlockchainStatus = models.NotOnchain
	inState.OnchainTime = time.Now().Unix()
	inState.Hash = RespThreeSetBill.Hash
	if dberr := db.DbClient.Client.Create(&inState); dberr.Error != nil {
		glog.Errorln(lib.Log("create database record err", reqg.OrderId, "WriteBlockchain"), "", dberr.Error, reqg.ScheduleName)
		return err
	}

	return nil
}

//排班有退款，要带退款的上链流程
func WriteBlockchainWithArrears(inState models.IncomeStatement) error {
	flag := false
	tx := db.DbClient.Client.Begin()
	defer func() {
		if !flag {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 交易流水金额为0时，不再上链，直接保存数据库记录
	if inState.TotalAmount < inState.Fees+models.GetScheduleAccount(inState.SubAccountNo) {
		if err := SaveTradeInfoToDatabaseWithArrears(inState); err != nil {
			glog.Errorln(lib.Log("save trade info to database err", "", "SaveTradeInfoToDatabase"), "", err)
			return err
		}
		return nil
	}

	// 初始化交易秘钥，准备上链
	//
	keyStr := utils.ParseKeyStore()
	if keyStr == "" {
		glog.Errorln(lib.Log("get key store err", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("get key store err: key string is empty")
	}

	// 上链数据结构初始化
	//
	reqg := &protocol.ReqThreeSetBill{
		UserAddress:  config.Opts().FinanceAddress,
		KeyString:    keyStr,
		ScheduleName: inState.SubAccountNo,
		Money:        inState.TotalAmount - inState.Fees - models.GetScheduleAccount(inState.SubAccountNo),
		OrderId:      inState.TradeNo,
	}

	out, err := json.Marshal(reqg)
	if err != nil {
		glog.Errorln(lib.Log("json marshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("json marshal error: %v", err)
	}

	// 存证上链太慢了，800ms+，简易先回包
	//
	//RespThreeSetBill, err := client.Cli.C.ThreeSetBill(context.Background(), reqg)
	reqBody := bytes.NewReader(out)
	reqUrl := fmt.Sprintf("http://%s:%s/threesetbill", config.Optional.ApiAddress, config.Optional.ApiPort)
	rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("write blockchain err", reqg.OrderId, "WriteBlockchain"), "", err, reqg.ScheduleName)
		return err
	}

	RespThreeSetBill := protocol.RespThreeSetBill{}
	if err := json.Unmarshal(rspBody, &RespThreeSetBill); err != nil {
		glog.Errorln(lib.Log("json unmarshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
		return fmt.Errorf("json unmarshal error: %v", err)
	}

	if RespThreeSetBill.Hash == "" || RespThreeSetBill.StatusCode != uint32(0) {
		glog.Errorln(lib.Log("RespThreeSetBill.GetHash == nil or StatusCode!=uint32(0)", reqg.OrderId, "WriteBlockchain"), RespThreeSetBill.StatusCode, RespThreeSetBill.Msg, reqg.ScheduleName)
		return fmt.Errorf("RespThreeSetBill.GetHash == nil or StatusCode!=uint32(0)")
	}

	// 更新排班欠款
	scheAccount := models.ScheduleAccount{ScheduleName: inState.SubAccountNo}
	if dberr := tx.Model(&scheAccount).Update("arrears", 0); dberr.Error != nil {
		glog.Errorln(lib.Log("update schedule arrears err", inState.TradeNo, "WriteBlockchain"), "", nil)
		return dberr.Error
	}
	models.ScheduleAccountArrears[inState.SubAccountNo] = 0
	delete(models.ScheduleAccountArrears, inState.SubAccountNo)

	// 交易流水记录存储sql
	//
	inState.BlockchainStatus = models.NotOnchain
	inState.OnchainTime = time.Now().Unix()
	inState.Hash = RespThreeSetBill.Hash
	inState.Arrears = models.GetScheduleAccount(inState.SubAccountNo)
	if dberr := db.DbClient.Client.Create(&inState); dberr.Error != nil {
		glog.Errorln(lib.Log("create database record err", reqg.OrderId, "WriteBlockchain"), "", dberr.Error, reqg.ScheduleName)
		return err
	}

	flag = true
	return nil
}

// 交易流水金额为0时，不再上链，直接保存数据库记录
func SaveTradeInfoToDatabase(inState models.IncomeStatement) error {
	inState.BlockchainStatus = models.Onchain
	if dberr := db.DbClient.Client.Create(&inState); dberr.Error != nil {
		glog.Errorln(lib.Log("create database record err", inState.TradeNo, "WriteBlockchain"), "", dberr.Error)
	}

	epBill := models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("trade_no = ?", inState.TradeNo).First(&epBill); dberr.Error == nil {
		// gorm 查询未判断error
	} else {
		epBill = models.ExpensesBill{
			OrderType:      inState.OrderType,
			AlipaySeller:   inState.AlipaySeller,
			BuyerLogonID:   inState.BuyerLogonID,
			BuyerPayAmount: inState.BuyerPayAmount,
			BuyerUserID:    inState.BuyerUserID,
			BuyerUserType:  inState.BuyerUserType,
			InvoiceAmount:  inState.InvoiceAmount,
			OutTradeNo:     inState.OutTradeNo,
			PointAmount:    inState.PointAmount,
			ReceiptAmount:  inState.ReceiptAmount,
			SendPayDate:    inState.SendPayDate,
			TotalAmount:    inState.TotalAmount,
			Fees:           inState.Fees,
			Arrears:        inState.Arrears,
			TradeNo:        inState.TradeNo,
			SubAccountNo:   inState.SubAccountNo,
			AutoTransfer:   inState.AutoTransfer,
			TradeStatus:    models.TransferOK,
			Hash:           inState.Hash,
			PayerName:      inState.PayerName,
		}
		if dberr := db.DbClient.Client.Create(&epBill); dberr.Error != nil {
			glog.Errorln(lib.Log("create database record err", inState.TradeNo, "CheckAccountFromBlockchain"), "", dberr.Error)
			return dberr.Error
		}
	}

	return nil
}

// 交易流水金额为0时，不再上链，直接保存数据库记录(退款未清空)
func SaveTradeInfoToDatabaseWithArrears(inState models.IncomeStatement) error {
	inState.BlockchainStatus = models.Onchain
	inState.Arrears = models.GetScheduleAccount(inState.SubAccountNo)
	if dberr := db.DbClient.Client.Create(&inState); dberr.Error != nil {
		glog.Errorln(lib.Log("create database record err", inState.TradeNo, "WriteBlockchain"), "", dberr.Error)
	}

	epBill := models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("trade_no = ?", inState.TradeNo).First(&epBill); dberr.Error == nil {
		// gorm 查询未判断error
	} else {
		epBill = models.ExpensesBill{
			OrderType:      inState.OrderType,
			AlipaySeller:   inState.AlipaySeller,
			BuyerLogonID:   inState.BuyerLogonID,
			BuyerPayAmount: inState.BuyerPayAmount,
			BuyerUserID:    inState.BuyerUserID,
			BuyerUserType:  inState.BuyerUserType,
			InvoiceAmount:  inState.InvoiceAmount,
			OutTradeNo:     inState.OutTradeNo,
			PointAmount:    inState.PointAmount,
			ReceiptAmount:  inState.ReceiptAmount,
			SendPayDate:    inState.SendPayDate,
			TotalAmount:    inState.TotalAmount,
			Fees:           inState.Fees,
			Arrears:        inState.Arrears,
			TradeNo:        inState.TradeNo,
			SubAccountNo:   inState.SubAccountNo,
			AutoTransfer:   inState.AutoTransfer,
			TradeStatus:    models.TransferOK,
			Hash:           inState.Hash,
			PayerName:      inState.PayerName,
		}
		if dberr := db.DbClient.Client.Create(&epBill); dberr.Error != nil {
			glog.Errorln(lib.Log("create database record err", inState.TradeNo, "CheckAccountFromBlockchain"), "", dberr.Error)
			return dberr.Error
		}
	}

	// 更新排班欠款
	scheAccount := models.ScheduleAccount{ScheduleName: inState.SubAccountNo}
	if dberr := db.DbClient.Client.Model(&scheAccount).Update("arrears", gorm.Expr("arrears - ?", inState.TotalAmount+inState.Fees)); dberr.Error != nil {
		glog.Errorln(lib.Log("update schedule arrears err", inState.TradeNo, "WriteBlockchain"), "", nil)
		return dberr.Error
	}
	models.ScheduleAccountArrears[inState.SubAccountNo] -= (inState.TotalAmount + inState.Fees)

	return nil
}

// 获取之前未上链完成的流水，看当前是否已上链
//
func CheckAccountFromBlockchain() error {
	inStates := []models.IncomeStatement{}
	if dbres := db.DbClient.Client.Where("blockchain_status = ?", models.NotOnchain).Find(&inStates); dbres.Error != nil {
		glog.Errorln(lib.Log("get income statement from database err", "", "CheckAccountFromBlockchain"), "", dbres.Error)
		return nil
	}

	//遍历所有未上链的流水，从区块链查询是否上链完成
	for _, inState := range inStates {
		//如果超过规定时间未返回结果，重新上链
		if time.Since(time.Unix(inState.OnchainTime, 0)).Seconds() >= float64(config.Opts().OnchainTimeLimit) &&
			inState.OnchainNum <= config.Opts().OnchainNumber {
			// 初始化交易秘钥，准备上链
			//
			keyStr := utils.ParseKeyStore()
			if keyStr == "" {
				glog.Errorln(lib.Log("get key store err", inState.TradeNo, "CheckAccountFromBlockchain"), "", nil)
				continue
			}

			// 上链数据结构初始化
			//
			reqg := &protocol.ReqThreeSetBill{
				UserAddress:  config.Opts().FinanceAddress,
				KeyString:    keyStr,
				ScheduleName: inState.SubAccountNo,
				Money:        inState.TotalAmount,
				OrderId:      inState.TradeNo,
			}

			out, err := json.Marshal(reqg)
			if err != nil {
				glog.Errorln(lib.Log("json marshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
				continue
			}

			// 存证上链太慢了，800ms+，简易先回包
			//
			//RespThreeSetBill, err := client.Cli.C.ThreeSetBill(context.Background(), reqg)
			reqBody := bytes.NewReader(out)
			reqUrl := fmt.Sprintf("http://%s:%s/threesetbill", config.Optional.ApiAddress, config.Optional.ApiPort)
			rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
			if err != nil {
				glog.Errorln(lib.Log("write blockchain err", reqg.OrderId, "WriteBlockchain"), "", err, reqg.ScheduleName)
				continue
			}

			respThreeSetBill := protocol.RespThreeSetBill{}
			if err := json.Unmarshal(rspBody, &respThreeSetBill); err != nil {
				glog.Errorln(lib.Log("json unmarshal error", inState.TradeNo, "WriteBlockchain"), "", nil)
				continue
			}

			if respThreeSetBill.Hash == "" || respThreeSetBill.StatusCode != uint32(0) {
				glog.Errorln(lib.Log("respThreeSetBill.GetHash == nil or StatusCode!=uint32(0)", reqg.OrderId, "WriteBlockchain"), respThreeSetBill.StatusCode, respThreeSetBill.Msg, reqg.ScheduleName)
				continue
			}

			if dberr := db.DbClient.Client.Model(&inState).Updates(map[string]interface{}{"on_chain_time": time.Now().Unix(),
				"on_chain_num": gorm.Expr("on_chain_num + ?", 1)}); dberr.Error != nil {
				glog.Errorln(lib.Log("update on_chain_num err", "", "CheckAccountFromBlockchain"), "", dberr.Error)
			}
			continue
		}
		reqg := protocol.ReqGetABById{
			OrderId:      inState.TradeNo,
			ScheduleName: inState.SubAccountNo,
		}

		out, err := json.Marshal(reqg)
		if err != nil {
			glog.Errorln(lib.Log("json marshal error", inState.TradeNo, "CheckAccountFromBlockchain"), "", nil)
			continue
		}

		reqBody := bytes.NewReader(out)
		reqUrl := fmt.Sprintf("http://%s:%s/getabbyid", config.Optional.ApiAddress, config.Optional.ApiPort)
		rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
		if err != nil {
			glog.Errorln(lib.Log("get ab by id err", reqg.OrderId, "GetABById"), "", err, reqg.ScheduleName)
			continue
		}

		respGetABById := protocol.RespGetABById{}
		if err := json.Unmarshal(rspBody, &respGetABById); err != nil {
			glog.Errorln(lib.Log("json unmarshal error", inState.TradeNo, "GetABById"), "", nil)
			continue
		}

		if respGetABById.StatusCode == uint32(protocol.Status_ABByIdFail) {
			glog.Errorln(lib.Log("response of GetABById StatusCode err", inState.TradeNo, "CheckAccountFromBlockchain"), respGetABById.Msg, fmt.Errorf("GetABById status code: %v", respGetABById.StatusCode))
			continue
		}

		if respGetABById.StatusCode != 0 {
			glog.Errorln(lib.Log("response of GetABById StatusCode err", inState.TradeNo, "CheckAccountFromBlockchain"), respGetABById.Msg, fmt.Errorf("GetABById status code: %v", respGetABById.StatusCode))
			continue
		}

		if dberr := db.DbClient.Client.Model(&inState).Update("blockchain_status", models.Onchain); dberr.Error != nil {
			glog.Errorln(lib.Log("udate account status err", inState.TradeNo, "CheckAccountFromBlockchain"), "", dberr.Error)
			continue
		}

		epBill := models.ExpensesBill{}
		if dberr := db.DbClient.Client.Where("trade_no = ?", inState.TradeNo).First(&epBill); dberr.Error == nil {
			// gorm 查询未判断error
		} else {
			epBill = models.ExpensesBill{
				OrderType:      inState.OrderType,
				AlipaySeller:   inState.AlipaySeller,
				BuyerLogonID:   inState.BuyerLogonID,
				BuyerPayAmount: inState.BuyerPayAmount,
				BuyerUserID:    inState.BuyerUserID,
				BuyerUserType:  inState.BuyerUserType,
				InvoiceAmount:  inState.InvoiceAmount,
				OutTradeNo:     inState.OutTradeNo,
				PointAmount:    inState.PointAmount,
				ReceiptAmount:  inState.ReceiptAmount,
				SendPayDate:    inState.SendPayDate,
				TotalAmount:    inState.TotalAmount,
				Fees:           inState.Fees,
				Arrears:        inState.Arrears,
				TradeNo:        inState.TradeNo,
				SubAccountNo:   inState.SubAccountNo,
				AutoTransfer:   inState.AutoTransfer,
				TradeStatus:    models.Onchain,
				Hash:           inState.Hash,
				PayerName:      inState.PayerName,
			}
			if dberr := db.DbClient.Client.Create(&epBill); dberr.Error != nil {
				glog.Errorln(lib.Log("create database record err", inState.TradeNo, "CheckAccountFromBlockchain"), "", dberr.Error)
			}
		}

		var uBill models.UserBill
		for _, ab := range respGetABById.Abs {
			uBill = models.UserBill{
				OrderType: epBill.OrderType,
				Name:      ab.Acco.Name,
				BankCard:  ab.Acco.BankCard,
				WeChat:    ab.Acco.WeChat,
				Alipay:    ab.Acco.Alipay,
				Telephone: ab.Acco.Telephone,
				BillId:    epBill.ID,
				Address:   ab.Address,
				OrderId:   ab.OrderId,
				//TradeNo:         inState.TradeNo,
				Money:           ab.Money,
				Rflag:           ab.Rflag,
				TransferDetails: ab.TransferDetails,
				Radio:           ab.Radio,
				SubWay:          ab.SubWay,
				PaySwitch:       config.Opts().PaySwitch,
				TradeStatus:     models.NotTransferOnChain,
			}
			if dberr := db.DbClient.Client.Create(&uBill); dberr.Error != nil {
				glog.Errorln(lib.Log("create database record err", uBill.Alipay, "CheckAccountFromBlockchain"), inState.TradeNo, dberr.Error)
			}
		}
		glog.Infoln("上链成功*****")
	}

	return nil
}

// 获取未分账的账单
//
func GetNoSubAccount() ([]models.ExpensesBill, error) {
	epBills := []models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("(trade_status = ? or trade_status = ?) and auto_transfer = 1",
		models.Onchain, models.Transfering).Find(&epBills); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "GetNoSubAccount"), "", dberr.Error)
		return []models.ExpensesBill{}, nil
	}

	return epBills, nil
}

// 获取正在分账的账单
//
func GetSubAccountOnchaining() ([]models.ExpensesBill, error) {
	epBills := []models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("(trade_status = ? or trade_status = ?)",
		models.Onchain, models.Transfering).Find(&epBills); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "GetNoSubAccount"), "", dberr.Error)
		return []models.ExpensesBill{}, nil
	}

	return epBills, nil
}

func UserSubAccountToBlockchain(uBill models.UserBill) error {
	blockchFalg := false
	defer func() {
		if !blockchFalg {
			uBill.TradeStatus = models.TransferOnChainFail
			if dberr := db.DbClient.Client.Model(&uBill).Update(uBill); dberr.Error != nil {
				glog.Errorln(lib.Log("udate account status err", uBill.TradeNo, "UserSubAccountToBlockchain"), "", dberr.Error)
			}
		}
	}()
	keyStr := utils.ParseKeyStore()
	if keyStr == "" {
		glog.Errorln(lib.Log("get key store err", "", "UserSubAccountToBlockchain"), "", nil)
		return fmt.Errorf("get blockchain key string error")
	}
	reqg := protocol.ReqConfirm{
		UserAddress:     uBill.Address,
		KeyString:       keyStr,
		OrderId:         uBill.OrderId,
		TransferDetails: uBill.TransferDetails,
	}

	out, err := json.Marshal(reqg)
	if err != nil {
		glog.Errorln(lib.Log("json marshal error", uBill.TradeNo, "UserSubAccountToBlockchain"), "", nil)
		return err
	}

	reqBody := bytes.NewReader(out)
	reqUrl := fmt.Sprintf("http://%s:%s/threeconfirm", config.Optional.ApiAddress, config.Optional.ApiPort)
	rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get three confirm err", reqg.OrderId, "GetABById"), "", err)
		return err
	}

	respConfirm := protocol.RespConfirm{}
	if err := json.Unmarshal(rspBody, &respConfirm); err != nil {
		glog.Errorln(lib.Log("json unmarshal error", uBill.TradeNo, "GetABById"), "", nil)
		return err
	}

	if respConfirm.StatusCode != 0 {
		glog.Errorln(lib.Log("get three confirm err", "", "UserSubAccountToBlockchain"), respConfirm.Msg, fmt.Errorf("respConfirm.StatusCode=%v", respConfirm.StatusCode))
		return fmt.Errorf("get three confirm status code error")
	}

	uBill.TradeStatus = models.TransferOnChaining
	uBill.OnchainTime = time.Now().Unix()
	if dberr := db.DbClient.Client.Model(&uBill).Update(uBill); dberr.Error != nil {
		glog.Errorln(lib.Log("udate account status err", "", "UserSubAccountToBlockchain"), "", dberr.Error)
		return dberr.Error
	}
	if dberr := db.DbClient.Client.Model(&uBill).Updates(map[string]interface{}{"on_chain_num": gorm.Expr("on_chain_num + ?", 1)}); dberr.Error != nil {
		glog.Errorln(lib.Log("update on_chain_num err", "", "UserSubAccountToBlockchain"), "", dberr.Error)
	}
	blockchFalg = true

	return nil
}

func CheckUserSubAccount(uBill models.UserBill) error {
	keyStr := utils.ParseKeyStore()
	if keyStr == "" {
		glog.Errorln(lib.Log("get key store err", "", "CheckUserSubAccount"), "", nil)
		return fmt.Errorf("get blockchain key string error")
	}
	reqg := protocol.ReqGetAccountBook{
		UserAddress: uBill.Address,
		OrderId:     uBill.OrderId,
	}

	out, err := json.Marshal(reqg)
	if err != nil {
		glog.Errorln(lib.Log("json marshal error", uBill.TradeNo, "CheckAccountFromBlockchain"), "", nil)
		return err
	}

	reqBody := bytes.NewReader(out)
	reqUrl := fmt.Sprintf("http://%s:%s/getaccountbook", config.Optional.ApiAddress, config.Optional.ApiPort)
	rspBody, err := utils.SendHttpRequest("POST", reqUrl, reqBody, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get account book err", reqg.OrderId, "GetAccountBook"), "", err)
		return err
	}

	respGetAccountBook := protocol.RespGetAccountBook{}
	if err := json.Unmarshal(rspBody, &respGetAccountBook); err != nil {
		glog.Errorln(lib.Log("json unmarshal error", uBill.TradeNo, "GetAccountBook"), "", nil)
		return err
	}

	if respGetAccountBook.StatusCode != 0 {
		glog.Errorln(lib.Log("get account book ", uBill.OrderId, "GetAccountBook"), respGetAccountBook.Msg, fmt.Errorf("respGetAccountBook.StatusCode=%v", respGetAccountBook.StatusCode))
		return fmt.Errorf("get account book error")
	}

	uBill.TradeStatus = models.TransferOnChainOK
	uBill.Rflag = true
	if dberr := db.DbClient.Client.Model(&uBill).Update(uBill); dberr.Error != nil {
		glog.Errorln(lib.Log("udate user account status err", uBill.OrderId, "CheckUserSubAccount"), uBill.TradeNo, dberr.Error)
		return dberr.Error
	}

	return nil
}
