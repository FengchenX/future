package manager

import (
	"github.com/golang/glog"
	"sub_account_service/order_server_zhengdao/config"
	"sub_account_service/order_server_zhengdao/lib"
	"time"
)

func InitTransfer() {

	go func() {
		for {
			glog.Infoln("1.定时获取未转账成功的订单，向企业用户转账")
			if time.Now().Hour() == config.Opts().ReTransferTime {
				if err := HandleTransferToCompany(); err != nil {
					glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				}
			}
			time.Sleep(config.Opts().TimeTicker1 * time.Second)
		}

	}()

	go func() {
		for {
			glog.Infoln("2.从数据库获取 添加到编号系统失败 的交易流水，重新添加")
			if time.Now().Hour() == config.Opts().ReTransferTime {
				if err := HandleTransferToOrderServer(); err != nil {
					glog.Errorln(lib.Log("timing get tarding stream err", "", "GetTradingStreamFromThirdPartPay"), "", err)
				}
			}
			time.Sleep(config.Opts().TimeTicker2 * time.Second)
		}
	}()
}
