package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pl, pr := 1, 1
	res := make([]int, len(nums))
	for idx := range res {
		res[idx] = 1
	}
	for idx := range nums {
		res[idx] *= pl
		res[n-idx-1] *= pr
		pl = pl * nums[idx]
		pr = pr * nums[n-idx-1]
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
