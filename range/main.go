package main

import "fmt"

func main() {
	nums := []int{11, 12, 13, 14, 15}
	sum := 0

	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum = ", sum)

	for idx, n := range nums {
		fmt.Printf("Element at index %d = %d\n", idx, n)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("key = %s\n", k)
	}

	for idx, c := range "golang" {
		fmt.Printf("%d => %c\n", idx, c)
	}
}
