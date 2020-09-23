package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n)
	}
}

func main() {
	go f(0)
	var input int
	fmt.Scanln(&input)
}
