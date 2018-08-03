package db

import (
	"strconv"
	"github.com/golang/glog"
	"sub_account_service/app_server_v2/model"
	"time"
	"github.com/jinzhu/gorm"
)

var (
	S_SUCCESS   = "0"
	S_PENDING   = "1"
	S_FAIL      = "2"
	S_OUTMODED  = "3"
	Ss_SUCCESS  = 0
	Ss_PENDING  = 1
	Ss_FAIL     = 2
	Ss_OUTMODED = 3
)

//Schedule 发布信息结构体
type Schedule struct {
	gorm.Model
	//ID         uint      `gorm:"PRIMARY KEY"`
	SubCode    string    `gorm:"type:text;not null"`             // 排班编号
	Publisher  string    `gorm:"not null"`                        // 发布者
	CreateTime time.Time `gorm:"not null default CURRENT_TIESTAMP"`                        // 创建时间
	Status     int64     `gorm:"not null"`                        // 是否当前
	Menthon    string    `gorm:"not null"`                        // 備註信息
	Hash       string    `gorm:"not null"`                        // Hash
	Times      int64     `gorm:"not null"`                        // Hash
	Rs []model.Rs 		 `gorm:"ForeignKey:ScheID;AssociationForeignKey:ID"`
	PaiBan []model.PaiBan  `gorm:"ForeignKey:ScheID;AssociationForeignKey:ID"`
}



//InsertSchedule 插入分配表
func InsertSchedule(a []model.Rs, sn, usr, hash, methon string, used bool) {

	DB := DbClient.Client
	if DB == nil {
		glog.Errorln("InserSchedule***************db is nil")
	}

	var schedules []Schedule
	if mydb := DB.Where("sub_code = ?", sn).Find(&schedules); mydb.Error != nil {
		glog.Errorln("IinsertSchedule******************db err", mydb.Error)
		return
	}
	if len(schedules) > 0 {
		//do update
		tx := DB.Begin()
		for _, sche := range schedules {
			temp := Schedule{
				SubCode: sn,
				Publisher: usr, 
				Hash: hash,
				Menthon: methon,
			}
			glog.Infoln("InsertSchedule*************old schedule:", sche, "new schedule", sn,"|", usr, "|", hash, "|", methon)
			if mydb := tx.Model(&sche).Updates(temp); mydb.Error != nil {
				glog.Errorln("InsertSchedule******************db err:", mydb.Error)
				continue
			}
			for i := 0; i < len(a); i++ {
				a[i].ScheID = sche.ID
			}

			glog.Infoln("InsertSchedule*************rs:", a)
			tx := DB.Begin()
			for _, v := range a {
				if mydb := tx.Create(&v); mydb.Error != nil {
					glog.Errorln("InsertSchedule****************db err:", mydb.Error)
				}
			}
			tx.Commit()
		}
		tx.Commit()
		return
	}

	status, _ := strconv.Atoi(S_PENDING)
	// inserting New one
	schedule := Schedule {
		SubCode: sn,
		Publisher: usr,
		CreateTime: time.Now(),
		Status: int64(status),
		Hash: hash,
		Times: 0,
		Menthon: methon,
	}
	
	glog.Infoln("InsertSchedule****************schedule:", schedule)
	mydb := DB.Create(&schedule)
	if mydb.Error != nil {
		glog.Errorln("InsertSchedule******************err:", sn, mydb.Error)
	}
	for i := 0; i < len(a); i++ {
		a[i].ScheID = schedule.ID
	}
	glog.Infoln("InsertSchedule************rs:", a)
	tx := DB.Begin()
	for _, v := range a {
		v.Accounts = ""  //防止客户端又将用户地址发过来
		if mydb := tx.Create(&v); mydb.Error != nil {
			glog.Errorln("SetSchedule****************db err:", mydb.Error)
		}
	}
	tx.Commit()
}

func GetSchedulesLimit(where string, limit, offset int, args ...string) ([]*Schedule, int) {
	ss := make([]*Schedule, 0)
	var count int

	DbClient.Client.Model(&Schedule{}).Where(where, args).Count(&count)

	db := DbClient.Client.Order("id desc").Where(where, args).Limit(limit).Offset(offset).Find(&ss)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
	return ss, count
}

func GetSchedules(where string, args ...interface{}) []*Schedule {
	ss := make([]*Schedule, 0)

	db := DbClient.Client.Model(&Schedule{}).Where(where, args...).Find(&ss)
	if db.Error != nil {
		glog.Errorln("GetSchedules err", db.Error)
	}
	return ss
}

func UpdatingSchedules(ss []*Schedule) {
	for _, s := range ss {
		if s.Times < 20 && s.Status != int64(Ss_FAIL) {
			db := DbClient.Client.Model(s).
				Where("sub_code = ?", s.SubCode).Update(map[string]interface{}{"status": S_SUCCESS, "times": s.Times + 1})
			if db.Error != nil {
				glog.Errorln("UpdatingSchedules err", db.Error)
				continue
			}
		} else if s.Times == 20 && s.Status != int64(Ss_FAIL) {
			db := DbClient.Client.Model(s).
				Where("status = ? AND sub_code = ?", S_PENDING, s.SubCode).Update("status", S_FAIL)
			if db.Error != nil {
				glog.Errorln("UpdatingSchedules err", db.Error)
				continue
			}
		}

	}
}
