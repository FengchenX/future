package service

import (
	"strings"
	"github.com/feng/future/go-kit/microsvr/app-server/model"
	"github.com/feng/future/go-kit/microsvr/app-server/util"
	"github.com/feng/future/go-kit/microsvr/app-server/dao"
	"github.com/sirupsen/logrus"
)

//SetSchedule 发布分配表
func (AppSvc) SetSchedule(userAddress, 
	userKeyStore, 
	userParse, 
	keyString, 
	scheduleName string, 
	Rss []model.Rs, 
	message string) (uint32, string, string, string) {
		if keyString == "" {
			keyString = util.KeyMap.Calc(userKeyStore, userParse)
		}
		scheduleName = strings.Trim(scheduleName, " ")

		var sqlstr string
		var params []interface{}
		sqlstr += "sub_code = ?"
		params = append(params, scheduleName)
		sn := ""
		if scheduleName != "" {
			schedules := dao.GetSchedules(sqlstr, params...)
			if len(schedules) == 0 {
				logrus.Errorln("SetSchedule****************ScheduleName:", scheduleName)	
				return 10001, "not find schedule name", "", ""
			}
			sn = scheduleName
		} else {
			var reqNewSchedule struct {
				Index string
			}
			var respNewSchedule struct {
				StatusCode   uint32
				ScheduleName string
				Msg          string
			}

			reqNewSchedule.Index = "AC"
			if err := doPost("/newscheduleid", reqNewSchedule, & respNewSchedule); err != nil {
				logrus.Errorln("SetSchedlule***************doPost", err)
				return 10001, "doPost err", "", ""
			}
			sn = respNewSchedule.ScheduleName
		}

		sum := float64(0)
		allByCount := true
		for _, rs := range Rss {
			if rs.SubWay == 0 {
				sum += rs.Radios
				allByCount = false
			}
		}
		if !util.IsEqual(sum, 100) {
			logrus.Errorln("SetSchedule************* sum is not 100, sum:", sum)
			return 10001, "比例之和要等于100", "", ""
		}

		if allByCount {
			logrus.Errorln("SetSchedule***********cannot all are quota, Rss: ", Rss)
			return 10001, "不可以全部是定额", "", ""
		}
		if len(Rss) == 0 {
			logrus.Errorln("SetSchedule***************no set job, rss: ", Rss)
			return 10001, "没有设置工作", "", ""
		}

		for i := 0; i < len(Rss); i++ {
			for j := i + 1; j < len(Rss); j++ {
				if Rss[i].Job == Rss[j].Job {
					return 10001, "不能设置两个相同名字的工作", "", ""
				}
			}
		}

		
	}

	