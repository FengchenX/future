package main

import (
	"fmt"
	"log"

	T "gorgonia.org/gorgonia"
)

func main() {
	g := T.NewGraph()

	var x, y, z *T.Node
	var err error

	// define the expression
	x = T.NewScalar(g, T.Float64, T.WithName("x"))
	y = T.NewScalar(g, T.Float64, T.WithName("y"))
	//z, err = Add(x, y)
	z, err = T.Sub(x, y)
	if err != nil {
		log.Fatal(err)
	}

	// create a VM to run the program on
	machine := T.NewTapeMachine(g)

	// set initial values then run
	T.Let(x, 2.0)
	T.Let(y, 2.5)
	if machine.RunAll() != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", z.Value())
	// Output: 4.5
}