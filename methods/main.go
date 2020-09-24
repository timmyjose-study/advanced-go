package main

import "fmt"

type Rect struct {
	length, width float32
}

func (r *Rect) area() float32 {
	return r.length * r.width
}

func (r *Rect) perimeter() float32 {
	return 2.0 * (r.length + r.width)
}

func main() {
	r := Rect{length: 10.0, width: 20.0}
	fmt.Printf("%v has area %f and perimeter %f\n", r, r.area(), r.perimeter())

	rp := &r
	fmt.Printf("%v has area %f and perimeter %f\n", rp, rp.area(), rp.perimeter())
}
