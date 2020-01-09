package main

import (
	"fmt"
	"time"
)

func main() {
	selectDefault()
}

func selectDefault() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

func selectfunc() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func channelClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func channels() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// 채널을 몇개 이용할지에 대해 버퍼되는 개수를 설정할 수 있다.
	// c := make(chan int, 100)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(c)
	fmt.Println(x, y, x+y)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func goroutine() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}