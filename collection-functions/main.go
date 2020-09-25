package main

import (
	"fmt"
	"strings"
)

// Index finds the index of the first match of v in vs
// else return -1
func Index(vs []string, v string) int {
	for idx, s := range vs {
		if s == v {
			return idx
		}
	}
	return -1
}

// Include checks if v is present in vs
func Include(vs []string, v string) bool {
	return Index(vs, v) >= 0
}

// Any checks if any of vs satisfies the predicate p
func Any(vs []string, p func(string) bool) bool {
	for _, v := range vs {
		if p(v) {
			return true
		}
	}
	return false
}

// All checks if all elements of vs satisfy the predicate p
func All(vs []string, p func(string) bool) bool {
	count := 0
	for _, v := range vs {
		if p(v) {
			count++
		}
	}

	return count == len(vs)
}

// Filter returns a new string slice with all elements of vs that
// satisfy the predicate p
func Filter(vs []string, p func(string) bool) []string {
	res := make([]string, 0)
	for _, v := range vs {
		if p(v) {
			res = append(res, v)
		}
	}

	return res
}

// Map returns a new string slice with the function f applied
// to each element of vs
func Map(vs []string, f func(string) string) []string {
	res := make([]string, 0)
	for _, v := range vs {
		res = append(res, f(v))
	}

	return res
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}
