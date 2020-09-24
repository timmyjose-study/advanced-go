package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		fmt.Println(msg)
	}("Hola")

	go func(msg string) {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		fmt.Println(msg)
	}("Hi")

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
