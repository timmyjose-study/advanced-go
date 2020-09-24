package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "one"
	messages <- "two"

	close(messages) // this is fine we have buffering

	for m := range messages {
		fmt.Println(m)
	}
}
