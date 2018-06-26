package main

import (
	"github.com/feng/future/design/factory/factory"
)

func main() {
	var pf factory.PenFactory
	p := pf.Produce("brush")
	p.Write()
}