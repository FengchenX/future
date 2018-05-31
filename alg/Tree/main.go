


package main


import (
	"fmt"

)


func main() {

}

/**
key >= left.key
key <= right.key
二叉树性质
*/



type TreeNode struct {
	key int
	left, right, p *TreeNode
}

func InorderTreeWalk(x *TreeNode) {
	if x != nil {
		InorderTreeWalk(x.left)
		fmt.Println(x.key)
		InorderTreeWalk(x.right)
	}
}

func TreeSearch(x *TreeNode, key int) *TreeNode {
	if x == nil || key == x.key {
		return x
	} 
	if key < x.key {
		return TreeSearch(x.left, key)
	} else {
		return TreeSearch(x.right, key)
	}
}
func TreeMinimum(x *TreeNode) *TreeNode {
	for x.left != nil {
		x = x.left
	}
	return x
}
func TreeMaximum(x *TreeNode) *TreeNode {
	for x.right != nil {
		x = x.right
	}
	return x
}

//查找后继
func TreeSuccessor(x *TreeNode) *TreeNode {
	if x.right!=nil {
		return TreeSuccessor(x.right)
	}
	y := x.p	
	for y != nil && x == y.right {
		x= y
		y=y.p
	}
	return y
}

type Tree struct {
	root *TreeNode

}
//插入
func TreeInsert(T *Tree, z *TreeNode) {
	var y *TreeNode
	x := T.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		T.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func Transplant(T *Tree, u, v *TreeNode) {
	
}