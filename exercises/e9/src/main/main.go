package main

import (
	"exer9"
	"fmt"
	"math"
)

func main() {
	// My own tests
	p := exer9.Point{1, 0}
	p.Rotate(math.Pi / 2)
	fmt.Println(p)
	p.Rotate(math.Pi / 2)
	fmt.Println(p)

	arr := exer9.RandomArray(6, 2000)
	for i := 0; i < 5; i++ {
		println(arr[i])
	}
}
