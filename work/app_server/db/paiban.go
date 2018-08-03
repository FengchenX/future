package db

import (
	"strconv"
	"time"
	//"sub_account_service/app_server/protocol"
	"encoding/json"
	"sub_account_service/app_server_v2/model"

	"github.com/golang/glog"
)

// 发布信息结构体
type PaiBan struct {
	Id         uint      "gorm:PRIMARY KEY"
	SubCode    string    `gorm:"type:text;not null` // 排班编号
	Publisher  string    "gorm:not null"            // 发布者
	Detall     string    `gorm:"type:text;not null` // 参与者&分配比例
	CreateTime time.Time "gorm:not null"            // 创建时间
	Status     int64     "gorm:not null"            // 是否当前
	Times      int64     "gorm:not null"            // Hash
}

func InsertPaiBan(a []model.PaiBan, sn, usr string, used bool) {

	jbuf, err := json.Marshal(a)

	glog.Infoln("json", string(jbuf))
	if err != nil {
		glog.Errorln("SetSchedule err", err)
	}
	status, _ := strconv.Atoi(S_PENDING)
	// updating old status
	if used {
		db := DbClient.Client.Model(&PaiBan{}).
			Where("status < ? AND sub_code = ?", S_FAIL, sn).Update(&PaiBan{
			SubCode:    sn,
			Publisher:  usr,
			Detall:     string(jbuf),
			CreateTime: time.Now(),
			Status:     int64(status),
			Times:      0,
		})
		if db.Error != nil {
			glog.Errorln("UpdateSchedule err used", sn, db.Error)
		}
		return
	}

	// inserting New one
	db := DbClient.Client.Create(&PaiBan{
		SubCode:    sn,
		Publisher:  usr,
		Detall:     string(jbuf),
		CreateTime: time.Now(),
		Status:     int64(status),
		Times:      0,
	})

	if db.Error != nil {
		glog.Errorln("SetSchedule err", sn, db.Error)
	}
}

func GetPaiBan(where string, args ...interface{}) *PaiBan {
	p := &PaiBan{}
	db := DbClient.Client.Model(&PaiBan{}).Where(where, args...).Find(p)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
	return p
}
