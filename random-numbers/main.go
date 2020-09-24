package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(rand.Intn(100), ", ")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, %f\n", r.Intn(100), r.Float64())
	}
}
