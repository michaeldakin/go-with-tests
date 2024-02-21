package main

import "fmt"

const defaultHelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	// If no name is supplied, defualt "World!"
	if name == "" {
		name = "World!"
	}

	// Determine prefix
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
	var prefix string
	// language check after valid name check
	switch lang {
	case "es":
		prefix = esHelloPrefix
	case "fr":
		prefix = frHelloPrefix
	default:
		prefix = defaultHelloPrefix
	}

	return prefix
}

func main() {
	greeting := Hello("Chris", "en")
	fmt.Println(greeting)
}
