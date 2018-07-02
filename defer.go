package main

import (
	"fmt"
	"os"
)

/*
Defer is used to ensure that a function call is performed later in a program’s execution,
usually for purposes of cleanup.
defer is often used where e.g. ensure and finally would be used in other languages.
*/

// Defer func to illustrate deffer in go
func Defer() {
	// Suppose we wanted to create a file,
	// write to it, and then close when we’re done.
	// Here’s how we could do that with defer.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("Creating ...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing ...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing ...")
	f.Close()
}

// Running the program confirms that the file is closed after being written.
