package main

import (
	"flag"
	"github.com/golang/glog"
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
	defer glog.Flush()
	flag.Parse()
	p := Point {
		x: 10,
		y: 20,
	}
	glog.Infoln(p)
}
type Point struct {
	x int
	y int
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
