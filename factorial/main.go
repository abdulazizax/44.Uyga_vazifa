package main

func Factorial(n int64) int64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
