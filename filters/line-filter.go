package main

import (
	"bufio"
	"fmt"
	os "os"
	"strings"
)

/*
	A line filter is a common type of program that reads input on stdin,
	processes it, and then prints some derived result to stdout.
	grep and sed are common line filters.

	To run it:
	$ echo 'hello'   > /tmp/lines
	$ echo 'filter' >> /tmp/lines
	Then use the line filter to get uppercase lines.
	$ cat /tmp/lines | go run line-filters.go
*/

// LineFilter function to illustrate how to create a line filter in go.
func main() {
	// Wrapping the unbuffered os.Stdin with a buffered scanner
	// gives us a convenient Scan method that advances the scanner
	// to the next token; which is the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// Text returns the current token, here the next line, from the input.
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
