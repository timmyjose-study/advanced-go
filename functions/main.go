package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return plus(a, b) + c
}

func main() {
	a, b, c := 1, 2, 3
	fmt.Println(plus(a, b))
	fmt.Println(plusPlus(a, b, c))
}
