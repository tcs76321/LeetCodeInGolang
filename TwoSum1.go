package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) [2]int {
	var i, k int
	result := [2]int{-1, -1}
	for i = 0; i < (len(nums) - 1); i++ {
		for k = 1; k < (len(nums)); k++ {
			if nums[i]+nums[k] == target {
				result[0] = i
				result[1] = k
				return result
			}
		}
	}
	return result
}

// this is correct code, needs to be modified for leetcode though, [2]int -> []int
// runtime beats 14.00 % of golang submissions
// 
