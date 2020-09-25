package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(numbers ...int) int {
	tot := 0
	for _, n := range numbers {
		tot += n
	}
	return tot
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := []int{}

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	fmt.Printf("sum = %d\n", sum(numbers...))
}
