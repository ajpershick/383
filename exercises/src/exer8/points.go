package exer8

// TODO: The Point struct, NewPoint function, .String and .Norm methods
type Point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) Point {
	return Point{x, y}
}
