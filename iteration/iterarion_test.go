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
