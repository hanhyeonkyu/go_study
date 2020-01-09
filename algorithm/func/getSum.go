package main

import "fmt"

func getSumInt(std int) int {
	if std <= 0 {
		return 0
	}
	var sum int = 0
	for p := 1; p < std+1; p++ {
		sum += p
	}
	return sum
}

func getSumArr(data []int) int {
	var sum int = 0
	for p := 0; p < len(data); p++ {
		sum += data[p]
	}
	return sum
}

func main() {
	// fmt.Println(getSumArr([]int{4, 5, 6, 23, 54}))
	fmt.Println(getSumInt(1100))
}
