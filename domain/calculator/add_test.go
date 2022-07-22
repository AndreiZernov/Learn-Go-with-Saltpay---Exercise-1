package calculator_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	adderTest := []struct {
		Name    string
		Numbers []int64
		Sum     int64
		Error   bool
	}{
		{
			Name:    "Given a three positive Numbers 1, 2, 3 should return 6",
			Numbers: []int64{1, 2, 3},
			Sum:     6,
			Error:   false,
		},
		{
			Name:    "Given a three Numbers with some of them negatives, such as -7,8,-9, should return -8",
			Numbers: []int64{-7, 8, -9},
			Sum:     -8,
			Error:   false,
		},
		{
			Name:    "Given a Numbers which higher than maxInt should return 0",
			Numbers: []int64{2, 9223372036854775807},
			Sum:     0,
			Error:   true,
		},
		{
			Name:    "Given a Numbers which smaller than minInt should return 0",
			Numbers: []int64{-2, -9223372036854775808},
			Sum:     0,
			Error:   true,
		},
		{
			Name:    "Given a Numbers with duplications should return the sum of the unique numbers",
			Numbers: []int64{2, 2, 2},
			Sum:     2,
			Error:   false,
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
