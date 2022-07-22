package calculator_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	adderTest := []struct {
		Name    string
		Numbers string
		Sum     int
		Error   bool
	}{
		{
			Name:    "Given a one number should return the same number",
			Numbers: "1",
			Sum:     1,
			Error:   false,
		},
		{
			Name:    "Given a two zeros should return 0",
			Numbers: "0,0",
			Sum:     0,
			Error:   false,
		},
		{
			Name:    "Given a three positive Numbers 1, 2, 3 should return 6",
			Numbers: "1,2,3",
			Sum:     6,
			Error:   false,
		},
		{
			Name:    "Given a three Numbers with some of them negatives, such as -7,8,-9, should return -8",
			Numbers: "-7,8,-9",
			Sum:     -8,
			Error:   false,
		},
		{
			Name:    "Given a Numbers and non-Numbers should ignore non-Numbers and return calculator of Numbers",
			Numbers: "add,2,2",
			Sum:     4,
			Error:   false,
		},
		{
			Name:    "Given a float in a list of Numbers should return the calculator of only integers",
			Numbers: "0,2.4,2",
			Sum:     2,
			Error:   false,
		},
		{
			Name:    "Given a Numbers which higher than maxInt should return 0",
			Numbers: "3,9223372036854775807",
			Sum:     0,
			Error:   true,
		},
		{
			Name:    "Given a Numbers which smaller than minInt should return 0",
			Numbers: "-2,-9223372036854775808",
			Sum:     0,
			Error:   true,
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			calculate := calculator.New()
			got, err := calculate.Add(tt.Numbers)

			assert.Equal(t, tt.Error, err != nil)
			assert.Equal(t, tt.Sum, got)
		})
	}
}
