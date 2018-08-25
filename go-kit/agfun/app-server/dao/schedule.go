package dao

import (
	"time"
	"strconv"
	"github.com/feng/future/go-kit/agfun/app-server/model"
	"github.com/sirupsen/logrus"
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

//GetSchedules 获取分配表
func GetSchedules(where string, args ...interface{}) []model.Schedule {
	var schedules []model.Schedule
	if mydb := DBInst().Where(where, args...).Find(&schedules); mydb.Error != nil {
		logrus.Errorln("GetSchedules err", mydb.Error)	
		return nil
	}
	return schedules
}

//GetRss 获取分配详情
func GetRss(sqlstr string, args ...interface{}) []model.Rs {
	var rss []model.Rs
	if mydb := DBInst().Where(sqlstr, args...).Find(&rss); mydb.Error != nil {
		logrus.Errorln("GetRss***************db err:", mydb.Error)
		return nil
	}
	return rss
}

//GetRssConnSchedule 获得分配详情
func GetRssConnSchedule(sqlstr string, args ...interface{}) []model.Rs {
	var rss []model.Rs
	if mydb := DBInst().Table("rs").Joins("join schedules on rs.sche_id = schedules.id").Where(sqlstr, args...); mydb.Error != nil {
		logrus.Errorln("GetRssConnSchedule************db err:", mydb.Error)
		return nil
	}
	return rss
}

//GetPaiBans 获取排班
func GetPaiBans(sqlstr string, args ...interface{}) []model.PaiBan {
	var paiBans []model.PaiBan
	if mydb := DBInst().Where(sqlstr, args...).Find(&paiBans); mydb.Error != nil {
		logrus.Errorln("GetRss***************db err:", mydb.Error)
		return nil
	}
	return paiBans
}

//DelOne 删除某个记录
func DelOne(obj interface{}) {
	logrus.Infoln("DelOne*******************obj:", obj)
	if mydb := DBInst().Delete(obj); mydb.Error != nil {
		logrus.Errorln("DelOne***************db err", mydb.Error)
	}
}

//InsertSchedule 插入分配表
func InsertSchedule(rss []model.Rs, subCode, userAddr, hash, message string) error {
	status, _ := strconv.Atoi(S_PENDING)
	schedule := model.Schedule {
		SubCode: subCode,
		Publisher: userAddr,
		CreateTime: time.Now(),
		Status: int64(status),
		Hash: hash,
		Times: 0,
		Menthon: message,
	}
	logrus.Infoln("InsertSchedule**************schedule:", schedule)
	tx := DBInst().Begin()
	if mydb := tx.Create(&schedule); mydb.Error != nil {
		logrus.Errorln("InsertSchedule**************db err:", subCode, mydb.Error)
		tx.Rollback()
		return mydb.Error
	}
	for _, rs := range rss {
		if mydb := tx.Create(&rs); mydb.Error != nil {
			tx.Rollback()
			logrus.Errorln("InsertSchedule**************db err:", mydb.Error)
			return mydb.Error
		}
	}
	tx.Commit()
	return nil
}

//UpdateSchedule 更新分配表
func UpdateSchedule(rss []model.Rs, subCode, userAddr, hash, message string, old *model.Schedule) error {
	logrus.Infoln("InsertSchedule***********old sche:", old, "new sche:", rss, userAddr, hash, message)
	tx := DBInst().Begin()
	temp := model.Schedule {
		SubCode: subCode,
		Publisher: userAddr,
		Hash: hash,
		Menthon: message,
	}
	if mydb := tx.Model(old).Updates(temp); mydb.Error != nil {
		logrus.Errorln("UpdateSchedule************")
	}
}