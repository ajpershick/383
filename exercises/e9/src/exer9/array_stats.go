package exer9

import "math/rand"

type Sums struct {
	X float64
	Y float64
}

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxInt)
	}

	return arr
}

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.
	
	return

}
