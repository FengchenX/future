package main

import (
	"fmt"
	"github.com/feng/future/design/factory/factory"
)

func main() {
	var pf factory.PenFactory
	p := pf.Produce("brush")
	p.Write()

	var of factory.OperationFactory

	o := of.Produce("*")
	result := o.Operation(10, 20)
	fmt.Println(result)

	var bp factory.BrushPen
	f(bp)


	var u = factory.User{"xxxx", 29}
	var fac factory.AccessFactory

	var absfactory factory.DBAbsFactory

	absfactory = fac
	iu := absfactory.Produce()

	iu.Insert(u)
	
}

func f(af factory.AbstractFactory) {
	p := af.Produce()
	p.Write()
}

