package main

import (
	"fmt"
	"time"
	"math"
)

func main() {

	now := time.Now().Unix()
	sdfs := time.Unix(now, 0)
	fmt.Println(sdfs)
	assigner := Assigner {
		LastTime: "2018-08-31 00:04:05",
		Quota: 0,
		QuotaWay: 1,
		ResetTime: 5,
	}
	var allocation Allocation
	allocation.Assigners = append(allocation.Assigners, assigner)
	for i := range allocation.Assigners {
		quota := Quota {
			Money: 10,
		}
		nowTime := "2018-09-01 11:04:05"
		_, err := quotaAlgorithm(&allocation.Assigners[i], &quota, nowTime)
		if err != nil {
			fmt.Println(err)
		}
		
	}
	for _, allo := range allocation.Assigners {
		fmt.Println("################", allo.LastTime)
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
			lastTime, _ := time.Parse("2006-01-02 15:04:05", assigner.LastTime)
			if math.Abs(float64(now.Day() - lastTime.Day())) >= 1 {
				//发起重置
				fmt.Println("resetQuo*********day start")
				assigner.LastTime = nowTime
				assigner.Quota = 0

				quota.Money = 0

				hadChange = true
				return hadChange, nil
			}
		}
	} else if assigner.QuotaWay == 2 {
		//按月重置
		lastTime, _ := time.Parse("2006-01-02 15:04:05", assigner.LastTime)
		if math.Abs(float64(now.Month()-lastTime.Month())) >= 1 {
			//发起重置
			fmt.Println("resetQuo*********month start")
			assigner.LastTime = nowTime
			assigner.Quota = 0

			quota.Money = 0

			hadChange = true
			return hadChange, nil
		}
	}
	return hadChange, nil
}

type Allocation struct {
	Assigners []Assigner
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
