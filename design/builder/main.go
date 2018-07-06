package main

import (
	"fmt"

	"github.com/feng/future/design/builder/builder"
)

func main() {
	var hb builder.Hamburger
	var fc builder.FriedChicken
	var bear builder.Beer
	var mb builder.MealBuilder
	fmt.Println(mb.MealTwo(&hb, &fc, &bear))
}
