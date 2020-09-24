package main

import "fmt"

func sum(nums ...int) {
	fmt.Println("nums: ", nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{100, 200, 300, 400, 500}
	sum(nums...)
}
