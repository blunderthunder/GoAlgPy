package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	if t == 0 {
		magic := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if j, exist := magic[nums[i]]; exist {
				if i-j <= k {
					return true
				}
			}
			magic[nums[i]] = i
		}
		return false
	}
	// t > 0 here
	// key: num/t, e.g., t=3, num=5, key=1; t=3, num=-1, key=-1
	// value: value[0] index, value[1] num
	magic := make(map[int][2]int)
	for index, num := range nums {
		// t=3, num=2, key=0, it' ok
		// t=3, num=-2, key=0, it' not ok, so we make num -= t-1, like:
		// t=3, num=[+3,+6), key=+1
		// t=3, num=[+0,+3), key=+0
		// t=3, num=[-3,-0), key=-1
		// t=3, num=[-6,-3), key=-2
		if num < 0 {
			num -= t - 1
		}
		if v, exist := magic[num/t]; exist && index-v[0] <= k {
			return true
		}
		if v, exist := magic[num/t+1]; exist && index-v[0] <= k && v[1]-num <= t {
			return true
		}
		if v, exist := magic[num/t-1]; exist && index-v[0] <= k && num-v[1] <= t {
			return true
		}
		magic[num/t] = [2]int{index, num}
	}
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}
