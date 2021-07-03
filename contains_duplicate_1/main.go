package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	hashmap := make(map[int]int)
	for _, num := range nums {
		if _, ok := hashmap[num]; ok {
			return true
		} else {
			hashmap[num] = 1
		}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
