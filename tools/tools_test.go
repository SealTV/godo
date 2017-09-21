package tools

import "testing"

type sumTestParams struct {
	a int
	b int
	s int
}

var sumTests = []sumTestParams{
	{0, 0, 0},
	{1, 0, 1},
	{0, 1, 1},
	{1, -1, 0},
}

func TestSum(t *testing.T) {
	for _, data := range sumTests {
		s := Sum(data.a, data.b)
		if s != data.s {
			t.Errorf("For %d + %d expected %d got %d", data.a, data.b, data.s, s)
		}
	}
}

type testpair struct {
	value  int
	result int
}

var fibonacciTests = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
}

func TestFibonacci(t *testing.T) {
	for _, data := range fibonacciTests {
		s := Fibonacci(data.value)
		if s != data.result {
			t.Errorf("For %d expected %d got %d", data.value, data.result, s)
		}
	}
}

var factorialTests = []testpair{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
	{6, 720},
	{7, 5040},
	{8, 40320},
	{9, 362880},
	{10, 3628800},
}

func TestFactorial(t *testing.T) {
	for _, data := range factorialTests {
		s := Factorial(data.value)
		if s != data.result {
			t.Errorf("For %d expected %d got %d", data.value, data.result, s)
		}
	}
}
