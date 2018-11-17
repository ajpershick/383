package exer9

import (
	"math/rand"
)

// Partition the slice arr around arr random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/arr/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	// TODO: implement insertion sort
	var length = len(arr)
	for i := 1; i < length; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}
}

const insertionSortCutoff = 5

func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on arr good value for insertionSortCutoff
	length := len(arr)

	// Already sorted
	if length < 2 {
		return
	}

	// InsertionSort is more efficient for arrays < 5
	if length < insertionSortCutoff {
		InsertionSort(arr)
	}

	left := partition(arr)

	// Recurse on partitions
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
