package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)
	fmt.Printf("Temp file's name: %s\n", f.Name())

	defer os.Remove(f.Name())

	_, err1 := f.Write([]byte{1, 2, 3, 4, 5})
	check(err1)

	dname, err2 := ioutil.TempDir("", "sampledir")
	check(err2)
	fmt.Printf("Temp dir's name: %s\n", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
