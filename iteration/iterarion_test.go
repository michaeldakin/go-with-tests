package iteration

import "testing"

func utilAssertCorrectValue(t *testing.T, expected, repeated string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	t.Run("expected 'a' chars repeated 5 times", func(t *testing.T) {
		utilAssertCorrectValue(t, expected, repeated)
	})
}

func TestRepeatTrim(t *testing.T) {
	repeated := RepeatUpper("go", 10)
	expected := "gggggggggg"

	t.Run("expected 10 'g' after 'o' is trimmmed", func(t *testing.T) {
		utilAssertCorrectValue(t, expected, repeated)
	})
}

func TestRepeatUpper(t *testing.T) {
	repeated := RepeatUpper("hi", 3)
	expected := "HIHIHI"

	t.Run("expected 'hi' repeated 3 times", func(t *testing.T) {
		utilAssertCorrectValue(t, expected, repeated)
	})
}

func TestRepeatLower(t *testing.T) {
	repeated := RepeatUpper("HI", 3)
	expected := "hihihi"

	t.Run("expected 'hi' repeated 3 times", func(t *testing.T) {
		utilAssertCorrectValue(t, expected, repeated)
	})
}

/*
~/personal/go/go-with-tests/iteration ~/personal/go/go-with-tests
goos: darwin
goarch: arm64
pkg: github.com/michaeldakin/go-with-tests/iteration
BenchmarkTest-11        17173435                70.35 ns/op
PASS
ok      github.com/michaeldakin/go-with-tests/iteration 2.441s
*/

// go test -bench=.
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
