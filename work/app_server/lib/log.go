package lib

import "fmt"

func Log(sign, addr, msg string) string {
	if addr == "" {
		return fmt.Sprintf("[type : %v  act: %v]  ", sign, msg)
	}
	return fmt.Sprintf("[type : %v user: %v  act: %v]   ", sign, addr, msg)
}

func Loger(sign, msg string) string {
	if sign == "api" {
		return fmt.Sprintf("< --API-- act: %v >   ", msg)
	}
	return fmt.Sprintf("[type : %v act: %v]   ", sign, msg)
}
