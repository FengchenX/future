package main

import "github.com/feng/future/design/adapter/adapter"

func main() {
	//adapter.AdapterTest()

	// var b adapter.NonA
	// nty := adapter.AdapterNonToYes{b}
	// f(nty)

	b := adapter.Fowards{}
	b.Player("巴蒂尔")
	f2(&b)

	m := adapter.Guards{}
	m.Player("麦克格雷迪")
	f2(&m)

	ym := adapter.Translator{}
	ym.Player("姚明")
	f2(&ym)

}

func f(battery adapter.IReBattery) {
	battery.Use()
	battery.Charge()
}


func f2(player adapter.IPlayer) {
	player.Attack()
	player.Defense()
}