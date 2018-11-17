package exer9

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomArrays(t *testing.T) {
	length := 1000
	maxint := 100
	arr1 := RandomArray(length, maxint)
	assert.Equal(t, length, len(arr1))
	for _, v := range arr1 {
		assert.True(t, 0 <= v, "contains a negative integer")
		assert.True(t, v < maxint, "contains an integer >=maxint")
	}

	// check that different calls return different results
	arr2 := RandomArray(length, maxint)
	assert.False(t, reflect.DeepEqual(arr1, arr2))
}

func TestScale(t *testing.T) {
	p := Point{1, 2}
	p.Scale(5)
	assert.Equal(t, p, Point{5, 10})
	p.Scale(0.5)
	assert.Equal(t, p, Point{2.5, 5})
}

// func TestRotate(t *testing.T) {
// 	p := Point{0, 1}
// 	p.Rotate(math.Pi)
// 	assert.Equal(t, Point{int(p.X), int(p.Y)}, Point{0, -1})
// 	p.Rotate(math.Pi)
// 	assert.Equal(t, p, Point{0, 1})
// }

// func TestArrayStatParallel(t *testing.T) {
// 	length := 30000000
// 	maxint := 10000
// 	arr2 := RandomArray(length, maxint)

// 	// call MeanStddev single-threaded
// 	start := time.Now()
// 	mean2, stddev2 := MeanStddev(arr2, 1)
// 	end := time.Now()
// 	dur1 := end.Sub(start)

// 	// now turn on cuncurrency and make sure we get the same results, but faster
// 	start = time.Now()
// 	mean3, stddev3 := MeanStddev(arr2, 3)
// 	end = time.Now()
// 	dur2 := end.Sub(start)

// 	speedup := float64(dur1) / float64(dur2)
// 	assert.True(t, speedup > 1.25, "Running MeanStddev with concurrency didn't speed up as expected. Sped up by %g.", speedup)
// 	assert.Equal(t, float32(mean2), float32(mean3)) // compare as float32 to avoid rounding differences
// 	assert.Equal(t, float32(stddev2), float32(stddev3))
// }

func benchmarkSorting(b *testing.B, sort func(arr []float64)) {
	const length = 1000
	arr := make([]float64, length)
	for i := 0; i < length; i++ {
		arr[i] = float64(i)
	}

	// run the benchmark
	for n := 0; n < b.N; n++ {
		rand.Shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
		sort(arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) { benchmarkSorting(b, InsertionSort) }

func BenchmarkQuickSort(b *testing.B) { benchmarkSorting(b, QuickSort) }

func BenchmarkFloat64s(b *testing.B) { benchmarkSorting(b, sort.Float64s) }
