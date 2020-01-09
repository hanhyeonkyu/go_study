package main

import "fmt"
import "math"

func getMax(a, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else if a == b {
		return a
	} else {
		return a
	}
}

func getMaxArr(data []int) int {
	var best int = math.MinInt32
	for i := 0; i < len(data)-1; i++ {
		var upper = getMax(data[i], data[i+1])
		best = getMax(upper, best)
	}
	return best
}

func main() {
	// fmt.Println(getMax(10, 5))
	fmt.Println(getMaxArr([]int{-546, 0, 6, 20, 4, 2000}))
}
