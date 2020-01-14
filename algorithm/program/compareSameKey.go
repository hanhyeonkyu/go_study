package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	fmt.Scanln(&num)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var heights string = scanner.Text()
	var heightArr = strings.Split(heights, " ")
	heightArr = heightArr[:num]
	scanner.Scan()
	var months string = scanner.Text()
	var monthArr = strings.Split(months, " ")
	monthArr = monthArr[:num]
	var targetMonth string
	fmt.Scanln(&targetMonth)
	var index int = -1
	for p := num - 1; p >= 0; p-- {
		if monthArr[p] == targetMonth {
			index = p
			break
		}
	}

	if index == -1 {
		fmt.Println(index)
	} else {
		fmt.Println(heightArr[index])
	}
}
