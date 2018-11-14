package main

import (
	"exer8"
	"fmt"
)

func main() {
	// My own tests
	fmt.Println("Hailstone(18) =", exer8.Hailstone(18))
	fmt.Println("Hailstone sequence starting at 5:", exer8.HailstoneSequenceAppend(5))
	fmt.Println("Hailstone sequence starting at 5:", exer8.HailstoneSequenceAllocate(5))
	fmt.Println("newPoint(5, 2) = ", exer8.NewPoint(5, 2))
	fmt.Println("NewPoint(5, 2).String() == (5, 2) =", exer8.NewPoint(5, 2).String() == "(5, 2)")
	fmt.Println("NewPoint(3, 4).Norm() = ", exer8.NewPoint(3, 4).Norm())
}
