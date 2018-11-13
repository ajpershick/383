package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions
func Hailstone(n uint) uint {
	if n%2 == 0 {
		return (n / 2)
	} else {
		return 3*n + 1
	}
}

// 3891 ns/op
func HailstoneSequenceAppend(n uint) []uint {
	slice := []uint{n}

	for n > 1 {
		n = Hailstone(n)
		slice = append(slice, n)
	}

	return slice
}

// 1740 ns/op
func HailstoneSequenceAllocate(n uint) []uint {
	val := n
	count := 1

	for n > 1 {
		n = Hailstone(n)
		count++
	}

	sequence := make([]uint, count)
	count = 0
	n = val
	sequence[0] = n

	for n > 1 {
		n = Hailstone(n)
		count++
		sequence[count] = n
	}

	return sequence
}
