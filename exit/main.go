package main

import (
	"fmt"
	"os"
)

func main() {
	// defers are not run when explicitly quitting using os.Exit
	defer fmt.Println("This line will not be printed")
	os.Exit(3)
}
