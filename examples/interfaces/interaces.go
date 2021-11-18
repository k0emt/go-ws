package main

import (
	"fmt"
	"math"
)

func main() {
	square := Square{10.0}
	circle := Circle{5.0}

	shapes := []Shape{square, circle}

	for _, shape := range shapes {
		fmt.Printf("%T %v %v\n", shape, shape, fmt.Sprintf("%f", shape.Area()))
		fmt.Printf("%T %f %f\n", shape, shape, shape.Area())
	}

	totalArea := sumAreas(shapes)
	fmt.Println(totalArea)
}

// ---

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

// ---

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ---

type Shape interface {
	Area() float64
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}
