package main

import "fmt"
import "strings"

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func main() {
	fmt.Println(arrayToString([]int{3, 4, 5, 6, 8}, " "))
}
