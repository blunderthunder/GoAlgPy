package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	hashmap := make(map[rune]int)
	longestlen := 0
	track := 0
	for idx, char := range s {
		if idt, ok := hashmap[char]; ok {
			if idt+1 > track {
				track = idt + 1
			}
		}
		hashmap[char] = idx
		if idx+1-track > longestlen {
			longestlen = idx + 1 - track
		}
	}
	return longestlen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
