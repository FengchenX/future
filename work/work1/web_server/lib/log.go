package lib

import "fmt"

func Log(sign, addr, msg string) string {
	if addr==""{
		return fmt.Sprintf("[类型 : %v  操作: %v]  ", sign, msg)
	}
	return fmt.Sprintf("[类型 : %v 用户: %v  操作: %v]   ", sign, addr, msg)
}

func Loger(sign, msg string) string {
	return fmt.Sprintf("[类型 : %v 操作: %v]   ", sign, msg)
}
