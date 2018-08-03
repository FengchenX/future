package third_part_pay

import (
	//	"time"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	"sub_account_service/finance/blockchain"
	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	"sub_account_service/finance/models"
	"sub_account_service/finance/payment"
	"sub_account_service/finance/payment/alipayModels"
	svcModels "sub_account_service/number_server/models"
)

type SubAccountInfo struct {
}

// 从数据库获取分账信息，然后通过第三方支付转账
func (this *SubAccountInfo) HandleSubAccoutFromBlockchain() error {
	var err error
	epBills, err := blockchain.GetNoSubAccount()
	if err != nil {
		glog.Errorln(lib.Log("get account from blockchain err", "", "HandleSubAccoutFromBlockchain"), "", err)
		return err
	}

	if len(epBills) != 0 {
		glog.Infoln("HandleSubAccoutFromBlockchain*********************未分账完成的流水数量=", len(epBills))
	}

	for _, epBill := range epBills {
		if dberr := db.DbClient.Client.Model(&epBill).Update("trade_status", models.Transfering); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), "", dberr.Error)
			continue
		}
		glog.Infoln("HandleSubAccoutFromBlockchain*********************未分账完成的流水epbills=", epBill.ID)
		uBills := []models.UserBill{}
		if dberr := db.DbClient.Client.Where("bill_id = ? and rflag = false and (trade_status = ? or trade_status = ?)",
			epBill.ID, models.NotTransferOnChain, models.PayFailed).Find(&uBills); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", "", "HandleSubAccoutFromBlockchain"), "", dberr.Error)
			continue
		}

		for _, uBill := range uBills {
			glog.Infoln("HandleSubAccoutFromBlockchain*********************未付款的分账信息的user bills=", uBill.ID)

			if config.Opts().PaySwitch > 0 { //转账开关打开
				uBillOp := models.UserBill{}
				tx := db.DbClient.Client.Begin()
				glog.Infoln(fmt.Sprintf("id = %v and bill_id = %v and rflag = false and (trade_status = %v or trade_status = %v)",
					uBill.ID, epBill.ID, models.NotTransferOnChain, models.PayFailed))
				if dberr := tx.Where("id = ? and bill_id = ? and rflag = false and (trade_status = ? or trade_status = ?)",
					uBill.ID, epBill.ID, models.NotTransferOnChain, models.PayFailed).First(&uBillOp); dberr.Error != nil {
					glog.Errorln(lib.Log("get account from database err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
					tx.Rollback()
					continue
				}

				if dberr := tx.Model(&uBillOp).Update("trade_status", models.Paying); dberr.Error != nil {
					glog.Errorln(lib.Log("update trade status to database err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
					tx.Rollback()
					continue
				}
				tx.Commit()

				if arrearsFlag, err := this.CheckUserAccountArrears(&uBill); err != nil {
					glog.Errorln(lib.Log("check subaccount arrears err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), "", err)
					continue
				} else if arrearsFlag {
					continue
				}

				if uBillOp.OrderType == 0 { //支付宝支付
					uBillOp, err = this.AlipayForUserAccount(uBillOp, epBill.PayerName)
					if err != nil {
						glog.Errorln(lib.Log("alipay for user account err", uBillOp.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, err)
						continue
					}

				} else if uBillOp.OrderType == 1 { //微信支付

				} else {

				}

				if dberr := db.DbClient.Client.Model(&uBillOp).Update("trade_status", models.PayOK); dberr.Error != nil {
					glog.Errorln(lib.Log("update trade status to database err", uBillOp.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
					continue
				}
				if err := blockchain.UserSubAccountToBlockchain(uBillOp); err != nil {
					glog.Errorln(lib.Log("get account from blockchain err", uBillOp.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, err)
					continue
				}
			} else { //转账开关关闭
				if err := blockchain.UserSubAccountToBlockchain(uBill); err != nil {
					glog.Errorln(lib.Log("get account from blockchain err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, err)
					continue
				}
			}
		}
	}

	return nil
}

func (this *SubAccountInfo) AlipayForUserAccount(uBill models.UserBill, payerName string) (models.UserBill, error) {
	payflag := true
	defer func() {
		if !payflag {
			if dberr := db.DbClient.Client.Model(&uBill).Update("trade_status", models.PayFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), uBill.OrderId, dberr.Error)
			}
		}
	}()
	if uBill.Money-uBill.Arrears < config.Opts().TransferUserLimitMoney {
		if err := this.AlipayForUserAccountLessThanLimit(uBill); err != nil {
			payflag = false
			return uBill, err
		}
	}
	uid, err := uuid.NewV4()
	if err != nil {
		payflag = false
		return uBill, err
	}
	taReq := alipayModels.TransferToAccountReq{
		OrderId:       fmt.Sprintf("%s", uid), //需要唯一id，例如uuid
		Amount:        uBill.Money - uBill.Arrears,
		PayeeType:     "ALIPAY_LOGONID",
		PayeeAccount:  uBill.Alipay,
		PayerShowName: payerName,
		PayeeRealName: uBill.Name,
		Remark:        "",
	}
	tAccount, err := payment.TransferToAccount(taReq)
	if err != nil {
		glog.Errorln(lib.Log("transfer to account err", "", "HandleSubAccoutFromBlockchain"), "", err)
		payflag = false
		return uBill, err
	}

	glog.Infoln("付款成功**********************************分账信息的user bills=", uBill.ID)

	//	loc, _ := time.LoadLocation("Local")                                             //重要：获取时区
	//	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tAccount.PayDate, loc) //使用模板在对应时区转化为time.time类型
	//	payDate := theTime.Unix()

	if dberr := db.DbClient.Client.Model(&models.UserAccount{Address: uBill.Address}).Update("arrears", 0); dberr.Error != nil {
		glog.Errorln(lib.Log("update trade status to database err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), uBill.OrderId, dberr.Error)
		return uBill, dberr.Error
	}

	if dberr := db.DbClient.Client.Model(&uBill).Updates(map[string]interface{}{"trade_status": models.PayOK,
		"trade_no": tAccount.TradeNo, "pay_date": tAccount.PayDate}); dberr.Error != nil {
		glog.Errorln(lib.Log("update trade status to database err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), uBill.OrderId, dberr.Error)
		return uBill, dberr.Error
	}

	out, err := json.Marshal(tAccount)
	if err != nil {
		glog.Errorln(lib.Log("marshal tAccount err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), uBill.OrderId, err)
		return uBill, err
	}
	uBill.TransferDetails = string(out)

	return uBill, nil
}

func (this *SubAccountInfo) AlipayForUserAccountLessThanLimit(uBill models.UserBill) error {
	tx := db.DbClient.Client.Begin()
	uAccount := models.UserAccount{Address: uBill.Address}
	if dberr := tx.Where("address = ?", uBill.Address).First(&uAccount); dberr.Error != nil {
		uAccount.TradeStatus = models.InitState
		if dberr := tx.Create(&uAccount); dberr.Error != nil {
			glog.Errorln(lib.Log("create useraccount err", "", "AlipayForUserAccountLessThanLimit"), "", dberr.Error)
			tx.Rollback()
			return dberr.Error
		}
	}
	tx.Commit()
	if dberr := tx.Model(&uAccount).Update("no_pay_sum", gorm.Expr("no_pay_sum + ?", uBill.Money-uBill.Arrears)); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "AlipayForUserAccountLessThanLimit"), "", dberr.Error)
		tx.Rollback()
		return dberr.Error
	}

	return nil
}

func (this *SubAccountInfo) CheckUserAccountArrears(uBill *models.UserBill) (bool, error) {
	uAccount := models.UserAccount{}
	if dberr := db.DbClient.Client.Where("address = ?", uBill.Address).First(&uAccount); dberr.Error != nil {
		return false, nil
	}

	if dberr := db.DbClient.Client.Model(&uBill).Update("arrears", uAccount.Arrears); dberr.Error != nil {
		glog.Errorln(lib.Log("update user account err", "", "CheckUserAccountArrears"), "", dberr.Error)
		db.DbClient.Client.Rollback()
		return true, dberr.Error
	}
	if uAccount.Arrears >= uBill.Money {
		tx := db.DbClient.Client.Begin()
		if dberr := tx.Model(&uAccount).Update("arrears", gorm.Expr("arrears - ?", uBill.Money)); dberr.Error != nil {
			glog.Errorln(lib.Log("update user account err", "", "CheckUserAccountArrears"), "", dberr.Error)
			tx.Rollback()
			return true, dberr.Error
		}
		if dberr := tx.Model(&uAccount).Update("trade_status", models.PayOK); dberr.Error != nil {
			glog.Errorln(lib.Log("update user account err", "", "CheckUserAccountArrears"), "", dberr.Error)
			tx.Rollback()
			return true, dberr.Error
		}
		tx.Commit()
		if err := blockchain.UserSubAccountToBlockchain(*uBill); err != nil {
			glog.Errorln(lib.Log("get account from blockchain err", uBill.Telephone, "HandleSubAccoutFromBlockchain"), uBill.TradeNo, err)
			return true, err
		}
		return true, nil
	}

	return false, nil
}

// 从区块链获取分账信息上链中的状态
func (this *SubAccountInfo) HandleSubAccoutOnBlockchain() error {

	epBills, err := blockchain.GetSubAccountOnchaining()
	if err != nil {
		glog.Errorln(lib.Log("get account from blockchain err", "", "HandleSubAccoutOnBlockchain"), "", err)
		return err
	}

	for _, epBill := range epBills {

		uBills := []models.UserBill{}
		if dberr := db.DbClient.Client.Where("bill_id = ? and rflag = false and trade_status = ?",
			epBill.ID, models.TransferOnChaining).Find(&uBills); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", epBill.TradeNo, "HandleSubAccoutOnBlockchain"), "", dberr.Error)
			continue
		}

		// 如果没有正在上链的流水，轮询下一个
		if len(uBills) == 0 {
			continue
		}

		for _, uBill := range uBills {
			//如果超过规定时间未返回结果，重新上链
			if time.Since(time.Unix(uBill.OnchainTime, 0)).Seconds() >= float64(config.Opts().OnchainTimeLimit) &&
				uBill.OnchainNum <= config.Opts().OnchainNumber {
				if err := blockchain.UserSubAccountToBlockchain(uBill); err != nil {
					glog.Errorln(lib.Log("get account from blockchain err", uBill.Telephone, "HandleSubAccoutReOnBlockchain"), epBill.TradeNo, err)
					continue
				}
			}
			//检查是否上链完成
			glog.Infoln("HandleSubAccoutOnBlockchain===================正在上链的分账信息uBill.ID=", uBill.ID)
			if err := blockchain.CheckUserSubAccount(uBill); err != nil {
				glog.Errorln(lib.Log("get account from blockchain err", uBill.Telephone, "HandleSubAccoutOnBlockchain"), epBill.TradeNo, err)
				continue
			}
			glog.Infoln("HandleSubAccoutOnBlockchain===================上链成功的分账信息uBill.ID=", uBill.ID)
		}
		uBills1 := []models.UserBill{}
		uBills2 := []models.UserBill{}
		if dberr := db.DbClient.Client.Where("bill_id = ? and rflag = true and trade_status = ?",
			epBill.ID, models.TransferOnChainOK).Find(&uBills1); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", epBill.TradeNo, "HandleSubAccoutOnBlockchain"), "", dberr.Error)
			continue
		}
		if dberr := db.DbClient.Client.Where("bill_id = ?", epBill.ID).Find(&uBills2); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", epBill.TradeNo, "HandleSubAccoutOnBlockchain"), "", dberr.Error)
			continue
		}

		if len(uBills1) == len(uBills2) {
			//更新流水状态
			if dberr := db.DbClient.Client.Model(&epBill).Update("trade_status", models.TransferOK); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", epBill.TradeNo, "HandleSubAccoutOnBlockchain"), "", dberr.Error)
				continue
			}
			glog.Infoln("HandleSubAccoutOnBlockchain===================转账成功的流水信息uBill.ID=", epBill.ID)
		}
	}

	return nil
}

// 从数据库获取上链失败的分账信息，重新上链
func (this *SubAccountInfo) HandleSubAccoutReOnBlockchain() error {

	epBills, err := blockchain.GetSubAccountOnchaining()
	if err != nil {
		glog.Errorln(lib.Log("get account from blockchain err", "", "HandleSubAccoutReOnBlockchain"), "", err)
		return err
	}

	for _, epBill := range epBills {

		uBills := []models.UserBill{}
		if dberr := db.DbClient.Client.Where("bill_id = ? and rflag = false and trade_status = ?",
			epBill.ID, models.TransferOnChainFail).Find(&uBills); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", epBill.TradeNo, "HandleSubAccoutReOnBlockchain"), "", dberr.Error)
			continue
		}

		for _, uBill := range uBills {
			glog.Infoln("HandleSubAccoutReOnBlockchain===================上链失败的分账信息uBill.ID=", uBill.ID)
			//重新上链
			if err := blockchain.UserSubAccountToBlockchain(uBill); err != nil {
				glog.Errorln(lib.Log("get account from blockchain err", uBill.Telephone, "HandleSubAccoutReOnBlockchain"), epBill.TradeNo, err)
				continue
			}
		}
	}

	return nil
}

// 从数据库获取用户累积未转账的money，准备转账
func (this *SubAccountInfo) HandleUserAccountLessThanLimit() error {
	uAccounts := []models.UserAccount{}
	if dberr := db.DbClient.Client.Where("(trade_status=? or trade_status=?)", models.InitState, models.PayFailed).Find(&uAccounts); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "HandleSubAccoutReOnBlockchain"), "", dberr.Error)
		return dberr.Error
	}

	for _, uAccount := range uAccounts {
		tx := db.DbClient.Client.Begin()
		if dberr := tx.Where("address=? and (trade_status=? or trade_status=?)", uAccount.Address, models.InitState, models.PayFailed).First(&uAccount); dberr.Error != nil {
			tx.Rollback()
			continue
		}

		if uAccount.NoPaySum < config.Opts().TransferUserLimitMoney {
			tx.Commit()
			continue
		}

		if dberr := tx.Model(&uAccount).Update("trade_status", models.Paying); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", uAccount.Telephone, "HandleSubAccoutReOnBlockchain"), "", dberr.Error)
			tx.Rollback()
			continue
		}
		tx.Commit()

		uid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		taReq := alipayModels.TransferToAccountReq{
			OrderId:       fmt.Sprintf("%s", uid), //需要唯一id，例如uuid
			Amount:        uAccount.NoPaySum,
			PayeeType:     "ALIPAY_LOGONID",
			PayeeAccount:  uAccount.Alipay,
			PayerShowName: "",
			PayeeRealName: uAccount.Name,
			Remark:        "",
		}
		_, err = payment.TransferToAccount(taReq)
		if err != nil {
			glog.Errorln(lib.Log("transfer to account err", "", "HandleSubAccoutFromBlockchain"), "", err)
			if dberr := db.DbClient.Client.Model(&uAccount).Update("trade_status", models.PayFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", uAccount.Telephone, "HandleSubAccoutFromBlockchain"), "", dberr.Error)
				continue
			}
			continue
		}

		glog.Infoln("累积金额付款成功**********************************user address=", uAccount.Address)

		if dberr := tx.Model(&uAccount).Update(map[string]interface{}{"trade_status": models.InitState, "no_pay_sum": gorm.Expr("no_pay_sum - ?", uAccount.NoPaySum)}); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", uAccount.Telephone, "HandleSubAccoutReOnBlockchain"), "", dberr.Error)
			continue
		}
	}

	return nil
}

// 主动转账函数，根据分账用户转账
func (this *SubAccountInfo) InitiativeSubAccount(uBillIds []uint) (map[uint]int, error) {
	res := map[uint]int{}
	for _, uBillId := range uBillIds {
		// 初始为未成功
		res[uBillId] = 0
		uBillOp := models.UserBill{}
		tx := db.DbClient.Client.Begin()
		if dberr := tx.Where("id = ? and rflag = false and (trade_status = ? or trade_status = ?)",
			uBillId, models.NotTransferOnChain, models.PayFailed).First(&uBillOp); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", dberr.Error)
			tx.Rollback()
			continue
		}
		if dberr := tx.Model(&uBillOp).Update("trade_status", models.Paying); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", dberr.Error)
			tx.Rollback()
			continue
		}
		tx.Commit()

		uid, err := uuid.NewV4()
		if err != nil {
			return map[uint]int{}, err
		}
		taReq := alipayModels.TransferToAccountReq{
			OrderId:      fmt.Sprintf("%s", uid), //需要唯一id，例如uuid
			Amount:       uBillOp.Money,
			PayeeType:    "ALIPAY_LOGONID",
			PayeeAccount: uBillOp.Alipay,
			//PayerShowName: epBill.PayerName,
			PayeeRealName: uBillOp.Name,
			Remark:        "",
		}
		glog.Infoln("HandleSubAccoutFromBlockchain===================未付款的分账信息的user bills=", uBillOp)
		if config.Opts().PaySwitch > 0 {
			if uBillOp.OrderType == 0 { //支付宝支付
				tAccount, err := payment.TransferToAccount(taReq)
				if err != nil {
					glog.Errorln(lib.Log("transfer to account err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", err)
					if dberr := db.DbClient.Client.Model(&uBillOp).Update("trade_status", models.PayFailed); dberr.Error != nil {
						glog.Errorln(lib.Log("update trade status to database err", "", "HandleSubAccoutFromBlockchain"), "", dberr.Error)
						continue
					}
					continue
				}
				// 转账成功
				res[uBillId] = 1

				if dberr := db.DbClient.Client.Model(&uBillOp).Update(map[string]interface{}{"trade_status": models.PayOK, "trade_no": tAccount.TradeNo}); dberr.Error != nil {
					glog.Errorln(lib.Log("update trade status to database err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", dberr.Error)
					continue
				}

				out, err := json.Marshal(tAccount)
				if err != nil {
					glog.Errorln(lib.Log("marshal tAccount err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", err)
					continue
				}
				uBillOp.TransferDetails = string(out)
			} else if uBillOp.OrderType == 1 { //微信支付

			} else {

			}
		} else {
			return map[uint]int{}, fmt.Errorf("pay switch is close")
		}

		// 转账成功
		res[uBillId] = 1

		if dberr := db.DbClient.Client.Model(&uBillOp).Update(map[string]interface{}{"trade_status": models.PayOK, "pay_switch": config.Opts().PaySwitch}); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", dberr.Error)
			continue
		}

		if err := blockchain.UserSubAccountToBlockchain(uBillOp); err != nil {
			glog.Errorln(lib.Log("get account from blockchain err", fmt.Sprintf("%v", uBillId), "HandleSubAccoutFromBlockchain"), "", err)
			continue
		}

	}

	return res, nil
}

// 给顾客退款
func CustomerRefundTrade(order svcModels.Orders) error {
	flag := false
	defer func() {
		if !flag {
			if err := AddOrderToFailedQueue(order); err != nil {
				glog.Errorln(lib.Log("add order to failed to queue err, trade num: ", order.ThirdTradeNo, "AddOrderToFailedQueue"), "", err)
			}
		}
	}()
	epBill := models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("trade_no=? and trade_status=?", order.ThirdTradeNo, models.TransferOK).First(&epBill); dberr.Error != nil {
		if err := CustomerRefundWithNewTradeNo(order); err != nil {
			glog.Errorln(lib.Log("customer refund with new trade number err", fmt.Sprintf("%v", order.ThirdTradeNo), "CustomerRefundTrade"), "", err)
			return err
		}
	} else {
		if err := CustomerRefundWithTradeNo(order); err != nil {
			glog.Errorln(lib.Log("customer refund with new trade number err", fmt.Sprintf("%v", order.ThirdTradeNo), "CustomerRefundTrade"), "", err)
			return err
		}
	}
	flag = true
	return nil
}

// 给顾客退款 (交易号可查询)
func CustomerRefundWithTradeNo(order svcModels.Orders) error {
	epBill := models.ExpensesBill{}
	if dberr := db.DbClient.Client.Where("trade_no=? and trade_status=?", order.ThirdTradeNo, models.TransferOK).First(&epBill); dberr.Error != nil {
		return fmt.Errorf("Not found trade record")
	}

	tx := db.DbClient.Client.Begin()

	uBills := []models.UserBill{}
	fmt.Println("bill_id=", epBill.ID)
	if dberr := tx.Where("bill_id = ?", epBill.ID).Find(&uBills); dberr.Error != nil {
		glog.Errorln(lib.Log("get from database err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
		tx.Rollback()
		return dberr.Error
	}
	for _, uBill := range uBills {
		uAccount := models.UserAccount{}
		if dberr := tx.Where("address = ?", uBill.Address).First(&uAccount); dberr.Error != nil {
			uAccount.Address = uBill.Address
			uAccount.TradeStatus = models.InitState
			uAccount.Arrears = uBill.Money
			if dberr := tx.Create(&uAccount); dberr.Error != nil {
				tx.Rollback()
				return dberr.Error
			}
			UserAccountCache[uAccount.Address] = uBill.Money
		} else {
			if dberr := tx.Model(&models.UserAccount{Address: uBill.Address}).Update("arrears", gorm.Expr("arrears + ?", uBill.Money)); dberr.Error != nil {
				tx.Rollback()
				return dberr.Error
			}
			UserAccountCache[uAccount.Address] += uBill.Money
		}

		if dberr := tx.Model(&uBill).Update("trade_status", models.RefundOK); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade refund err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
			tx.Rollback()
			return dberr.Error
		}
	}

	refTrade := models.RefundTrade{
		OrderType:      epBill.OrderType,
		AlipaySeller:   epBill.AlipaySeller,
		BuyerLogonID:   epBill.BuyerLogonID,
		BuyerPayAmount: epBill.BuyerPayAmount,
		BuyerUserID:    epBill.BuyerUserID,
		BuyerUserType:  epBill.BuyerUserType,
		InvoiceAmount:  epBill.InvoiceAmount,
		OutTradeNo:     epBill.OutTradeNo,
		PointAmount:    epBill.PointAmount,
		ReceiptAmount:  epBill.ReceiptAmount,
		SendPayDate:    epBill.SendPayDate,
		TotalAmount:    epBill.TotalAmount,
		Fees:           epBill.Fees,
		TradeNo:        epBill.TradeNo,
		SubAccountNo:   epBill.SubAccountNo,
		PayerName:      epBill.PayerName,
		AutoTransfer:   epBill.AutoTransfer,
	}

	//生成退款记录
	if dberr := db.DbClient.Client.Create(&refTrade); dberr.Error != nil {
		glog.Errorln(lib.Log("create refund trade record err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
		tx.Rollback()
		return dberr.Error
	}

	// 交易流水标记为已退款
	if dberr := db.DbClient.Client.Model(&epBill).Update("trade_status", models.RefundOK); dberr.Error != nil {
		glog.Errorln(lib.Log("update trade refund err", epBill.TradeNo, "HandleSubAccoutFromBlockchain"), epBill.TradeNo, dberr.Error)
		tx.Rollback()
		return dberr.Error
	}
	tx.Commit()

	return nil
}

// 给顾客退款 (新的交易号)
func CustomerRefundWithNewTradeNo(order svcModels.Orders) error {
	tx := db.DbClient.Client.Begin()
	scheAccount := models.ScheduleAccount{}
	if dberr := tx.Where("schedule_name = ?", order.SubAccountNo).First(&scheAccount); dberr.Error != nil {
		scheAccount.ScheduleName = order.SubAccountNo
		scheAccount.Arrears = order.TransferAmount
		if dberr := tx.Create(&scheAccount); dberr.Error != nil {
			tx.Rollback()
			return dberr.Error
		}
		models.ScheduleAccountArrears[order.SubAccountNo] = order.TransferAmount
	} else {
		if dberr := tx.Model(&scheAccount).Update("arrears", gorm.Expr("arrears + ?", order.TransferAmount)); dberr.Error != nil {
			tx.Rollback()
			return dberr.Error
		}
		if dberr := tx.Model(&models.ScheduleAccount{ScheduleName: order.SubAccountNo}).First(&scheAccount); dberr.Error != nil {
			tx.Rollback()
			return dberr.Error
		}
		models.ScheduleAccountArrears[order.SubAccountNo] = order.TransferAmount + scheAccount.Arrears
	}

	refTrade := models.RefundTrade{
		OrderType:    int64(order.OrderType),
		TotalAmount:  order.TransferAmount,
		TradeNo:      order.ThirdTradeNo,
		SubAccountNo: order.SubAccountNo,
		PayerName:    order.Company,
	}

	//生成退款记录
	if dberr := db.DbClient.Client.Create(&refTrade); dberr.Error != nil {
		glog.Errorln(lib.Log("create refund trade record err", order.ThirdTradeNo, "HandleSubAccoutFromBlockchain"), order.ThirdTradeNo, dberr.Error)
		tx.Rollback()
		return dberr.Error
	}
	tx.Commit()

	return nil
}
