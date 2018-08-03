package third_part_pay

import (
	//"runtime/pprof"
	"time"

	"github.com/golang/glog"

	"sub_account_service/finance/config"
	"sub_account_service/finance/db"
	"sub_account_service/finance/lib"
	"sub_account_service/finance/models"
)

var UserAccountCache = make(map[string]float64, 100)

//对接三方支付的定时任务
func InitThirdPartPay() {
	HandleTimeTicker1 := time.NewTicker(config.Opts().TimeTicker1 * time.Second)
	HandleTimeTicker2 := time.NewTicker(config.Opts().TimeTicker2 * time.Second)
	HandleTimeTicker3 := time.NewTicker(config.Opts().TimeTicker3 * time.Second)
	HandleTimeTicker5 := time.NewTicker(config.Opts().TimeTicker5 * time.Second)
	HandleTimeTicker6 := time.NewTicker(config.Opts().TimeTicker6 * time.Second)
	HandleTimeTicker7 := time.NewTicker(config.Opts().TimeTicker7 * time.Second)
	HandleTimeTicker8 := time.NewTicker(config.Opts().TimeTicker8 * time.Second)

	ts := &TardingStreamInfo{}
	sa := &SubAccountInfo{}
	go func() {
		glog.Infoln("1.定时从三方支付获取流水信息，存储redis")
		if err := ts.GetTradingStreamFromThirdPartPay(); err != nil {
			glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker1.C:
				glog.Infoln("1.定时从三方支付获取流水信息，存储redis")
				if err := ts.GetTradingStreamFromThirdPartPay(); err != nil {
					glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				}
				//pprof.StopCPUProfile()
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("2.定时查询获取详情失败的账单，重新获取账单详情")
		if err := ts.GetTradingStreamDetailsFromThirdPartPay(); err != nil {
			glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker2.C:
				glog.Infoln("2.定时查询获取详情失败的账单，重新获取账单详情")
				if err := ts.GetTradingStreamDetailsFromThirdPartPay(); err != nil {
					glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("3.定时获取未做上链操作的流水，进行上链操作")
		if err := ts.TradingStreamWriteToBlockchain(); err != nil {
			glog.Errorln(lib.Log("timing write tarding stream to blockchain err", "", "TradingStreamWriteToBlockchain"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker3.C:
				glog.Infoln("3.定时获取未做上链操作的流水，进行上链操作")
				if err := ts.TradingStreamWriteToBlockchain(); err != nil {
					glog.Errorln(lib.Log("timing write tarding stream to blockchain err", "", "TradingStreamWriteToBlockchain"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("5.定时从数据库获取未分账信息，开始分账")
		if err := sa.HandleSubAccoutFromBlockchain(); err != nil {
			glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutFromBlockchain"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker5.C:
				glog.Infoln("5.定时从数据库获取未分账信息，开始分账")
				if err := sa.HandleSubAccoutFromBlockchain(); err != nil {
					glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutFromBlockchain"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("6.定时查询正在上链的分账信息，是否上链成功")
		if err := sa.HandleSubAccoutOnBlockchain(); err != nil {
			glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutOnBlockchain"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker6.C:
				glog.Infoln("6.定时查询正在上链的分账信息，是否上链成功")
				if err := sa.HandleSubAccoutOnBlockchain(); err != nil {
					glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutOnBlockchain"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("7.定时查询上链失败或者未上链的分账信息，重新上链")
		if err := sa.HandleSubAccoutReOnBlockchain(); err != nil {
			glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutReOnBlockchain"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker7.C:
				glog.Infoln("7.定时查询上链失败或者未上链的分账信息，重新上链")
				if err := sa.HandleSubAccoutReOnBlockchain(); err != nil {
					glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleSubAccoutReOnBlockchain"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		glog.Infoln("8.定点获取用户累积未转账的money，准备转账")
		if time.Now().Hour() == config.Opts().TransferUserLimitMoneyTime {
			if err := sa.HandleUserAccountLessThanLimit(); err != nil {
				glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleUserAccountLessThanLimit"), "", err)
			}
		}
		for {
			select {
			case <-HandleTimeTicker8.C:
				glog.Infoln("8.定点获取用户累积未转账的money，准备转账")
				if time.Now().Hour() == config.Opts().TransferUserLimitMoneyTime {
					if err := sa.HandleUserAccountLessThanLimit(); err != nil {
						glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleUserAccountLessThanLimit"), "", err)
					}
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func InitScheduleArrears() {
	scheAccounts := []models.ScheduleAccount{}
	if dberr := db.DbClient.Client.Find(&scheAccounts); dberr.Error != nil {
		glog.Errorln(lib.Log("handle account from blockchain err", "", "HandleUserAccountLessThanLimit"), "", dberr.Error)
	}
	for _, v := range scheAccounts {
		if v.Arrears != 0 {
			models.ScheduleAccountArrears[v.ScheduleName] = v.Arrears
		}
	}
}
