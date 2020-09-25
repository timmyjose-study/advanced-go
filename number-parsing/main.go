package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Printf("f = %f\n", f)

	i, _ := strconv.ParseInt("12345", 0, 32)
	fmt.Printf("i = %d\n", i)

	h, _ := strconv.ParseInt("0xdeadbeef", 0, 32)
	fmt.Printf("h = %x\n", h)

	u, _ := strconv.ParseUint("9292", 0, 64)
	fmt.Printf("u = %d\n", u)

	a, _ := strconv.Atoi("12123")
	fmt.Printf("a = %d\n", a)

	_, e := strconv.Atoi("abc123")
	fmt.Println(e)
}
