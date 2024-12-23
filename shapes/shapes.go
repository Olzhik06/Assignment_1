package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

type Triangle struct {
	SideA, SideB, SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}
func PrintShapeDetails(s Shape) {
	fmt.Println("Shape details:")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	shapes := []Shape{
		Rectangle{Length: 5, Width: 3},
		Circle{Radius: 4},
		Square{Length: 4},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}
	for _, shape := range shapes {
		PrintShapeDetails(shape)
		fmt.Println()
	}
}
