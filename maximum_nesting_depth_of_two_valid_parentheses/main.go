package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	depth := 0
	ans := []int{}
	for _, val := range seq {
		if val == '(' {
			ans = append(ans, depth%2)
			depth += 1
		} else {
			depth -= 1
			ans = append(ans, depth%2)
		}
	}
	return ans
}

func main() {
	seq := "(()())"
	expected_result := []int{0, 1, 1, 1, 1, 0}

	fmt.Println("Output of input sequence ", seq, maxDepthAfterSplit(seq))
	fmt.Println("expected return is ", expected_result)
}
