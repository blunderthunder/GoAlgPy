package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cache map[int][]*TreeNode

func init() {
	cache = make(map[int][]*TreeNode)
	cache[0] = []*TreeNode{}
	cache[1] = []*TreeNode{&TreeNode{}}
}

func allPossibleFBT(n int) []*TreeNode {

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	result := []*TreeNode{}
	for l := 0; l < n; l++ {
		r := n - 1 - l

		for _, lefttree := range allPossibleFBT(l) {
			for _, righttree := range allPossibleFBT(r) {
				result = append(result, &TreeNode{Val: 0, Left: lefttree, Right: righttree})
			}
		}
	}

	cache[n] = result
	return result
}

func main() {
	fmt.Println(allPossibleFBT(7))
}
