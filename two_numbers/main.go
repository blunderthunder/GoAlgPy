package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(arr *[]int) *ListNode {
	l := &ListNode{}
	curr := &ListNode{}
	l.Next = curr

	for _, val := range *arr {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = val
	}
	return l.Next.Next
}

func getGraph(l *ListNode) string {
	graph := ""
	for l != nil {
		graph = graph + fmt.Sprintf("%s ( %T ) ->", strconv.Itoa(l.Val), *l)
		l = l.Next
	}
	return graph
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	root := &ListNode{}
	n := &ListNode{}
	root.Next = n

	for l1 != nil || l2 != nil || carry != 0 {
		n.Next = &ListNode{}
		n = n.Next
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		val := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		n.Val = val
		// n.Next = &ListNode{Val: val, Next: n.Next}
	}

	return root.Next.Next
}
func main() {
	l1 := makeListNode(&[]int{2, 4, 3})
	l2 := makeListNode(&[]int{5, 6, 4})

	fmt.Println(getGraph(l1))
	fmt.Println(getGraph(l2))
	fmt.Println(getGraph(addTwoNumbers(l1, l2)))
}
