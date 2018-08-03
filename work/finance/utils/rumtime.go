package utils

import (
	"runtime"
)

// 设置同时执行的cpu数量
func SetGOMAXPROCS() {
	cupNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cupNum)
}
