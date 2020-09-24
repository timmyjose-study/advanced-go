package main

import "fmt"

func vals() (int, int) {
	return 42, 7
}

func main() {
	a, b := vals()
	fmt.Printf("a = %d, b = %d\n", a, b)

	aa, _ := vals()
	fmt.Printf("aa = %d\n", aa)

	_, bb := vals()
	fmt.Printf("bb = %d\n", bb)
}
