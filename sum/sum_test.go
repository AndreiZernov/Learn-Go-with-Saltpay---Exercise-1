package sum

import "testing"

func TestAdder(t *testing.T) {
	adderTest := []struct {
		name    string
		numbers Numbers
		sum     int
	}{
		{numbers: []string{"2", "2"}, sum: 4},
		{numbers: []string{"6", "2"}, sum: 8},
		{numbers: []string{"0", "0"}, sum: 0},
		{numbers: []string{"1", "2", "3"}, sum: 6},
		{numbers: []string{"4", "5", "6"}, sum: 15},
		{numbers: []string{"-7", "8", "-9"}, sum: -8},
		{numbers: []string{"-1", "-2", "-3"}, sum: -6},
		{numbers: []string{"add", "2", "2"}, sum: 4},
		{numbers: []string{"add", "2.4", "2"}, sum: 2},
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
