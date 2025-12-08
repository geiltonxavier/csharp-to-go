package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printAreas(shapes []Shape) {
	for _, s := range shapes {
		fmt.Printf("area: %.2f\n", s.Area())
	}
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
	}
	printAreas(shapes)
}
