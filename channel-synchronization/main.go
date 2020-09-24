package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(done chan<- bool) {
	fmt.Println("working...")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("Done")
	done <- true
}

func main() {
	var c chan bool = make(chan bool, 1)
	go worker(c)
	<-c
}
