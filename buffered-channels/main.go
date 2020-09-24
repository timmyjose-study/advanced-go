package main

import "fmt"

func main() {
	var c chan string = make(chan string, 2)

	c <- "First"
	c <- "Second"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
