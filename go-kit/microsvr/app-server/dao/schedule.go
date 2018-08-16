package dao

import (
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"github.com/sirupsen/logrus"
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
	if mydb := DBInst().Delete(obj); mydb.Error != nil {
		logrus.Errorln("DelOne***************db err", mydb.Error)
	}
}