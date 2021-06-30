package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, num := range nums {
		fmt.Println(num, i, hashmap)
		d := target - num
		index, ok := hashmap[d]
		if ok {
			return []int{index, i}
		} else {
			hashmap[num] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
