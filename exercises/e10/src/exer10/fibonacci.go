package exer10

// func FibConcurrent(n uint, c chan uint) {
// 	if n == 0 || n == 1 {
// 		c <- 1
// 	}
// 	c <- n
// 	FibConcurrent(n-1, c) + FibConcurrent(n-2, c)
// }

// func Fib(n uint, cutoff uint) uint {
// 	if n > cutoff {
// 		// Do concurrency
// 		c := make(chan uint)
// 		go FibConcurrent(n, c)
// 	} else {

// 	}
// }

// func Fibonacci(n uint) uint {
// 	return Fib(n, 0)
// }
