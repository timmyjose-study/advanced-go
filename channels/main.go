package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan<- string) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		c <- "Ping"
	}
}

func ponger(c chan<- string) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		c <- "Pong"
	}
}

func printer(c <-chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Received: ", <-c)
	}
}

func main() {
	var c chan string = make(chan string)
	go func() {
		c <- "Hi!"
	}()

	msg := <-c
	fmt.Println(msg)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
