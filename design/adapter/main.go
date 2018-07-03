package main

import "github.com/feng/future/design/adapter/adapter"

func main() {
	//adapter.AdapterTest()

	var b adapter.NonA
	nty := adapter.AdapterNonToYes{b}
	f(nty)
}

func f(battery adapter.IReBattery) {
	battery.Use()
	battery.Charge()
}