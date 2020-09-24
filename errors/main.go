package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42!")
	}

	return arg + 10, nil
}

// custom error type
type ArgError struct {
	arg     int
	problem string
}

// we just need to provide the Error() method which has the signature:
// func Error() string
func (e *ArgError) Error() string {
	return fmt.Sprintf("%d -- %s\n", e.arg, e.problem)
}

type FooError struct {
	arg     int
	problem string
}

func (e *FooError) Error() string {
	return fmt.Sprintf("%d <--> %s", e.arg, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "Can't work with it"}
	}
	return arg + 200, nil
}

func f3(arg int) (int, error) {
	if arg == 96 {
		return -1, &FooError{arg, "Cannot work with this!"}
	}
	return arg + 300, nil
}

func main() {
	for _, n := range []int{7, 42} {
		if r, e := f1(n); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, m := range []int{7, 42} {
		if r, e := f2(m); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*ArgError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.problem)
	}

	// this will not match
	_, err := f3(96)
	if fe, ok := err.(*ArgError); ok {
		fmt.Println(fe.arg)
		fmt.Println(fe.problem)
	}

	// while this will since the types match up
	// The Poor Man's pattern matching!
	_, err1 := f3(96)
	if fe, stat := err1.(*FooError); stat {
		fmt.Println(fe.arg)
		fmt.Println(fe.problem)
	}
}
