package main

import (
	"fmt"
)

// shape
type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Printf("Area: %v\n", s.getArea())
}

// triangle
type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// square
type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// ----
func shapes() {
	t := triangle{base: 2.0, height: 3.0}
	s := square{sideLength: 3.0}

	printArea(t)
	printArea(s)
}
