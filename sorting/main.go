package main

import (
	"fmt"
	"sort"
)

func main() {
	// since Go does not yet have Generics, there are different sorting functions for
	// differen types
	strings := []string{"hello", "world", "how", "are", "you?"}
	fmt.Println("Before sorting, strings = ", strings)
	sort.Strings(strings)
	fmt.Println("After sorting, strings = ", strings)
	fmt.Println(sort.StringsAreSorted(strings))

	numbers := []int{11, -11, 99, 1, 2, 3, 3, 2, 1, 0, 11, 99, 199}
	fmt.Println("Before sorting, numbers = ", numbers)
	sort.Ints(numbers)
	fmt.Println("After sorting, strings = ", numbers)
	fmt.Println(sort.IntsAreSorted(numbers))
}
