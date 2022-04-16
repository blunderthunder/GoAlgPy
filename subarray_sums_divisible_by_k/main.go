package main

import "fmt"

func subarraysDivByK(nums []int, k int) int {
	mod := make([]int, k+1)
	cumsum := 0

	for i := 0; i <= k; i++ {
		mod = append(mod, 0)
	}

	for i := 0; i < len(nums); i++ {
		cumsum += nums[i]
		mod[((cumsum%k)+k)%k] = mod[((cumsum%k)+k)%k] + 1
	}
	result := 0

	for i := 0; i < k; i++ {
		if mod[i] > 1 {
			result = result + (mod[i]*(mod[i]-1))/2
		}
	}
	result = result + mod[0]
	return result
}

func main() {
	fmt.Println("no of  subarray sums of array", []int{4, 5, 0, -2, -3, 1}, "divisible by", 5, "is ", subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
