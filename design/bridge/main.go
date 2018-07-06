package main

import (
	"fmt"
	"github.com/feng/future/design/bridge/bridge"
)

func main() {
	bridge.TestBridge()

	game := bridge.HandsetGame{}
	book := bridge.HandsetAddrList{}
	m := &bridge.HandsetBrandM{}
	n := &bridge.HandsetBrandN{}
	f(game, m)
	f(game, n)
	f(book, m)
	f(book, n)


	fmt.Println("*************************")
	var sch bridge.SchA
	var chushi bridge.Chushi
	var fuwuyuan bridge.Fuwuyuan
	f2(&sch, chushi, fuwuyuan)
}

func f(soft bridge.HandsetSoft, hb bridge.HandsetBrand) {
	hb.SetHandsetSoft(soft)
	hb.Run()
}

func f2(sch bridge.ISchedule, emps ...bridge.IEmployee) {
	sch.SetEmp(emps...)
	sch.Do()
}
