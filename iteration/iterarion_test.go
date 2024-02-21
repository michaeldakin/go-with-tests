package iteration

import "testing"

func utilAssertCorrectValue(t *testing.T, expected, repeated string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	t.Run("expected 'a' chars repeated 5 times", func(t *testing.T) {
		utilAssertCorrectValue(t, expected, repeated)
	})
}
