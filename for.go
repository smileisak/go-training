package main

import "fmt"

// For function details for loop in go
func For() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("Loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
