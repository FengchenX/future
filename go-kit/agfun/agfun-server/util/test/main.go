package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "suio"
	if !IsTecentAccount(a) {
		fmt.Println("bushi")
	}
	fmt.Println("shi")
}

const tecentAccountRe = "^ibs.+"

//IsTecentAccount 判断是否是腾讯账户
func IsTecentAccount(account string) bool {
	re := regexp.MustCompile(tecentAccountRe)
	if account == "" {
		return false
	}
	return re.MatchString(account)
}
