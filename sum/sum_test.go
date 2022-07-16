package sum

import "testing"
import "fmt"

func ExampleAdd() {
	numbers := []int{1, 5}

	sum := Numbers.add(numbers)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	adderTest := []struct {
		name    string
		numbers Numbers
		sum     int
	}{
		{numbers: []int{2, 2}, sum: 4},
		{numbers: []int{6, 2}, sum: 8},
		{numbers: []int{0, 0}, sum: 0},
		{numbers: []int{1, 2, 3}, sum: 6},
		{numbers: []int{4, 5, 6}, sum: 15},
		{numbers: []int{-7, 8, -9}, sum: -8},
	}

	for _, tt := range adderTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.numbers.add()
			checkAdder(t, tt.sum, got)
		})
	}
}

func checkAdder(t testing.TB, expected, sum int) {
	t.Helper()
	if expected != sum {
		t.Errorf("expected %d sum %d", expected, sum)
	}
}
