package main

import (
	"fmt"
	"os"
)

// Exit function to illustrate how to exit for go program.
func Exit() {

	// defers will not be run when using os.Exit,
	// so this fmt.Println will never be called.
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)

	/*
		Note that unlike e.g. C, Go does not use an integer return value
		from main to indicate exit status.
		If you’d like to exit with a non-zero status you should
		use os.Exit.
	*/
}
