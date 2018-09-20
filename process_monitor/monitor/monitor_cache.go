package monitor

import (
	"sub_account_service/process_monitor/config"
	"time"
)

// CacheInfo 缓存信息
type CacheInfo struct {
	ProcessName string
	StartPoints []*StartPoint
	StartTimes  int
	FailTimes   int
}

// StartPoint 重启点
type StartPoint struct {
	StartTime time.Time
	Success   bool
}

// Cache 缓存
var Cache = make(map[string]*CacheInfo)

// return true means need sendEmail
func restarted(processName string, success bool) bool {
	if Cache[processName] == nil {
		Cache[processName] = &CacheInfo{
			ProcessName: processName,
			StartPoints: make([]*StartPoint, 0),
			StartTimes:  0,
			FailTimes:   0,
		}
	}

	cacheInfo := Cache[processName]
	cacheInfo.StartTimes++
	cacheInfo.StartPoints = append(cacheInfo.StartPoints, &StartPoint{
		StartTime: time.Now(),
		Success:   success,
	})

	if success { //重启成功
		return true
	}
	//重启失败
	cacheInfo.FailTimes++
	if len(cacheInfo.StartPoints) == 0 {
		return true
	} else if time.Now().Unix()-cacheInfo.StartPoints[len(cacheInfo.StartPoints)-1].StartTime.Unix() > config.GetConfigInstance().Interval*1000*2 {
		//每三次检查才会发送一次邮件才发送邮件
		return true
	} else {
		return false
	}
}
