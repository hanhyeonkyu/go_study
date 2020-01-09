package main

import "fmt"
import "math"
import "math/cmplx"
import "runtime"
import "time"

func main() {
	switches3()
}

func switches1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s.", os)
	}
}

func switches2() {
	fmt.Println("when's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func switches3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func ifElseFunc() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3,2,10), pow(3,3,20))
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, u, u)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// math.Sqrt(x) => square-root x
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// math.Pow(x, n) = x^n
	if v:= math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func forLoop() {
	// for 문
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// while 문
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// 무한루프 문
	for {
		if(sum <= 10) {
			break
		}
		sum -= 3
	}
	fmt.Println(sum)
}

func doubleOperatorFunc() {
	// (2^5 * 1) * 2^-3
	var doubleOperator = (1<<5) >> 3
	fmt.Println(doubleOperator)
}

func variableFunc() {
	fmt.Println(add(43, 29))
	a, b := swap("hi", "there")
	fmt.Println(swap("world", "hello"))
	fmt.Println(b, a)
	fmt.Println(split(27))
	fmt.Println(x,y,z,c,python,java)
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"
	fmt.Println(x,y,z,c,python,java)
	const World = "hi"
	fmt.Println("hello", World)
	fmt.Printf("happy %g", Pi)
	const Truth = true
	fmt.Println("Go Go", Truth)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

var x, y, z int
var c, python, java bool

const Pi float64= math.Pi

const (
	Big = 1 << 100
	Small = Big >> 99
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	u complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

func needInt(x int) int { return x * 10 + 1}
func needFloat(x float64) float64 { return x * 0.1}