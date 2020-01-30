package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var ftext = scanner.Text()
	scanner.Scan()
	var stext = scanner.Text()
	fbyte := []byte(ftext)
	sbyte := []byte(stext)
	std := 0
	for p := 0; p < len(fbyte); p++ {
		if fbyte[p] < sbyte[p] {
			std = -1
			break
		} else if fbyte[p] > sbyte[p] {
			std = 1
			break
		} else if fbyte[p] == sbyte[p] {
			std = 0
			continue
		}
	}
	fmt.Println(std)
}
