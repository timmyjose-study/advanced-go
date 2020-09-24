package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		c2 <- "Two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}
}
