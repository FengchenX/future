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

	
}