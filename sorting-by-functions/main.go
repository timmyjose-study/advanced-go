package main

import (
	"fmt"
	"sort"
)

// to sort using a custom type, we have to define the type and also
// make it implement the sort.Interface interface which defines three
// methods - Len, Swap, and Less

type byLength []string
type byReverseLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byReverseLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byReverseLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return (len(s[i]) < len(s[j]))
}

func (s byReverseLength) Less(i, j int) bool {
	return (len(s[i]) > len(s[j]))
}

func main() {
	fruits := []string{"apple", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	sort.Sort(byReverseLength(fruits))
	fmt.Println(fruits)
}
