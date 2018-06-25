package main

import "fmt"

// Varibales is a declaration example
func Varibales() {
	// Declatre 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared withiout a corresponding initialization are zero-valued. the zero value of init is 0
	var e int
	fmt.Println(e)

	// The syntax is shorthand for declaring and initializing a variable
	// In this case var f string = "short"
	f := "short"
	fmt.Println(f)
}
