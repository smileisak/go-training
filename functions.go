package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

// Functions details functions in golang
func Functions() {
	// Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)

	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)

}
