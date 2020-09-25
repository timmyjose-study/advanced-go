package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "HOLA")
	fmt.Printf("FOO = %s\n", os.Getenv("FOO"))
	fmt.Printf("BAR = %s\n", os.Getenv("BAR"))

	// list all environment variables
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s has the value %s\n", pair[0], pair[1])
	}
}
