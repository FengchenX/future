package util
 
/*
#include "util.h"
*/
import "C"
 
import "fmt"
 
func GoSum(a,b int) int {
    s := C.sum(C.int(a),C.int(b))
	fmt.Println(s)
	return int(s)
}

func GoSub(a, b int) int {
	s := C.sub(C.int(a), C.int(b))
	fmt.Println(s)
	return int(s)
}