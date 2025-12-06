package main

import "fmt"

// Figure demonstrates implicit interface satisfaction.
type Figure interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 { return r.Length * r.Width }

type Square struct {
	Side float64
}

func (s Square) Area() float64 { return s.Side * s.Side }

func main() {
	// Interface values can store any type that implements the methods.
	var f1 Figure = Rectangle{Length: 10.5, Width: 12.25}
	var f2 Figure = Square{Side: 15}

	fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())
	fmt.Printf("Area of square: %.3f unit sq.\n", f2.Area())
}
