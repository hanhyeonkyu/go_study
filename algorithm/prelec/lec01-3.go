package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	var input1 int
	fmt.Println("입력할 정수의 개수 : ")
	fmt.Scanln(&input1)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input2 = scanner.Text()
	var inputArr = strings.Split(input2, " ")
	inputArr = inputArr[0:input1]
	var inputArr2 = make([]string, input1)
	copy(inputArr2, inputArr)
	var q int = 0
	for q < len(inputArr) {
		inputArr[q] = inputArr2[len(inputArr)-q-1]
		q = q + 1
	}
	fmt.Println(inputArr)
}
