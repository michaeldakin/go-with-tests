package main

import "fmt"

const enHelloPrefix = "Hello, "
const esHelloPrefix = "Hola, "

func Hello(name, lang string) string {
	// If no name is supplied, defualt "World!"
	if name == "" {
		name = "World!"
	}

	// language check after valid name check
	if lang == "es" {
		return esHelloPrefix + name
	}

	// We must have a name and are not a language opt we offer, default to English
	return enHelloPrefix + name

}

func main() {
	greeting := Hello("Chris", "en")
	fmt.Println(greeting)
}
