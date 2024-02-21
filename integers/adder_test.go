package integers

import "testing"

func utilAssertCorrectValue(t *testing.T, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	t.Run("add 2 + 3", func(t *testing.T) {
		utilAssertCorrectValue(t, sum, expected)
	})
}
