package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const arabicHelloPrefix = "أهلا, "

const french = "French"
const spanish = "Spanish"
const arabic = "Arabic"

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case arabic:
		prefix = arabicHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello return a hello message
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
