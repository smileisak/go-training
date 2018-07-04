package main

import st "strings"
import "fmt"

/*
The standard library’s strings package provides many useful string-related functions.
Here are some examples to give you a sense of the package.
*/

var p = fmt.Println

// StringFunctions illustrate `strings` package in go
func StringFunctions() {
	p("Contains : ", st.Contains("test", "es"))
	p("Count	: ", st.Count("test", "t"))
	p("HasPrefix: ", st.HasPrefix("test", "te"))
	p("HasSuffix: ", st.HasSuffix("test", "st"))
	p("Index	: ", st.Index("test", "e"))
	p("Join		: ", st.Join([]string{"a", "b"}, "-"))
	p("Repeat	: ", st.Repeat("a", 5))
	p("Replace	: ", st.Replace("foo", "o", "0", -1))
	p("Replace	: ", st.Replace("foo", "o", "0", 1))
	p("Split	: ", st.Split("a-b-c-d-e-f", "-"))
	p("ToLower	: ", st.ToLower("TEST"))
	p("ToUpper	: ", st.ToUpper("test"))
	p()
	p("Len		: ", len("Hello"))
	p("Char		: ", "Hello"[1])
}
