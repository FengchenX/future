package main

import (
	"fmt"
	"time"
)

func main() {
	// var a = []int{1, 2, 3}

	// f1(a)
	// fmt.Println(a)
	// f2(&a)
	// fmt.Println(a)
	// num := 89.46
	// fmt.Println(int(num*10000))
	// fmt.Println(int(num*100*100))
	var a = 10
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("tick befor")
			if a == 10 {
				continue
			}
			fmt.Println("tick after")
		default:
			fmt.Println("default")
			time.Sleep(3 * time.Second)
		}
	}
}

// func f1(x interface{}) {
// 	v := x.([]int)
// 	for j := 0; j < 3; j++ {
// 		v[j] = v[j] + 3
// 	}
// 	for i := 0; i < 10; i++ {
// 		v = append(v, i)
// 	}
// }

// func f2(x *[]int) {

// 	for i := 0; i < 10; i++ {
// 		*x = append(*x, i)
// 	}
// }
