package main

import "fmt"

func isEqual(a, b int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func getEqualCount(data []int, std int) int {
	var count int = 0
	for p := 0; p < len(data); p++ {
		if isEqual(data[p], std) {
			count++
		}
	}
	return count
}

func getEqualCount2p(data []int, std1 int, std2 int) int {
	var count int = 0
	for p := 0; p < len(data); p++ {
		if isEqual(data[p], std1) || isEqual(data[p], std2) {
			count++
		}
	}
	return count
}

func main() {
	// fmt.Println(isEqual(5, 5))
	// fmt.Println(getEqualCount([]int{4, 5, 4, 6, 7, 5, 6, 6}, 5))
	fmt.Println(getEqualCount2p([]int{4, 5, 6, 7, 5, 6, 6}, 4, 5))
}
