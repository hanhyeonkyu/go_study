package main

import "fmt"
import "sort"

func getSumArr(data []int) int {
	var sum int = 0
	for p := 0; p < len(data); p++ {
		sum += data[p]
	}
	return sum
}

func getPos(sum, std int) bool {
	if sum >= std {
		return false
	} else {
		return true
	}
}

func getPepPos(data []int, std1 int, std2 int) pepInfo {
	var count int = 0
	var weight []int
	for p := 0; p < len(data); p++ {
		if data[p] <= 6 {
			count++
			weight = append(weight, data[p])
		}
	}
	var weightSum = getSumArr(weight)
	var isPossible = getPos(weightSum, std2)
	sort.Slice(weight, func(i, j int) bool {
		return weight[i] < weight[j]
	})
	var posWeightSum int = 0
	var posArr []int
	var posCount int = 0
	for q := 0; q < len(weight); q++ {
		posWeightSum += weight[q]
		if posWeightSum >= std2 {
			posWeightSum -= weight[q]
			break
		}
		posArr = append(posArr, weight[q])
		posCount++
	}

	return pepInfo{
		weight:       weight,
		count:        count,
		weightSum:    weightSum,
		isPossible:   isPossible,
		posArr:       posArr,
		posCount:     posCount,
		posWeightSum: posWeightSum,
	}
}

// 놀이기구를 타기 위해 사람들의 몸무게에 따라 허용 가능한 수준과 허용 불가능한 수준을 나누고 허용이 불가능할 경우
// 낮은 무게 순으로 최대한 탈 수 있는 조합을 만든다.
func main() {
	var result = getPepPos([]int{4, 5, 6, 8, 3, 4, 6, 7, 8, 10}, 6, 15)
	fmt.Println(result)
}

type pepInfo struct {
	weight       []int // 허용 가능한 몸무게를 가진 사람들의 슬라이스
	count        int   // 허용 가능한 몸무게를 가진 사람들의 수
	weightSum    int   // 허용 가능한 몸무게를 가진 사람들의 몸무게의 총합
	isPossible   bool  // 총 허용 가능한 무게에 대해 weightSum 을 허용할 수 있는지에 대한 bool 값
	posArr       []int // isPossible 을 만족하는 몸무게가 낮은 사람들의 배열
	posCount     int   // isPossible 을 만족하는 몸무게가 낮은 사람들의 수
	posWeightSum int   // isPossible 을 만족하는 몸무게가 낮은 사람들의 몸무게의 총합
}
