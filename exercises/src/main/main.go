package main

import (
	"exer8"
	"fmt"
)

func main() {
	fmt.Println("Hailstone(18) =", exer8.Hailstone(18))
	fmt.Println("Hailstone sequence starting at 5:", exer8.HailstoneSequenceAppend(5))
	fmt.Println("Hailstone sequence starting at 5:", exer8.HailstoneSequenceAllocate(5))
	fmt.Println("newPoint(5,2) = ", exer8.newPoint(5, 2))
}
