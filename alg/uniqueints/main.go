package main

import (
	"fmt"
)

func main() {
	var A = []int{9, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Println(uniqueInts(A))
}

func uniqueInts(A []int) []int {
	if len(A) == 0 || len(A) == 1 {
		a := make([]int, 0)
		a = append(a, A...)
		return a
	}
	for i := 1; i < len(A); i++ {
		key := A[i]
		for j := i - 1; j >= 0; j-- {
			if key == A[j] {
				A = append(A[:i], A[i+1:]...)
				i = i - 1
				break
			}
		}
	}
	a := make([]int, 0)
	a = append(a, A...)
	return a
}
