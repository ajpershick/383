package main

import (
	"exer10"
)

func main() {
	// My own tests
	println(exer10.Fibonacci(10))

	t1 := exer10.Triangle{exer10.Point{1, 2}, exer10.Point{2, 3}, exer10.Point{3, 4}}
	t2 := exer10.Triangle{exer10.Point{5, 10}, exer10.Point{10, 15}, exer10.Point{15, 20}}
	t1.Scale(5)
	println(t1 == t2)
	println(int(t1.A.X))
}
