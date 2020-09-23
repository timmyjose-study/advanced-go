package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c1 <- "from 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c2 <- "from 2"
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case t := <-time.After(time.Second):
				fmt.Printf("Timeout at %v!", t)
				return
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
