package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Duration(rand.Intn(250)) * time.Millisecond)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
