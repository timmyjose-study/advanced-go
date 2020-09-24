package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

type Rect struct {
	length, width float64
}

func (r *Rect) area() float64 {
	return r.length * r.width
}

func (r *Rect) perimeter() float64 {
	return 2.0 * (r.length + r.width)
}

func area(s Shape) float64 {
	return s.area()
}

func perimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	c := Circle{x: 0.0, y: 0.0, radius: 10.0}
	fmt.Printf("Circle %v has area %f and perimeter %f\n", c, c.area(), c.perimeter())
	fmt.Println(area(&c))
	fmt.Println(perimeter(&c))

	r := Rect{length: 10.0, width: 5.0}
	fmt.Printf("Rectangle %v has area %f and perimeter %f\n", r, r.area(), r.perimeter())
	fmt.Println(area(&r))
	fmt.Println(perimeter(&r))
}
