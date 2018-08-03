package test

import (
	"testing"
	"fmt"
)

type element struct {
	c chan int
}
var m map[string]element

func TestT(t *testing.T) {
	ele := element{
		c:make(chan int,1),
	}
	//go func() {
	//	ele.c <- 1
	//}()
	select {
	case<- ele.c:
		fmt.Println("----------")
	default:
		fmt.Println(",,,,,,,,,")
	}
	fmt.Println("121212")
}

