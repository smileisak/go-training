package main

import "fmt"

// Go supports recursive functions. Here’s a classic factorial example.
func fact(n int) int {
	// This fact function calls itself until it reaches the base case of fact(0).
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Recursive illustrate a recursive example
func Recursive() {
	fmt.Println(fact(7))
}
