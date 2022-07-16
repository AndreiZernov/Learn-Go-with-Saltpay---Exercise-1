package sum

import "testing"
import "fmt"

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	checkAdder := func(t testing.TB, expected, sum int) {
		t.Helper()
		if expected != sum {
			t.Errorf("expected %d sum %d", expected, sum)
		}
	}

	t.Run("expected 2 + 2 equal 4", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4
		checkAdder(t, expected, sum)
	})

	t.Run("expected 2 + 6 equal 8", func(t *testing.T) {
		sum := add(6, 2)
		expected := 8
		checkAdder(t, expected, sum)
	})

	t.Run("expected 0 + 0 equal 0", func(t *testing.T) {
		sum := add(0, 0)
		expected := 0
		checkAdder(t, expected, sum)
	})

	t.Run("expected 0 + 1 equal 1", func(t *testing.T) {
		sum := add(0, 1)
		expected := 1
		checkAdder(t, expected, sum)
	})
}
