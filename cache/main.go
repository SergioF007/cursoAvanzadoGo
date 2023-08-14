package main

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

// Constructor del Memory
