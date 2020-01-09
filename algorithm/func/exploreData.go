package main

import (
	"fmt"
	"math"
)

func getSumArr(data []int) int {
	var sum int = 0
	for p := 0; p < len(data); p++ {
		sum += data[p]
	}
	return sum
}

func getIndexByVal(data []int, std int) []int {
	var indexArr []int
	for p := 0; p < len(data); p++ {
		if data[p] == std {
			indexArr = append(indexArr, p)
		}
	}
	return indexArr
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
	// fmt.Println(getIndexByVal([]int{3, 5, 6, 6, 5, 4, 4, 2, 4}, 4))
	// fmt.Println(getMinIndex([]int{3, 5, 6, 6, 5, 4, 4, 2, 4}))
	// fmt.Println(selectionSort([]int{6, 5, 4, 3, 2, 1, 8, 9, 0}, 3, 7))
	fmt.Println(selectionSort([]int{3, 5, 6, 6, 5, 4, 4, 2, 4}, 3, 6))
}
