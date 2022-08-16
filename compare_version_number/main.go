package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// compare rivision verions in golang
	v1_map := make(map[int]int64)
	v2_map := make(map[int]int64)
	longest_length := 0

	for idx, val := range strings.Split(version1, ".") {
		v1_map[idx], _ = strconv.ParseInt(val, 10, 32)
	}
	for idx, val := range strings.Split(version2, ".") {
		v2_map[idx], _ = strconv.ParseInt(val, 10, 32)
	}

	if len(v1_map) > len(v2_map) {
		longest_length = len(v1_map)
	} else {
		longest_length = len(v2_map)
	}
	for idx := 0; idx < longest_length; idx++ {
		_, ok := v1_map[idx]
		if !ok {
			v1_map[idx] = 0
		}
		_, ok = v2_map[idx]
		if !ok {
			v2_map[idx] = 0
		}
		val1 := v1_map[idx]
		val2 := v2_map[idx]

		if val1 > val2 {
			return 1
		} else if val2 > val1 {
			return -1
		} else {
			continue
		}
	}

	return 0
}

func main() {
	version1 := "0.1"
	version2 := "1.01"

	result := compareVersion(version1, version2)

	fmt.Println("compare versions of {} and {} is {}", version1, version2, result)
}
