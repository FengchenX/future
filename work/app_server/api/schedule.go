package api

import (
	"strings"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/db"
	"sub_account_service/app_server_v2/lib"

	"github.com/golang/glog"
	"net/http"
	"sub_account_service/app_server_v2/model"
	"time"

	"github.com/gin-gonic/gin"
)

//SetSchedule 发布分配表服务
func SetSchedule(c *gin.Context) {
	var req model.ReqSchedule
	var resp model.RespSchedule
	if err := lib.ParseReq(c, "SetSchedule", &req); err != nil {
		return
	}

	glog.Infoln("SetSchedule****************req:", req)

	keyString := lib.KeyMap.Calc(req.UserKeyStore, req.UserParse)
	req.KeyString = keyString

	req.ScheduleName = strings.Trim(req.ScheduleName, " ")
	used := false
	sn := ""
	if req.ScheduleName != "" {
		arr := db.GetSchedules("sub_code = ?", req.ScheduleName)
		used = true
		if len(arr) == 0 {
			glog.Errorln("SetSchedule**************ScheduleName:", req.ScheduleName)
			resp.StatusCode = S_NO_FIND
			c.JSON(http.StatusOK, resp)
			return
		}
		sn = req.ScheduleName
	} else {
		req1 := model.ReqNewScheduleId{
			Index: "AC",
		}
		var resp1 model.RespNewScheduleId
		if err := doPost("/newscheduleid", req1, &resp1); err != nil {
			glog.Errorln("SetSchedule*****************doPost", err)
			resp1.Msg = err.Error()
			resp1.StatusCode = S_DOPOST_ERR
			c.JSON(http.StatusOK, resp1)
			return
		}
		sn = resp1.ScheduleName
	}

	sum := float64(0)
	allByCount := true
	for _, v := range req.Rss {
		if v.SubWay == 0 {
			sum += v.Radios
			allByCount = false
		}
	}
	if sum != 100 {
		glog.Errorln("SetSchedule****************** sum is not 100, sum:", sum)
		resp.StatusCode = S_SUM_NOT_100
		resp.Msg = "比例之和要等于100"
		c.JSON(http.StatusOK, resp)
		return
	}

	if allByCount {
		glog.Errorln("SetSchedule***************cannot all are quota, Rss:", req.Rss)
		resp.StatusCode = S_ALL_QUO
		resp.Msg = "不可以全部是定额"
		c.JSON(http.StatusOK, resp)
		return
	}

	if len(req.Rss) == 0 {
		glog.Errorln("SetSchedule**************no set job, rss:",req.Rss)
		resp.StatusCode = S_NO_JOB
		resp.Msg = "没有设置工作"
		c.JSON(http.StatusOK, resp)
		return
	}

	for i := 0; i < len(req.Rss); i++ {
		for j := i + 1; j < len(req.Rss); j++ {
			if req.Rss[i].Job == req.Rss[j].Job {
				resp.StatusCode = S_TWO_JOB_SAME
				resp.Msg = "不能设置两个相同名字的工作"
				c.JSON(http.StatusOK, resp)
				return
			}
		}
	}

	req.ScheduleName = sn
	var resp2 model.RespSchedule
	glog.Infoln("SetSchedule******************req:", req)
	if err := doPost("/setschedule", req, &resp2); err != nil {
		glog.Errorln("SetSchedule******************doPost", err)
		resp.Msg = err.Error()
		resp.StatusCode = S_DOPOST_ERR
		c.JSON(http.StatusOK, resp)
		return
	}
	glog.Infoln("SetSchedule***************doPost:", resp2)
	if resp2.StatusCode != 0 {
		glog.Errorln("SetSchedule***************doPost", resp2)
		resp = resp2
		resp.StatusCode = S_API_RETURN_ERR
		c.JSON(http.StatusOK, resp)
		return
	}

	clearOld(req.ScheduleName)
	resp = resp2
	used = false
	db.InsertSchedule(req.Rss, req.ScheduleName,
		req.UserAddress, resp.Hash, req.Message, used)

	//添加重置
	var resets []db.Reset
	for _, rs := range req.Rss {
		var r db.Reset
		if rs.SubWay == 0 {
			continue
		}
		if rs.ResetWay == 2 {
			//每月
			now := time.Now()
			day := now.Day()
			dua := time.Duration(day - 1)
			x := time.Hour * (-1) * 24 * dua
			r.LastTime = now.Add(x)
			glog.Infoln("SetSchedule**********resetway=2 day:", day, "| lastTime:", r.LastTime)
		} else if rs.ResetWay == 1 {
			now := time.Now()
			h := now.Hour()
			if h > int(rs.ResetTime) {
				dua := time.Duration(h - int(rs.ResetTime))
				r.LastTime = now.Add(-dua * time.Hour)
			} else {
				dua := time.Duration(h + 24 - int(rs.ResetTime))
				r.LastTime = now.Add(-dua * time.Hour)
			}
			glog.Infoln("SetSchedule*******resetway=1 lasttime", r.LastTime, "| resettime:", rs.ResetTime)
		}
		r.Job = rs.Job
		r.ResetTime = rs.ResetTime
		r.ResetWay = rs.ResetWay
		r.SubCode = req.ScheduleName
		//r.UserAddr = rs.Accounts
		r.Publisher = req.UserAddress
		resets = append(resets, r)
	}
	Adds(req.UserAddress, req.ScheduleName, resets)

	glog.Infoln("SetSchedule************************resp: ", resp)
	resp.StatusCode = S_SUCCESS
	c.JSON(http.StatusOK, resp)
}

func clearOld(subCode string) {
	if subCode == "" {
		return
	}
	DB := db.DbClient.Client
	if DB == nil {
		return
	}
	var rs []model.Rs
	if mydb := DB.Table("rs").Joins("join schedules on rs.sche_id = schedules.id").Where("sub_code = ?", subCode).Find(&rs); mydb.Error != nil {
		glog.Errorln("clearOld**************db err:", mydb.Error)
		return
	}
	glog.Infoln("clearOld******************rs:", rs)
	tx := DB.Begin()
	for _, r := range rs {
		if mydb := tx.Delete(&r); mydb.Error != nil {
			glog.Errorln("clearOld*******************db err", mydb.Error)
			return
		}
	}
	tx.Commit()

	var paiBans []model.PaiBan
	if mydb := DB.Table("pai_bans").Joins("join schedules on pai_bans.sche_id = schedules.id").Where("sub_code = ?", subCode).Find(&paiBans); mydb.Error != nil {
		glog.Errorln("clearOld******************db err:", mydb.Error)
		return
	}
	glog.Infoln("clearOld****************paiBan:", paiBans)
	tx = DB.Begin()
	for _, paiBan := range paiBans {
		if mydb := tx.Delete(&paiBan); mydb.Error != nil {
			glog.Errorln("clearOld**********************db err", mydb.Error)
			return
		}
	}
	tx.Commit()
}

//GetAllSchedule 查询排班接口
func GetAllSchedule(c *gin.Context) {
	var req model.ReqGetAllSchedule
	var resp model.RespGetAllSchedule
	if err := lib.ParseReq(c, "GetAllSchedule", &req); err != nil {
		return
	}
	glog.Infoln("GetAllSchedule**************req:", req)

	DB := db.DbClient.Client
	if DB == nil {
		glog.Errorln("GetAllSchedule************db is nil")
	}

	page := int(req.Pages)
	if page == 0 {
		page = 1
	}
	limit := config.Optional.CheckLimit

	ss, count := db.GetSchedulesLimit("status < ?", limit, limit*(page-1), db.S_FAIL)
	ress := make([]model.Sd, 0)
	for _, v := range ss {
		rss := make([]model.Rs, 0)

		mydb := DB.Select("*").
			Table("rs").
			Where("sche_id = ?", v.ID).
			Find(&rss)
		if mydb.Error != nil {
			glog.Errorln("GetAllSchedule************db err:", mydb.Error)
			resp.Msg = mydb.Error.Error()
			resp.StatusCode = AS_DB_ERR
			c.JSON(http.StatusOK, resp)
		}

		uas := make([]model.UserAccount, 0)

		flag := false
		for i, rs := range rss {
			var resp1 model.RespGetAccount
			var req1 model.ReqGetAccount
			req1.UserAddress = rs.Accounts
			if rs.Accounts != "" {
				flag = true
				glog.Infoln("GetAllSchedule*****************req1:", req1)
				if err := doPost("/getaccount", req1, &resp1); err != nil {
					glog.Errorln("GetAllSchedule***************doPost", err)
					resp.Msg = err.Error()
					resp.StatusCode = AS_DOPOST_ERR
					c.JSON(http.StatusOK, resp)
					return
				}
				glog.Infoln("GetAllSchedule*******************resp1:", resp1)
				uas = append(uas, resp1.UserAccount)
			} else {
				var ua model.UserAccount
				ua.Address = ""
				ua.Alipay = ""
				ua.BankCard = ""
				ua.Name = ""
				ua.Telephone = ""
				ua.WeChat = ""
				uas = append(uas, ua)
			}

			if rs.SubWay != 0 {
				var req2 model.ReqGetQuo
				var resp2 model.RespGetQuo
				//req2.UserAddr = resp1.UserAccount.Address
				req2.SubCode = v.SubCode
				req2.UserAddr = rs.Accounts
				glog.Infoln("GetAllSchedule********************subCode, userAddr:", v.SubCode, rs.Accounts)
				if err := doPost("/getquo", req2, &resp2); err != nil {
					glog.Errorln("GetAllSchedule***************doPost", err)
					resp.Msg = err.Error()
					resp.StatusCode = AS_DOPOST_ERR
					c.JSON(http.StatusOK, resp)
					return
				}
				if resp2.StatusCode != 0 {
					glog.Warningln("GetAllSchedule*************warn resp2:", resp2)
				}
				glog.Infoln("GetAllSchedule*************money:", resp2.Money)
				//rss[i].GetMoney = float64(resp2.Money) / 10000
				rss[i].GetMoney = resp2.Money
			}

		}
		res := model.Sd{
			Owner:        v.Publisher,
			CreateTime:   v.CreateTime.Unix(),
			Rss:          rss,
			Status:       uint32(v.Status),
			UserAccounts: uas,
			SubCode:      v.SubCode,
			Message:      v.Menthon,
			HasPaiBan:    flag,
		}
		ress = append(ress, res)
	}

	pageC := uint64(count) / uint64(limit)
	if uint64(count)%uint64(limit) > 0 {
		pageC++
	}

	resp.StatusCode = AS_SUCCESS
	resp.Schedules = ress
	resp.Pages = req.Pages
	resp.PagesCount = pageC
	resp.Msg = "Success"
	glog.Infoln("GetAllSchedule****************resp:", resp)
	c.JSON(http.StatusOK, resp)
}

//SetPaiBan 2.7.发布排班  <--（客户端）v2新增  todo 需要判断是否包含自己
func SetPaiBan(c *gin.Context) {
	var req model.ReqSetPaiBan
	var resp model.RespSetPaiBan
	if err := lib.ParseReq(c, "SetPaiBan", &req); err != nil {
		return
	}
	glog.Infoln("SetPaiBan************************req: ", req)
	keyString := lib.KeyMap.Calc(req.UserKeyStore, req.UserParse)
	req.KeyString = keyString
	glog.Infoln("SetPaiBan******************keystring:", req.KeyString)

	req.OwnerAddress = strings.Trim(req.OwnerAddress, " ")
	req.SubCode = strings.Trim(req.SubCode, " ")
	isHadSelf := false
	isHadNil := false
	ownerRole := ""
	for i, paiBan := range req.PaiBans {
		if strings.ToLower(strings.Trim(paiBan.UserAddress, " ")) == strings.ToLower(strings.Trim(req.OwnerAddress, " ")) {
			isHadSelf = true
			ownerRole = paiBan.JobName
		}
		if paiBan.UserAddress == "" {
			isHadNil = true
		}
		req.PaiBans[i].UserAddress = strings.Trim(paiBan.UserAddress, " ")
	}
	if !isHadSelf {
		glog.Errorln("SetPaiBan**************** no include self")
		resp.Msg = "发布排班必须包含自己"
		resp.StatusCode = PB_NO_YOURSELF //todo
		c.JSON(http.StatusOK, resp)
		return
	}
	if isHadNil {
		glog.Errorln("SetPaiBan**************** addr had nil")
		resp.Msg = "某些排班地址未设置"
		resp.StatusCode = PB_HAD_NIL //todo
		c.JSON(http.StatusOK, resp)
		return
	}

	DB := db.DbClient.Client
	if DB == nil {
		glog.Errorln(" SetPaiBan********************db is nil")
		return
	}
	var schedule db.Schedule
	if mydb := DB.Where("sub_code = ?", req.SubCode).Preload("Rs").First(&schedule); mydb.Error != nil {
		glog.Errorln("  SetPaiBan*****************db err: ", mydb.Error)
		resp.StatusCode = PB_DB_ERR
		resp.Msg = mydb.Error.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	var rss []model.Rs
	if mydb := DB.Where("sche_id = ?", schedule.ID).Find(&rss); mydb.Error != nil {
		glog.Errorln("SetPaiBan*************************db err:", mydb.Error)
		resp.StatusCode = PB_DB_ERR
		resp.Msg = mydb.Error.Error()
		return
	}
	if len(rss) != len(req.PaiBans) {
		glog.Errorln("SetPaiBan******************paiban len not equal schedule")
		resp.StatusCode = PB_SCH_PB_LEN_NOTSAME
		resp.Msg = "排班人数和分配表不一致"
		c.JSON(http.StatusOK, resp)
		return
	}

	var d_level int64 = 99
	var d_flag bool = false

	for sdx := 0; sdx < len(schedule.Rs); sdx++ {
		if schedule.Rs[sdx].SubWay > 0 {
			// 有定额分配的
			//d_level = schedule.Rs[sdx].Level
			if d_level > schedule.Rs[sdx].Level {
				d_level = schedule.Rs[sdx].Level
			}
			d_flag = true
		}
	}

	glog.Infoln("SetPaiBan d_level:", d_level, "ownerRole:", ownerRole)

	if d_flag == true {
		for sdx := 0; sdx < len(schedule.Rs); sdx++ {
			if schedule.Rs[sdx].Job == ownerRole {
				// 有定额分配的
				glog.Infoln("SetPaiBan for d_level:", d_level, "schedule.Rs[sdx].Job:", schedule.Rs[sdx].Job, "schedule.Rs[sdx].Level:", schedule.Rs[sdx].Level)

				if d_level > schedule.Rs[sdx].Level {
					glog.Errorln("SetPaiBan fail set paiban owner level Put behind the quota.")
					resp.StatusCode = 70005
					resp.Msg = "set paiban owner level Put behind the quota."
					c.JSON(http.StatusOK, resp)
					return
				}
			}
		}
	}

	var resp1 model.RespSetPaiBan
	glog.Infoln("SetPaiBan****************req:", req)
	if err := doPost("/setpaiban", req, &resp1); err != nil {
		glog.Errorln("SetPaiBan*******************doPost", err)
		resp = resp1
		resp.StatusCode = PB_DOPOST_ERR
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	glog.Infoln("SetPaiBan**************resp:", resp1)

	resp = resp1

	//修改schedule表将对应的account填上
	go setScheAccount(req.OwnerAddress, req.SubCode, req.PaiBans)
	glog.Infoln("SetPaiBan********************resp: ", resp)
	if resp.StatusCode == 0 {
		resp.StatusCode = PB_SUCCESS
	} else {
		resp.StatusCode = PB_API_RETURN_ERR
	}
	c.JSON(http.StatusOK, resp)
}

func setScheAccount(publisher, subCode string, paiBans []model.PaiBan) {
	DB := db.DbClient.Client
	if DB == nil {
		glog.Errorln(" setScheAccount********************db is nil")
		return
	}

	var schedule db.Schedule
	if mydb := DB.Where("sub_code = ?", subCode).First(&schedule); mydb.Error != nil {
		glog.Errorln(" setScheAccount*****************db err: ", mydb.Error)
		return
	}

	var rss []model.Rs
	if mydb := DB.Where("sche_id = ?", schedule.ID).Find(&rss); mydb.Error != nil {
		glog.Errorln("setScheAccount*******************db err:", mydb.Error)
	}
	for i, rs := range rss {
		for _, paiBan := range paiBans {
			if rs.Job == paiBan.JobName {
				rss[i].Accounts = paiBan.UserAddress
				if mydb := DB.Model(&rs).Update("accounts", paiBan.UserAddress); mydb.Error != nil {
					glog.Errorln("setScheAccount*******************db err:", mydb.Error)
					continue
				}
				break
			}
		}
	}


	//todo 添加更新别直接写数据库
	tx := DB.Begin()
	for _, paiBan := range paiBans {
		paiBan.ScheID = schedule.ID
		tx.Create(&paiBan)
	}
	tx.Commit()
	Updates(publisher, subCode, paiBans)
}

//GetPaiBan 2.8.查询排班  <--（客户端）v2新增
func GetPaiBan(c *gin.Context) {
	var req model.ReqGetPaiBan
	var resp model.RespGetPaiBan
	if err := lib.ParseReq(c, "GetPaiBan", &req); err != nil {
		return
	}
	glog.Infoln(" GetPaiBan**********************req: ", req)
	if err := doPost("/getpaiban", req, &resp); err != nil {
		glog.Errorln("GetPaiBan********************doPost", err)
		resp.Msg = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//resp.StatusCode = 0
	glog.Infoln("GetPaiBan***********************resp: ", resp)
	c.JSON(http.StatusOK, resp)
}
