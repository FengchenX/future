


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



type Tree struct {
	key int
	left, right, p *Tree
}

func InorderTreeWalk(x *Tree) {
	if x != nil {
		InorderTreeWalk(x.left)
		fmt.Println(x.key)
		InorderTreeWalk(x.right)
	}
}

func TreeSearch(x *Tree, key int) *Tree {
	if x == nil || key == x.key {
		return x
	} 
	if key < x.key {
		return TreeSearch(x.left, key)
	} else {
		return TreeSearch(x.right, key)
	}
}
func TreeMinimum(x *Tree) *Tree {
	for x.left != nil {
		x = x.left
	}
	return x
}
func TreeMaximum(x *Tree) *Tree {
	for x.right != nil {
		x = x.right
	}
	return x
}
func TreeSuccessor(x *Tree) *Tree {
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

func TreeInsert(T *Tree, z *Tree) {
	//y := nil
	//x := T.root
}