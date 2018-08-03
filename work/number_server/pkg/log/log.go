package log

import (
	"time"
	"fmt"
)

func Info(action string, data interface{})  {
	fmt.Printf("[ action: %s time: %s ]   info:%+v  \n",action,time.Now().Format("2006-01-02 15:04:05"),data)
}

func Error(action string, data interface{})  {
	fmt.Errorf("[ action: %s time: %s ]   info:%+v  \n",action,time.Now().Format("2006-01-02 15:04:05"),data)
}

func DbInfo(action string,curd string, data interface{},err error)  {
	fmt.Printf("[ action: %s time: %s curd:%s  ]   info:%+v  \n errmsg:%+v \n",action,time.Now().Format("2006-01-02 15:04:05"),curd,data,err)
}