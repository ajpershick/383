package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions
func Hailstone(n uint) uint {
	if n%2 == 0 {
		return (n / 2)
	} else {
		return 3*n + 1
	}
}

func HailstoneSequenceAppend(n uint) []uint {
	slice := []uint{n}
	for n > 1 {
		n = Hailstone(n)
		slice = append(slice, n)
	}
	return slice
}

func HailstoneSequenceAllocate(n uint) []uint {
	count := 1
	for n > 1 {
		n = Hailstone(n)
		count++
	}
	sequence := make([]uint, count)
	for n > 1 {
		n = Hailstone(n)
		sequence
	}
}
