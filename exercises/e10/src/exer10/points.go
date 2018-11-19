package exer10

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

type Triangle struct {
	A, B, C Point
}

func (t Triangle) String() string {
	return fmt.Sprintf("[%s %s %s]", t.A, t.B, t.C)
}

func (t Triangle) Scale(f float64) {
	t.A.Scale(f)
	t.B.Scale(f)
	t.C.Scale(f)
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
