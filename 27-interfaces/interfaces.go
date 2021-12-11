package main

import (
	"fmt"
	"math"
)

type geometric interface {
    area() float64
}

func writeArea(shape geometric) {
    fmt.Printf("The area of shape is %0.2f\n", shape.area())
}

type rectangle struct {
    height float64
    width float64
}

func (rectangle rectangle) area() float64 {
    return rectangle.height * rectangle.width
}

type circle struct {
    radius float64
}

func (circle circle) area() float64 {
    return math.Pi * (circle.radius * circle.radius)
}

func main() {
    simpleRectangle := rectangle{10, 15}
    writeArea(simpleRectangle)

    simpleCircle := circle{10}
    writeArea(simpleCircle)
}

