package exer9

// TODO: Point (with everything from exercise 8 and) and methods that modify them in-place
import (
	"fmt"
	"math"
)

// TODO: The Point struct, NewPoint function, .String and .Norm methods
type Point struct {
	X float64
	Y float64
}

// Point constructor
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (point Point) String() string {
	return fmt.Sprintf("(%v, %v)", point.X, point.Y)
}

func (point Point) Norm() float64 {
	return math.Sqrt(point.X*point.X + point.Y*point.Y)
}

func (point *Point) Scale(f float64) {
	point.X = point.X * f
	point.Y = point.Y * f
}

func (point *Point) Rotate(a float64) {
	tempX := point.X*math.Cos(a) - point.Y*math.Sin(a)
	tempY := point.X*math.Sin(a) + point.Y*math.Cos(a)
	point.X = tempX
	point.Y = tempY
}
