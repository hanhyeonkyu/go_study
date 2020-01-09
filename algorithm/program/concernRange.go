package main

import "fmt"

func main() {
	const std string = "AJOU"
	var num int
	fmt.Scanln(&num)
	var schoolArray = make([]string, num)
	for p := 0; p < num; p++ {
		var tmp string
		fmt.Scanln(&tmp)
		schoolArray[p] = tmp
	}

	var indexf int = -1
	var indexl int = -1
	for q := 0; q < len(schoolArray); q++ {
		if schoolArray[q] == "AJOU" {
			indexf = q + 1
			break
		}
	}
	for q := len(schoolArray) - 1; q >= 0; q-- {
		if schoolArray[q] == "AJOU" {
			indexl = q + 1
			break
		}
	}
	fmt.Println(indexf, indexl)
}
