package main

import "fmt"

func sumEachSum(x int) int {
	var finalSum int = 0;
	for p := 1; p < x + 1; p++ {
 		var eachSum int = 0;
		for q := 1; q < p + 1; q++ {
			eachSum += q
		}
		finalSum += eachSum
	}
	return finalSum
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(sumEachSum(num))
}