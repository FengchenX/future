package blockchain

import (
	"time"

	"github.com/golang/glog"

	"sub_account_service/finance/config"
	"sub_account_service/finance/lib"
)

// 上链检查定时器的初始化
func InitBlockchain() {
	HandleTimeTicker4 := time.NewTicker(config.Opts().TimeTicker4 * time.Second)

	go func() {
		glog.Infoln("4.定时获取之前未上链完成的流水，看当前是否已上链")
		if err := CheckAccountFromBlockchain(); err != nil {
			glog.Errorln(lib.Log("check account from block err", "", "CheckAccountFromBlockchain"), "", err)
		}
		for {
			select {
			case <-HandleTimeTicker4.C:
				glog.Infoln("4.定时获取之前未上链完成的流水，看当前是否已上链")
				if err := CheckAccountFromBlockchain(); err != nil {
					glog.Errorln(lib.Log("check account from block err", "", "CheckAccountFromBlockchain"), "", err)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

}
