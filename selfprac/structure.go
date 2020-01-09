package main

import "fmt"
// import "math"
import "time"
import "os"

func main() {
	errorFunc()
}

func errorFunc() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func ReadWriterFunc() {
	var w Writer
	w = os.Stdout
	fmt.Println(w, "Hello, Writer\n")
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// func interfaceFunc() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}
// 	a = f
// 	a = &v
// 	a = v
// 	fmt.Println(a.Abs())
// }

// type Abser interface {
// 	Abs() float64
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
//     if f < 0 {
//         return float64(-f)
//     }
//     return float64(f)
// }

// type Vertex struct {
//     X, Y float64
// }

// func (v *Vertex) Abs() float64 {
//     return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func typePrint() {
// 	v := Vertex{2, 4}
// 	v.X = 4
// 	fmt.Println(v)
// 	p := Vertex{4,7}
// 	q := &p
// 	q.X = 1e3
// 	fmt.Println(p)
// 	w := new(Vertex)
// 	fmt.Println(w)
// 	w.X, w.Y = 11, 9
// 	fmt.Println(w)
// }

// var (
//     p = Vertex{1, 2}  // has type Vertex
//     q = &Vertex{1, 2} // has type *Vertex
//     r = Vertex{X: 1}  // Y:0 is implicit
//     s = Vertex{}      // X:0 and Y:0
// )

// type Vertex struct {
// 	X int
// 	Y int
// }