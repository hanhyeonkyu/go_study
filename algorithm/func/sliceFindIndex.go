package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func main() {
	var slicelen int
	var targetnum int
	fmt.Scanln(&slicelen, &targetnum)
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var arraytext = scanner.Text()
	var inputArr = strings.Split(arraytext, " ")
	inputArr = inputArr[:slicelen]
	var targetnumtemp string = strconv.Itoa(targetnum)
	var index int = -1
	for p := 0; p < len(inputArr); p++ {
		if inputArr[p] == targetnumtemp {
			index = p
		}
	}
	fmt.Println(index)
	
}