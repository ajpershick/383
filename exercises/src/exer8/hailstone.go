package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions
func Hailstone(n uint) uint {
	if n%2 == 0 {
		return (n / 2)
	} else {
		return 3*n + 1
	}
}
