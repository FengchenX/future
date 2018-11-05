package time

import (
	"strconv"
	"time"
)

func GetTimeStd(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetDateStd(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetUnix() string {
	now := time.Now().Unix()
	return strconv.FormatInt(now, 10)
}

func GetUnixNano() string {
	now := time.Now().UnixNano()
	return strconv.FormatInt(now, 10)
}
