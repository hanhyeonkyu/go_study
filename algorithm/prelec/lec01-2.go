package main

import "fmt"
import "strconv"

func main() {
	var z int = 0
	for z < 2 {
		var input1 string
		var input2 string
		fmt.Println("insert number : ")
		fmt.Scanln(&input1, &input2)
		var x, _ = strconv.Atoi(input1)
		var y, _ = strconv.Atoi(input2)
		var result = x + y
		fmt.Println(result)
		z = z + 1
	}
}