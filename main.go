package main

import (
	"fmt"
	
)

func main() {
	var a []int	
	fmt.Println(len(a))
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
