package exer9

func RandomArray(length int, maxInt int) []int {
	// TODO: create a new random generator with a decent seed; create an array with length values from 0 to values-1.
}

func MeanStddev(arr []int, chunks int) (mean, stddev float64) {
	if len(arr)%chunks != 0 {
		panic("You promised that chunks would divide slice size!")
	}
	// TODO: calculate the mean and population standard deviation of the array, breaking the array into chunks segments
	// and calculating on them in parallel.

}
