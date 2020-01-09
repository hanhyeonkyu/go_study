package main

import "fmt"

func eachSlices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func sliceSlices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("p == %d\np[1:4] == %d\np[:3] == %d\np[4:] == %d", p, p[1:4], p[:3], p[4:])
}

func makeSlices() {
	a := make([]int, 5)
	b := make([]int, 0, 5)
	c := b[:2]
	d := c[2:5]
	printSlice("a", a)
	printSlice("b", b)
	printSlice("c", c)
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func emptySlices() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func calpow() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i ,v)
	}
}

func rangeUse() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func map1() {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
func map2() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": Vertex{37.42202, -122.08408},
	}
	fmt.Println(m)
}

func map3() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func map4() {
	m := make(map[string]int)
	m["kasy"] = 24
	fmt.Println("kasy: ", m["kasy"])
	delete(m, "kasy")
	fmt.Println("after delete kasy : ", m["kasy"])
	v, ok := m["kasy"]
	fmt.Println("kasy : ", v, "is?", ok)
}

func main() {
	map4()
}

type Vertex struct {
	Lat, Long float64
}