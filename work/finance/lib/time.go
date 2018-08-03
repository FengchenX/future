package lib

import (
	"github.com/golang/glog"
	"time"
)

var timeLayout = "2006-01-02 15:04:05" //转化所需模板
func ParseDateTimeStr(time_str string) time.Time {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	t, err := time.ParseInLocation(timeLayout, time_str, loc)
	if err != nil {
		glog.Infoln("uploading time fail ", err, time_str)
		return time.Now()
	}
	return t
}
