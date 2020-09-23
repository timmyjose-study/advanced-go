package main

import (
	"fmt"
	"math"
)

const s string = "constants"

func main() {
	fmt.Println(s)

	const n = 5000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(math.Sin(n))
}
