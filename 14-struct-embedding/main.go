package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2.0) + math.Pow(p1.Y-p2.Y, 2.0))
}

// Go does inheritance through struct embedding, also known as composition.
type Circle struct {
	Point  // anonymous field (no name)
	X      float64
	Radius float64
}

func (c Circle) Distance(p2 Point) float64 {
	return math.Abs(c.Radius - c.Point.Distance(p2))
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {

	c1 := Circle{}
	c1.Point.X = 3
	c1.Point.Y = 4

	fmt.Println(c1.Point.X, c1.Point.Y)
	fmt.Println(c1.X, c1.Y) // due to anonymous embedding, we can access Point's fields directly

	var p1 Point = c1.Point

	fmt.Printf("%v, %v\n", p1.X, p1.Y)

	dist := p1.Distance(Point{X: 0, Y: 0}) // method call on Point
	fmt.Printf("Distance from origin: %v\n", dist)
}
