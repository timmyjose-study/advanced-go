package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, ", as ")
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
		fmt.Println("Time for fun!")
	default:
		fmt.Println("Time for work!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon!")
	default:
		fmt.Println("It's after noon!")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I'm a Bool!")
		case int:
			fmt.Println("I'm an int!")
		default:
			fmt.Println("I don't know who I am!")
		}
	}

	whatAmI(100)
	whatAmI(false)
	whatAmI("hello")
}
