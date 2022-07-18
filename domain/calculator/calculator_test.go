package calculator_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	adderTest := []struct {
		name    string
		numbers string
		sum     int
	}{
		{
			name:    "Given a one number should return the same number",
			numbers: "1",
			sum:     1,
		},
		{
			name:    "Given a two zeros should return 0",
			numbers: "0,0",
			sum:     0,
		},
		{
			name:    "Given a three positive numbers 1, 2, 3 should return 6",
			numbers: "1,2,3",
			sum:     6,
		},
		{
			name:    "Given a three numbers with some of them negatives, such as -7,8,-9, should return -8",
			numbers: "-7,8,-9",
			sum:     -8,
		},
		{
			name:    "Given a numbers and non-numbers should ignore non-numbers and return calculator of numbers",
			numbers: "add,2,2",
			sum:     4,
		},
		{
			name:    "Given a float in a list of numbers should return the calculator of only integers",
			numbers: "0,2.4,2",
			sum:     2,
		},
		{
			name:    "Given a numbers which higher than maxInt should return 0",
			numbers: "3,9223372036854775807",
			sum:     0,
		},
		{
			name:    "Given a numbers which smaller than minInt should return 0",
			numbers: "-2,-9223372036854775808",
			sum:     0,
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.name, func(t *testing.T) {
			calculate := calculator.New()
			got, err := calculate.Add(tt.numbers)
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, tt.sum, got)
		})
	}
}
