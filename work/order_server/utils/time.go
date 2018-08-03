package utils

import (
	"time"
)

func ParseDateTimeStr(time_str string) (t time.Time,err error) {
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	t, err = time.ParseInLocation(timeLayout, time_str, loc)
	return
}
