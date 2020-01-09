package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
	var numcase int
	fmt.Println("입력할 case 의 개수 : ")
	fmt.Scanln(&numcase)
	var caseArr = make([]int, numcase)
	var q int = 0
	for q < len(caseArr) {
		var scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		var input = scanner.Text()
		var inputArr = strings.Split(input, " ")
		var r int = 0
		var sum int = 0
		for r < len(inputArr) {
			var t, _ = strconv.Atoi(inputArr[r])
			sum = sum + t
			r = r + 1
		}
		caseArr[q] = sum
		q = q + 1
	}
	var s int = 0
	for s < len(caseArr) {
		fmt.Printf("Case #%d", s+1)
		fmt.Println()
		fmt.Println(caseArr[s])
		s = s + 1
	}
}
