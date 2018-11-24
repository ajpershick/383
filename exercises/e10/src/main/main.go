package main

import (
	"exer10"
	"fmt"
	"math"
)

func main() {
	// My own tests
	// println(exer10.Fibonacci(10))

	// t1 := exer10.Triangle{exer10.Point{1, 2}, exer10.Point{2, 3}, exer10.Point{3, 4}}
	// t2 := exer10.Triangle{exer10.Point{5, 10}, exer10.Point{10, 15}, exer10.Point{15, 20}}
	// t1.Scale(5)
	// println(t1 == t2)
	// println(int(t1.A.X))
	// println(int(t1.A.Y))
	// t1.Rotate(math.Pi)
	// println(int(t1.A.X))
	// println(int(t1.A.Y))
	// t1.Rotate(math.Pi)
	// println(int(t1.A.X))
	// println(int(t1.A.Y)) // Rotated 360 degrees should be back to same values
	pt := exer10.Point{3, 4}
	exer10.TurnDouble(&pt, 3*math.Pi/2)
	fmt.Println(pt)
	tri := exer10.Triangle{exer10.Point{1, 2}, exer10.Point{-3, 4}, exer10.Point{5, -6}}
	exer10.TurnDouble(&tri, math.Pi)
	fmt.Println(tri)
}
