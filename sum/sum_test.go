package sum

import "testing"
import "fmt"

func ExampleAdd() {
	numbers := []int{1, 5}

	sum := add(numbers)
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
		numbers := []int{2, 2}
		sum := add(numbers)
		expected := 4
		checkAdder(t, expected, sum)
	})

	t.Run("expected 2 + 6 equal 8", func(t *testing.T) {
		numbers := []int{6, 2}

		sum := add(numbers)
		expected := 8
		checkAdder(t, expected, sum)
	})

	t.Run("expected 0 + 0 equal 0", func(t *testing.T) {
		numbers := []int{0, 0}

		sum := add(numbers)
		expected := 0
		checkAdder(t, expected, sum)
	})

	t.Run("expected 1 + 2 + 3 equal 1", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		sum := add(numbers)
		expected := 6
		checkAdder(t, expected, sum)
	})

	t.Run("expected 0 + 1 + 2 equal 1", func(t *testing.T) {
		numbers := []int{4, 5, 6}

		sum := add(numbers)
		expected := 15
		checkAdder(t, expected, sum)
	})
}
