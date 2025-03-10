package fib

func Fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func QuickFibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	a, b := uint64(0), uint64(1)
	for i := uint64(2); i <= n; i++ {
		temp := a + b
		a, b = b, temp
	}

	return b
}
