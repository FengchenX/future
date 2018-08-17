package main

import (
	"fmt"
	
)

func main() {
	var a []int	
	a = append(a, 1,2,3)	
	a = a[:0]
	a = append(a, 4)
	fmt.Println(a)
}
