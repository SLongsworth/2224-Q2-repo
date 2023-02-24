package main

import (
	"fmt"
	"math"
)

//Demonstrate Interfaces

type geometricShapes interface {
	area() float64
	perimeter() float64
	figure() string
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) figure() string {
	return "circular"
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return (4) * s.side
}
func (s Square) figure() string {
	return "boxed"
}

//Create function to print info
func PrintShape(gs geometricShapes) {
	fmt.Println("This Geometric Shape is", gs.figure(), "has an area of", gs.area(), "and has a perimeter of", gs.perimeter())
}

func main() {
	//Create circle
	circ := Circle{
		radius: 4,
	}
	//Create square
	sqre := Square{
		side: 10,
	}
	PrintShape(circ)
	PrintShape(sqre)

}
