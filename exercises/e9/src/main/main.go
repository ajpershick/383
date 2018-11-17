package main

import (
	"exer9"
	"fmt"
	"math"
)

func main() {
	// My own tests
	p := exer9.Point{0, 1}

	p.Rotate(math.Pi)
	fmt.Println(int(p.X))
	fmt.Println(int(p.Y))

	p.Rotate(math.Pi)
	fmt.Println(int(p.X))
	fmt.Println(int(p.Y))

	// arr := exer9.RandomArray(6, 2000)
	// for i := 0; i < 5; i++ {
	// 	println(arr[i])
	// }

	arr2 := []float64{1.2, 4.5, 2.3, 7.6, 4.5}
	for i := 0; i < 5; i++ {
		println(arr2[i])
	}
	exer9.InsertionSort(arr2)
	for i := 0; i < 5; i++ {
		println(arr2[i])
	}
}
