package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	
	assigner := Assigner {
		LastTime: "2018-08-31 00:04:05",
		Quota: 0,
		QuotaWay: 1,
		ResetTime: 5,
	}
	quota := Quota {
		Money: 10,
	}
	nowTime := "2018-09-01 11:04:05"
	change, err := quotaAlgorithm(&assigner, &quota, nowTime)
	if err != nil {
		fmt.Println(err)
	}
	if change {
		fmt.Println("change")
		fmt.Println(assigner.LastTime)
	}
	fmt.Println("*************")
}


func quotaAlgorithm(assigner *Assigner, quota *Quota,  nowTime string) (bool, error) {
	hadChange := false
	now, err := time.Parse("2006-01-02 15:04:05", nowTime)
	fmt.Println(nowTime, assigner.LastTime, assigner.QuotaWay)
	if err != nil {
		fmt.Println("resetQuo***********err:", err)
		return hadChange, err
	}
	
	if assigner.QuotaWay == 1 {
		//按天重置
		fmt.Println(now.Hour(), assigner.ResetTime)
		if now.Hour() >= int(assigner.ResetTime) {
			if math.Abs(float64(now.Day()-time.Unix(formToUnix(assigner.LastTime), 0).Day())) >= 1 {
				//发起重置
				fmt.Println("resetQuo*********day start")
				assigner.LastTime = unixToForm(now.Unix())
				assigner.Quota = 0

				quota.Money = 0

				hadChange = true
				return hadChange, nil
			}

		}
	} else if assigner.QuotaWay == 2 {
		//按月重置
		if math.Abs(float64(now.Month()-time.Unix(formToUnix(assigner.LastTime), 0).Month())) >= 1 {
			//发起重置
			fmt.Println("resetQuo*********month start")
			assigner.LastTime = unixToForm(now.Unix())
			assigner.Quota = 0

			quota.Money = 0

			hadChange = true
			return hadChange, nil
		}
	}
	return hadChange, nil
}

type Assigner struct {
	LastTime string
	Quota float64
	QuotaWay uint
	ResetTime uint
}

type Quota struct {
	Money float64
}

func formToUnix(t string) int64 {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//获取本地location
	loc, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                   //转化为时间戳 类型是int64
	return sr
}

func unixToForm(t int64) string {
	//获取本地location
	timeLayout := "2006-01-02 15:04:05" //转化所需模板

	//时间戳转日期
	dataTimeStr := time.Unix(t, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	return dataTimeStr
}