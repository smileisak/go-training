package main

import "fmt"

// Go supports pointers, allowing you to pass references to values and records within your program.

// We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

// Pointers function illustrate function in go
func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// Pointers can be printed too.
	fmt.Println("Pointer: ", &i)

	// zeroval doesn’t change the i in main,
	// but zeroptr does because it has a reference to the memory address for that variable.
}
