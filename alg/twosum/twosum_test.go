
package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 15}
	var target = 13
	a := twoSum(nums,target)
	if a[0]==0 && a[1]==2{
		t.Log(twoSum(nums, target))
	} else {
		t.Error("算法有误")
	}
}