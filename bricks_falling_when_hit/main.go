package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func valid(grid [][]int, node [2]int) bool {
	if node[0] < 0 || node[0] >= len(grid) || node[1] < 0 || node[1] >= len(grid[0]) {
		return false
	}
	if grid[node[0]][node[1]] <= 0 {
		return false
	}
	return true
}

func add(father map[[2]int][2]int, node [2]int, sizeofSet map[[2]int]int) {
	if _, found := father[node]; !found {
		father[node] = [2]int{-2, -2}
		sizeofSet[node] = 1
	}

}

func find(father map[[2]int][2]int, node [2]int) [2]int {
	root := node
	for father[root] != [2]int{-2, -2} {
		root = father[root]
	}

	for node != root {
		origFather := father[node]
		father[node] = root
		node = origFather
	}

	return root
}

func merge(father map[[2]int][2]int, a, b [2]int, sizeofSet map[[2]int]int) {
	rootA := find(father, a)
	rootB := find(father, b)
	if rootA != rootB {
		father[rootA] = rootB
		sizeofSet[rootB] += sizeofSet[rootA]
	}
	return
}

func hitBricks(grid [][]int, hits [][]int) []int {
	WALL := [2]int{-1, -1}
	connections := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	father := make(map[[2]int][2]int)
	sizeofSet := make(map[[2]int]int)

	//remove all hits bricks
	for _, hit := range hits {
		grid[hit[0]][hit[1]]--
	}
	//add WALL into disjoint set
	add(father, WALL, sizeofSet)

	//add all bricks into disjoint set
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				add(father, [2]int{i, j}, sizeofSet)
			}
		}
	}

	//merge row 0 bricks to WALL
	for j := 0; j < len(grid[0]); j++ {
		if grid[0][j] == 1 {
			merge(father, [2]int{0, j}, WALL, sizeofSet)
		}
	}

	//process row 1 to n, merge connections
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				for _, connection := range connections {
					if valid(grid, [2]int{i + connection[0], j + connection[1]}) {
						merge(father, [2]int{i, j}, [2]int{i + connection[0], j + connection[1]}, sizeofSet)
					}
				}
			}
		}
	}
	// fmt.Println("father    = ", father)
	// fmt.Println("sizeofSet = ", sizeofSet)
	res := make([]int, 0)

	//process hits backwards, merge the hit brick based on connections
	for i := len(hits) - 1; i >= 0; i-- {
		before := sizeofSet[find(father, WALL)]
		grid[hits[i][0]][hits[i][1]]++
		if grid[hits[i][0]][hits[i][1]] == 1 {
			//add the hit back to fathers
			add(father, [2]int{hits[i][0], hits[i][1]}, sizeofSet)
			for _, connection := range connections {
				if valid(grid, [2]int{hits[i][0] + connection[0], hits[i][1] + connection[1]}) {
					merge(father, [2]int{hits[i][0], hits[i][1]}, [2]int{hits[i][0] + connection[0], hits[i][1] + connection[1]}, sizeofSet)
				}
			}
			//if the hit is on row 0, remember merge it to the wall
			if hits[i][0] == 0 {
				merge(father, [2]int{hits[i][0], hits[i][1]}, WALL, sizeofSet)
			}
		}
		after := sizeofSet[find(father, WALL)]
		diff := max(0, after-before-1)
		res = append(res, diff)
	}
	result := make([]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
	}
	return result
}

func main() {
	grid := [][]int{[]int{1, 0, 0, 0}, []int{1, 1, 1, 0}}
	hit := [][]int{[]int{1, 0}}
	result := hitBricks(grid, hit)
	fmt.Println("result is :", result)
}
