package main

import (
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
	err := os.Mkdir("subdir", 0755)
	check(err)
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(ioutil.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")
	err1 := os.MkdirAll("subdir/parent/child", 0755)
	check(err1)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err2 := ioutil.ReadDir("subdir/parent")
	check(err2)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err3 := os.Chdir("subdir/parent/child")
	check(err3)

	c1, err4 := ioutil.ReadDir(".")
	check(err4)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c1 {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
}
