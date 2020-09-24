package main

import "fmt"

func zeroval(x int) {
	x = 0
}

func zeroptr(px *int) {
	*px = 0
}

func main() {
	val := 100
	fmt.Printf("val = %d\n", val)
	zeroval(val)
	fmt.Printf("val = %d\n", val)
	zeroptr(&val)
	fmt.Printf("val = %d\n", val)

	fmt.Println(&val)
}
