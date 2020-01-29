package main

import (
	"fmt"
	"math"
)

//Defining Interface

type shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	width, height float64
}

// value reciever method(for calculation )

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}
func main() {

	Circle := circle{x: 0, y: 0, radius: 5}

	Rectangle := rectangle{width: 10, height: 5}

	fmt.Printf("Area of circle : %f\n", getArea(Circle))
	fmt.Printf("Area of rectangle : %f \n", getArea(Rectangle))
}
