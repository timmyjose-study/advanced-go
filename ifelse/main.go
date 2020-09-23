package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Printf("%d is even\n", 7)
	} else {
		fmt.Printf("%d is odd\n", 7)
	}

	if 8%4 == 0 {
		fmt.Printf("%d is divisible by %d\n", 8, 4)
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, " has multiple digits")
	}
}
