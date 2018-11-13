package exer8

import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods
type Point struct {
	x float64
	y float64
}

// Point constructor
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (point Point) String() string {
	return fmt.Sprintf("(%v, %v)", point.x, point.y)
}

func (point Point) Norm() float64 {
	return math.Sqrt(point.x*point.x + point.y*point.y)
}
