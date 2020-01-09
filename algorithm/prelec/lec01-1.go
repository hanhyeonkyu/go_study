package main

import "fmt"
import "strconv"

func main() {
	var num string
	fmt.Println("출력할 횟수를 입력하세요: ")
	fmt.Scanln(&num)
	var input, _ = strconv.Atoi(num)
	var x int = 0
	for x < input {
		fmt.Println("Hello World!")
		x = x + 1
	}
}