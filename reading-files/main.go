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
	// slurp the entire contents of the file
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat)) // data are raw bytes

	// get a file handl
	f, err := os.Open("/tmp/dat")
	defer f.Close()

	check(err)
	buf := make([]byte, 5)
	nread, err1 := f.Read(buf)
	check(err1)
	fmt.Printf("%d bytes: %s\n", nread, string(buf[:nread]))
	nread, err2 := f.Read(buf)
	check(err2)
	fmt.Printf("Another %d bytes: %s\n", nread, string(buf[:nread]))

	beg, err := f.Seek(0, 0)
	check(err)
	buf2 := make([]byte, 16)
	nread, err3 := f.Read(buf2)
	check(err3)
	fmt.Printf("%d bytes read from %d: %s\n", nread, beg, string(buf2[:nread]))
	f.Seek(0, 0)

	// buffered reader
	bufread := bufio.NewReader(f)
	buf3, err := bufread.Peek(1024)
	fmt.Printf("%d bytes read (buffered): %s\n", 1024, string(buf3))
}
