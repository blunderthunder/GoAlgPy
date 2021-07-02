package main

import (
	"fmt"
	"math"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	var max_profit, min_price int = 0, math.MaxInt32

	for _, price := range prices {
		min_price = min(price, min_price)
		profit := price - min_price
		max_profit = max(profit, max_profit)
	}
	return max_profit
}

func main() {
	fmt.Println(maxProfit([]int{5, 4, 3, 2, 1}))
}
