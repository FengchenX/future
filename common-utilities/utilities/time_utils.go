//author xinbing
//time 2018/8/27 10:03
//时间工具
package utilities

import "time"

const(
		dateTimePattern = "2006-01-02 15:04:05"
	) //转化所需模板
func ParseDateTimeStr(time_str string) (time.Time,error) {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	t, err := time.ParseInLocation(dateTimePattern, time_str, loc)
	return t, err
}

func ParseDateTimeStrWithDefault(time_str string,defaultTime time.Time) (time.Time) {
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(dateTimePattern, time_str, loc)
	if err != nil {
		return defaultTime
	}
	return t
}

func FormatDateTime(t time.Time) string {
	return t.Format(dateTimePattern)
}
