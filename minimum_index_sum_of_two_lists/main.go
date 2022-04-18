package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	res := []string{}
	minimum := 2001

	for i, s := range list1 {
		m[s] = i
	}

	for i, s := range list2 {
		if _, ok := m[s]; ok {
			indexSum := m[s] + i
			if indexSum < minimum {
				minimum = indexSum
				res = []string{s}
			} else if indexSum == minimum {
				res = append(res, s)
			}
		}
	}
	return res
}

func main() {
	v1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	v2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	fmt.Println("minimum index sum of given list ", v1, v2, "is : ", findRestaurant(v1, v2))
}
