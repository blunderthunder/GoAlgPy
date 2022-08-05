package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	// sort the array
	sort.Ints(nums)

	result := [][]int{}

	if n < 3 {
		return [][]int{}
	} else if n == 3 {
		// check if they are triplet or not
		i, j, k := nums[0], nums[1], nums[2]
		if (nums[0] + nums[1] + nums[2]) == 0 {
			result = append(result, []int{i, j, k})
			return result
		}
	}

	// if all the condiditon fail  use pointer method
	// define 3 pointer
	for num1Idx := 0; num1Idx < (n - 2); num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Triplets from array ", nums, "is ", threeSum(nums))
}
