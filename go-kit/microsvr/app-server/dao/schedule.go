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