package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	greeting := Hello("Chris")
	fmt.Println(greeting)
}
