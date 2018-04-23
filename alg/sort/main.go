
package main

import (
	"fmt"

)

func main() {
	//sort_insertSort()
	test_binSearch()
}


func sort_insertSort() {
	var A = []int {3,2,4,6,1,7,5,9,8}
	insertSort(A)
	for a := range A {
		fmt.Println(a)
	}
}
func insertSort(A []int) {
	for j := 1; j<len(A); j++ {
		key:= A[j]
		i:= j-1
		for i>=0 && key < A[i] {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
}

func partition(key int, A []int, lo int, hi int) (int,int) {
	mid:= (lo+hi)/2
	if key > A[mid] {
		lo = mid
	} else {
		hi = mid
	}
	return lo,hi
}

func binSearch(key int, A []int, lo int, hi int) int {
	if lo != hi-1 {
		lo,hi = partition(key,A,lo,hi)
		return binSearch(key,A,lo,hi)
	} else {
		if key != A[hi] {
			return -1
		}
		return hi
	}
}

func test_binSearch() {
	var A = []int {3,2,4,6,1,7,5,9,8,15}
	insertSort(A)
	fmt.Println(binSearch(15,A,0,len(A)-1))
}
