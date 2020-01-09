package main

import "fmt"
import "math"

func main() {
	method3()
}

func method3() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func method2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X * v.X + v.Y * v.Y)
// }

// func method1() {
// 	v := &Vertex{3,4}
// 	fmt.Println(v.Abs())
// }

func funcCloser() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func praFunc(a, b float64) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(a, b))
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64