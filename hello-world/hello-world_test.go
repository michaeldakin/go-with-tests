// package helloworld
package main

import "testing"

/*
* Rules of writing test in Go
*  1. The file name must end with _test.go
*  2. The test function must start with the word Test
*  3. The test function takes one argument only t *testing.T
*  4. The test function must be in the same package as the code being tested
*  5. In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
 */

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
