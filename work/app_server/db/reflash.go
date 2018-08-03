package db

import (
	"time"

	"github.com/golang/glog"
)

// 发布信息结构体
type Reflash struct {
	Id             uint      "gorm:PRIMARY KEY"
	SubCode        string    `gorm:"type:text;not null` // 排班编号
	UserAddr       string    "gorm:not null"            // 用户地址
	CreateTime     time.Time "gorm:not null"            // 创建时间
	UpdateTime     time.Time "gorm:not null"            // 更新时间
	Used           bool      "gorm:not null"            // 是否使用
	UpdateRealTime int       "gorm:not null"            // 是否使用
	UpdateWay      int       "gorm:not null"            // 是否使用
}

func InsertReflash(sn, usr string) {
	// inserting New one
	db := DbClient.Client.Create(&Reflash{
		SubCode:    sn,
		CreateTime: time.Now(),
		Used:       true,
	})
	if db.Error != nil {
		glog.Errorln("SetSchedule err", db.Error)
	}
}

func GetReflashs(where string, args ...string) map[int][]*Reflash {
	ss := make([]*Reflash, 0)

	db := DbClient.Client.Where(where, args).Find(&ss)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
	m := make(map[int][]*Reflash)
	for _, v := range ss {
		if _, exist := m[v.UpdateRealTime]; !exist {
			m[v.UpdateRealTime] = []*Reflash{v}
		} else {
			m[v.UpdateRealTime] = append(m[v.UpdateRealTime], v)
		}
	}
	return m
}

func UpdatingReflash(where, args, update, str string) {
	db := DbClient.Client.Model(&Reflash{}).Where(where, args).Update(update, str)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
}

func UpdatingReflash2(where, arg1, arg2, update, argu interface{}) {
	db := DbClient.Client.Model(&Reflash{}).Where(where, arg1, arg2).Update(update, argu)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
}
