package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return helloPrefix + name

}

func main() {
	greeting := Hello("Chris")
	fmt.Println(greeting)
}
