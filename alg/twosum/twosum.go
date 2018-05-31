package twosum

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i:=0; i<len(nums) - 2; i++ {
		for j:=i+1; j<len(nums) - 1; j++ {
			if nums[i]+nums[j]==target {
				return []int{i,j} 
			}
		}
	}
	fmt.Println("not find")
	return nil
}