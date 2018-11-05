package main

import (
	"fmt"
	"grm-service/util"
)

func main() {
	a := "1 2345 "
	if util.IsNum(a) {
		fmt.Println("is num")
	} else {
		fmt.Println("not num")
	}
}
