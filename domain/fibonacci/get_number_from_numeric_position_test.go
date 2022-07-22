package fibonacci_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/fibonacci"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberFromNumericPosition(t *testing.T) {
	adderTest := []struct {
		Name                string
		PositionInFibonacci int64
		FibonacciNumber     int64
	}{
		{Name: "Given the position of the first number in the Fibonacci sequence should return 0", PositionInFibonacci: 1, FibonacciNumber: 0},
		{Name: "Given the any number below or equal 0 should return 0", PositionInFibonacci: 0, FibonacciNumber: 0},
		{Name: "Given the position of the third number in the Fibonacci sequence should return 1", PositionInFibonacci: 3, FibonacciNumber: 1},
		{Name: "Given the position 93 should be able to calculate based on int64 limitation", PositionInFibonacci: 93, FibonacciNumber: 7540113804746344448},
		{Name: "Given the position 94 should return overflow error and return 0", PositionInFibonacci: 94, FibonacciNumber: 0},
		{Name: "Given negative number should return 0 and outside of sequence error", PositionInFibonacci: -2, FibonacciNumber: 0},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			fib := fibonacci.New()
			got, err := fib.GetNumberFromNumericPosition(tt.PositionInFibonacci)
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.FibonacciNumber, got)
		})
	}

}
