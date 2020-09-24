package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) Person {
	return Person{name: name, age: 42}
}

func main() {
	fmt.Println(newPerson("Bob"))
	fmt.Println(Person{name: "Dave", age: 19})

	david := Person{name: "David", age: 0}
	fmt.Println(david)

	pp := &david
	pp.age = 21
	fmt.Println(david)
}
