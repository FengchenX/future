

package main

import (
	"fmt"

)

func main() {
	l1 := &ListNode{1, nil}
	l12 := &ListNode{2, nil}
	l13 := &ListNode{3, nil}
	l14 := &ListNode{4, nil}
	l1.Next = l12
	l12.Next = l13
	l13.Next = l14
	l2 := &ListNode{4, nil}
	l22 := &ListNode{5, nil}
	l23 := &ListNode{6, nil}
	l24 := &ListNode{7, nil}
	l2.Next = l22
	l22.Next = l23
	l23.Next = l24

	temp := addTwoNumbers(l1, l2)
	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

// ListNode 链表节点
type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lx := reverList(l1)
	ly := reverList(l2)

	var list *ListNode
	var head *ListNode
	for i:=0; i<listLen(lx); i++ {
		val := nodei(lx, i).Val + nodei(ly, i).Val
		node := &ListNode{val, nil}
		if i == 0 {
			head = node
			list = node
		}
		list.Next = node
		list = node
	}
	return head
}

//链表反转
func reverList(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	temp := reverList(list.Next)
	x := temp
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = list
	list.Next = nil 
	return x
}

func listLen(list *ListNode) int {
	var len int
	for list != nil {
		len = len + 1
		list = list.Next
	}
	return len
}

func nodei(list *ListNode, i int) *ListNode {
	if i >= listLen(list) {
		return nil
	}
	for j:=0; j != i; j++ {
		list = list.Next
	}
	return list
}




