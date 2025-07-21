package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow((p.x-other.x), 2) + math.Pow((p.y-other.y), 2))
}

func main() {
	p1 := NewPoint(5, 10)
	p2 := NewPoint(2, 7)

	fmt.Println(p1.Distance(p2))
}
