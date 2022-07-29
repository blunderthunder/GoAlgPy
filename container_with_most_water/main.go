package main

import (
	"fmt"
)

func maxArea(height []int) int {
	// check the lenght of height first
	if len(height) <= 1 {
		return 0
	} else if len(height) == 2 {
		if height[0] > height[1] {
			return height[1]
		} else {
			return height[0]
		}
	}
	startidx := 0
	endidx := len(height) - 1

	ans := 0

	for {
		if startidx >= endidx {
			break
		}
		if height[startidx] > height[endidx] {
			newans := height[endidx] * (endidx - startidx)
			if newans > ans {
				ans = newans
			}
			endidx = endidx - 1
		} else {
			newans := height[startidx] * (endidx - startidx)
			if newans > ans {
				ans = newans
			}
			startidx = startidx + 1
		}

	}

	return ans
}

func main() {

	data := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	fmt.Println(data)
}
