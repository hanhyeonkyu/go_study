package main

import "fmt"
import "math"
import "bufio"
import "os"
import "strings"
import "strconv"

func getSumArr(data []int) int {
	var sum int = 0
	for p := 0; p < len(data); p++ {
		sum += data[p]
	}
	return sum
}

func getNearAvgIndex(data []int) int {
	var index int = 0
	var best int = math.MaxInt32
	var avg = getSumArr(data) / len(data)
	for p := 0; p < len(data); p++ {
		var temp int
		if (data[p] - avg) >= 0 {
			temp = data[p] - avg
		} else {
			temp = avg - data[p]
		}
		if temp <= best {
			best = temp
			index = p
		}
	}
	return index
}

func main() {
	var num int
	fmt.Scanln(&num)
	var numArr = make([]int, num)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input string = scanner.Text()
	var inputArr = strings.Split(input, " ")
	for p := 0; p < len(inputArr); p++ {
		var tmp, _ = strconv.Atoi(inputArr[p])
		numArr[p] = tmp
	}
	var index = getNearAvgIndex(numArr)
	var value = numArr[index]
	fmt.Println(index + 1, value)
}