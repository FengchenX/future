package api

import (
	"sub_account_service/app_server_v2/db"
	"sub_account_service/app_server_v2/model"
	"time"

	"sub_account_service/app_server_v2/config"

	"github.com/golang/glog"
	"sub_account_service/app_server_v2/lib"
)

//InitReset 初始化reset
func InitReset() {
	go ResetQuo()
}

//Adds 添加reset publisher: 发布者, subCode: 分账编号, rs: 分配明细
func Adds(publisher string, subCode string, rs []db.Reset) {
	glog.Infoln("Adds********************publisher, subCode, rs:", publisher, rs, subCode)
	if len(rs) == 0 {
		return
	}
	DB := db.DbClient.Client
	if DB == nil {
		return
	}
	var temps []db.Reset
	if mydb := DB.Where("publisher = ? AND sub_code = ?", publisher, subCode).Find(&temps); mydb.Error != nil {
		glog.Errorln("Adds***************db err:", mydb.Error)
		return
	}
	for _, temp := range temps {
		if mydb := DB.Delete(&temp); mydb.Error != nil {
			glog.Errorln("Adds***************db err:", mydb.Error)
			continue
		}
	}

	for _, r := range rs {
		Add(publisher, r)
	}
}

//Updates 更新reset的用户地址发布排班时调用
func Updates(publisher string, subCode string, paiBans []model.PaiBan) {
	DB := db.DbClient.Client
	if DB == nil {
		return
	}
	if len(paiBans) == 0 {
		glog.Errorln("Updates******************paiBan len is 0")
		return
	}
	var temps []db.Reset
	if mydb := DB.Where("publisher = ? AND user_addr = ? AND sub_code = ?", publisher, "", subCode).Find(&temps); mydb.Error != nil {
		glog.Errorln("Updates*******************db err:", mydb.Error)
		return
	}
	tx := DB.Begin()
	for _, temp := range temps {
		for _, paiBan := range paiBans {
			if temp.Job == paiBan.JobName {
				if mydb := tx.Model(&temp).Update("user_addr", paiBan.UserAddress); mydb.Error != nil {
					glog.Errorln("Updates***********************db err:", mydb.Error)
					continue
				}
			}
		}
	}
	tx.Commit()
}

//Add 添加一个reset
func Add(publisher string, r db.Reset) {
	glog.Infoln("Add**********************Reset: ", r)
	DB := db.DbClient.Client
	if DB.Error != nil {
		return
	}
	isNew := DB.NewRecord(r)
	if isNew {
		DB.Create(&r)
		if isNew := DB.NewRecord(r); !isNew {
			glog.Infoln("Add***********************create sucesss")
		}
	} else {
		glog.Errorln("Add*************************r is not a new record")
	}
}

//ResetQuo 重置定额
func ResetQuo() {
	DB := db.DbClient.Client
	if DB == nil {
		return
	}
	t := time.NewTicker(time.Duration(config.ConfInst().ResetQuoTime) * time.Second)
	for {
		select {
		case <-t.C:
			glog.Infoln("ResetQuo********************tick")
			var temps []db.Reset
			if mydb := DB.Where("user_addr != ?", "").Find(&temps); mydb.Error != nil {
				glog.Errorln("ResetQuo*********************db err:", mydb.Error)
				continue
			}
			for _, temp := range temps {
				handleOne(temp)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
func handleOne(r db.Reset) {
	DB := db.DbClient.Client
	if DB == nil {
		return
	}
	if r.UserAddr == "" {
		glog.Infoln("handOne*****************userAddr is nil, subCode:", r.SubCode)
		return
	}
	if r.ResetWay == 0 {
		return
	}
	keyString := lib.KeyMap.Calc(config.ConfInst().KeyStore, config.ConfInst().Phrase)
	if r.ResetWay == 1 {
		//每天重置
		now := time.Now()
		if dua := now.Sub(r.LastTime); dua > 1*time.Hour {
			if now.Hour() == int(r.ResetTime) {
				//发起重置
				var req model.ReqResetQuo
				var resp model.RespResetQuo
				req.SubCode = r.SubCode
				req.UserAddr = r.UserAddr
				req.KeyStore = keyString
				glog.Infoln("handleOne****************** start")
				//r.LastTime = now
				//todo 测试
				// if mydb := DB.Model(&r).Update("last_time", now); mydb.Error != nil {
				// 	glog.Errorln("handOne************db err:", mydb.Error)
				// 	return
				// }
				if req.UserAddr == "" || req.SubCode == "" || req.KeyStore == "" {
					glog.Errorln("handleOne**********************req.UserAddr == '' || req.SubCode == '' || req.KeyStore == ''")
					return
				}
				glog.Infoln("handleOne***************req:", req)
				if err := doPost("/resetquo", req, &resp); err != nil {
					glog.Errorln("handleOne****************doPost:", err)
					return
				}
				glog.Infoln("handleOne******************resp: ", resp)
				if mydb := DB.Model(&r).Update("last_time", now); mydb.Error != nil {
					glog.Errorln("handleOne************db err:", mydb.Error)
					return
				}
			}
		}
	} else if r.ResetWay == 2 {
		//每月重置
		now := time.Now()
		if dua := now.Sub(r.LastTime); dua > 24*time.Hour {
			if now.Day() == 1 {
				if now.Hour() == int(r.ResetTime) {
					//发起重置
					var req model.ReqResetQuo
					var resp model.RespResetQuo
					req.SubCode = r.SubCode
					req.UserAddr = r.UserAddr
					req.KeyStore = keyString
					glog.Infoln("handleOne***************** start")
					glog.Infoln("handleOne***************mount req:", req)
					if err := doPost("/resetquo", req, &resp); err != nil {
						glog.Errorln("handleOne********************doPost:", err)
						return
					}
					glog.Infoln("handleOne*****************resp:", resp)
					if mydb := DB.Model(&r).Update("last_time", now); mydb.Error != nil {
						glog.Errorln("handOne************db err:", mydb.Error)
						return
					}
				}
			}
		}
	}
}
