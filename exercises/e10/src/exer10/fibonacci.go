package exer10

func fibConcurrent(n uint) uint {
	return Fib(n-1, 2) + Fib(n-2, 2)
}

func Fib(n uint, cutoff uint) uint {
	if n < cutoff {
		// Do concurrency
	} else {
		if n == 0 || n == 1 {
			return 1
		} else {
			return Fib(n-1, 2) + Fib(n-2, 2)
		}
	}
}

func Fibonacci(n uint) uint {
	return Fib(n, 0)
}
