package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxsofar, maxendinghere := nums[0], nums[0]
	for _, val := range nums[1:] {
		fmt.Println(val)
		maxendinghere = max(val, maxendinghere+val)
		maxsofar = max(maxsofar, maxendinghere)
	}
	return maxsofar
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
