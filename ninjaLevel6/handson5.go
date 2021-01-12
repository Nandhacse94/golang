package main

import (
	"fmt"
	"math"
)

// create two structs - circle, square
// create a interface declaring the area func.
// define func area with two different receivers (circle & square)
// create a func with interface type parameter and call the area func

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {

	//create a value and call the method
	a := circle{
		radius: 10,
	}

	fmt.Println("circle radius: ", a)

	b := square{
		length: 50,
	}

	fmt.Println("square length: ", b)

	info(a)
	info(b)
}
