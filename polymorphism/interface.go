// Golang program to illustrate the
// concept of interfaces

package main

import (
    "fmt"
)

// defining an interface
type Figure interface{

    Area() float64
}

// declaring a struct
type Rectangle struct{
    
    // declaring struct variables
    length float64
    width float64
}

// declaring a struct
type Square struct{
    
    // declaring struct variable
    side float64
}

// function to calculate
// area of a rectangle
func (rect Rectangle) Area() float64{

    // Area of rectangle = l * b
    area := rect.length * rect.width
    return area
}

// function to calculate
// area of a square
func (sq Square) Area() float64{
    
    // Area of square = a * a
    area := sq.side * sq.side
    return area
}

// main function
func main() {
    
    // declaring a rectangle instance
    rectangle := Rectangle{
    
        length: 10.5,
        width: 12.25,
    }
    
    // declaring a square instance
    square := Square{
    
        side: 15.0,
    }
    
    // The Figure interface can hold rectangle and square type as they both implements the interface
    var f1 Figure = rectangle
    var f2 Figure = square
      
    // printing the calculated result
    fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())
    fmt.Printf("Area of square: %.3f unit sq.\n", f2.Area())
}
