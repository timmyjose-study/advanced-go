package main

import "fmt"

func intSeq() func() int {
	val := 0

	return func() int {
		val++
		return val
	}
}

func main() {
	nextInt := intSeq()

	for i := 0; i < 10; i++ {
		fmt.Println(nextInt())
	}

	newInts := intSeq()
	fmt.Println(newInts())
}
