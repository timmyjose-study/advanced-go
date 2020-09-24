package main

import "os"

func main() {
	panic("something went wrong")

	_, err := os.Create("/tmp/foo.txt")
	if err != nil {
		panic(err)
	}
}
