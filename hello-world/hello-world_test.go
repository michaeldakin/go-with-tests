package main

import "testing"

/*
* Rules of writing test in Go
*  1. The file name must end with _test.go
*  2. The test function must start with the word Test
*  3. The test function takes one argument only t *testing.T
*  4. The test function must be in the same package as the code being tested
*  5. In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*
* Using %q adds quotation marks " " around strings
 */

// Utility function to reduce code duplication - not DRY
func utilAssertCorrectMessage(test testing.TB, got, want string) {
	test.Helper()
	if got != want {
		test.Errorf("got %q want %q", got, want)
	}
}

// Test Hello func to verify providing a string return Hello + string, or default hello + world
func TestHello(t *testing.T) {
	t.Run("saying hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		utilAssertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' if no name string is passed", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World!"
		utilAssertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Javier", "es")
		want := "Hola, Javier"
		utilAssertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Oui", "fr")
		want := "Bonjour, Oui"
		utilAssertCorrectMessage(t, got, want)
	})
}
