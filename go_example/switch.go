package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is a weekend!")
	default:
		fmt.Println("it is a weekday...")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am bool")
		case int:
			fmt.Println("i am int")
		default:
			fmt.Println("don't know what type is", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
