package main

import "fmt"
import "os"
import "math"
import "bufio"
import "strconv"
import "strings"

func getMinIndex(data []int) int {
	var index int = math.MaxInt32
	var lowest int = math.MaxInt32
	var auxindex bool
	if len(data) == 1 {
		return 0
	}
	for p := 0; p < len(data)-1; p++ {
		var lower int
		if data[p] <= data[p+1] {
			lower = data[p]
			auxindex = false
		} else if data[p] > data[p+1] {
			lower = data[p+1]
			auxindex = true
		}
		if lowest >= lower {
			if auxindex == false {
				index = p
				lowest = lower
			} else if auxindex == true {
				index = p + 1
				lowest = lower
			}
		}
	}
	return index
}

func selectionSort(data []int, begin int, end int) []int {
	var copyData []int = make([]int, len(data))
	copy(copyData, data)
	var temp []int = copyData[begin : end+1]
	for p := 0; p < len(temp); p++ {
		var origin int = temp[p]
		var minIndex int = getMinIndex(temp[p:])
		temp[p] = temp[minIndex+p]
		temp[minIndex+p] = origin
	}
	for r := begin; r < end+1; r++ {
		copyData[r] = temp[r-begin]
	}
	return copyData
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var prenum string = scanner.Text()
	var num, _ = strconv.Atoi(prenum)
	var numArr = make([]int, num)
	scanner.Scan()
	var input string = scanner.Text()
	inputArr := strings.Split(input, " ")
	for p := 0; p < len(inputArr); p++ {
		var tmp, _ = strconv.Atoi(inputArr[p])
		numArr[p] = tmp
	}

	var sortArr = selectionSort(numArr, 0, len(numArr)-1)
	var std bool = false
	for p := 0; p < num; p++ {
		if numArr[p] == sortArr[p] {
			std = true
		} else {
			std = false
			break
		}
	}
	if std {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
