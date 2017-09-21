package tools

func Sum(a, b int) int {
	return a + b
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n <= 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
