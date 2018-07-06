package main

import (
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
}

func f(soft bridge.HandsetSoft, hb bridge.HandsetBrand) {
	hb.SetHandsetSoft(soft)
	hb.Run()
}
