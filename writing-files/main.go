package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := []byte("Hola, mundo!\nComo estas, ese?\n")
	err := ioutil.WriteFile("/tmp/dat1", data, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	data2 := []byte{115, 111, 109, 101, 10}
	nread, err := f.Write(data2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", nread)

	const s = "This is the end, my friend!\n"
	nread1, err1 := f.WriteString(s)
	check(err1)
	fmt.Printf("Wrote %d bytes\n", nread1)

	f.Sync() // flush

	// buffered writer
	bufwrite := bufio.NewWriter(f)
	nread2, err2 := bufwrite.WriteString("My only friend, the end!\n")
	check(err2)
	fmt.Printf("Wrote %d bytes\n", nread2)
	bufwrite.Flush() // for buffered writers, use Flush
}
