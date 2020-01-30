package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var prenum = scanner.Text()
	var num, _ = strconv.Atoi(prenum)
	var numArr = make([]int, num)
	scanner.Scan()
	var input string = scanner.Text()
	inputArr := strings.Split(input, " ")
	for p := 0; p < len(inputArr); p++ {
		var tmp, _ = strconv.Atoi(inputArr[p])
		numArr[p] = tmp
	}
	var count int = 0
	for q := 0; q < len(numArr); q++ {
		if q == len(numArr)-1 {
			count += 1
			break
		}
		if numArr[q] < numArr[q+1] {
			count += 1
		} else if numArr[q] >= numArr[q+1] {
			continue
		}
	}
	fmt.Println(count)
}
