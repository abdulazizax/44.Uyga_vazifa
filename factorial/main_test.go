package main

import "testing"

func TestFactorial(t *testing.T) {
	test := []struct{
		n int64
		factorial int64
	}{
		{0, 1},
		{5, 120},
		{1, 1},
		{10, 3628800},
		{15, 1307674368000},
	}

	for _, v := range test {
		expected := v.factorial
		actual := Factorial(v.n)
		if expected != actual {
			t.Errorf("Expected %v do not match actual %v", expected, actual)
		}
	}
}
