package calculator_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	adderTest := []struct {
		Name    string
		Numbers string
		Sum     int
	}{
		{
			Name:    "Given a one number should return the same number",
			Numbers: "1",
			Sum:     1,
		},
		{
			Name:    "Given a two zeros should return 0",
			Numbers: "0,0",
			Sum:     0,
		},
		{
			Name:    "Given a three positive Numbers 1, 2, 3 should return 6",
			Numbers: "1,2,3",
			Sum:     6,
		},
		{
			Name:    "Given a three Numbers with some of them negatives, such as -7,8,-9, should return -8",
			Numbers: "-7,8,-9",
			Sum:     -8,
		},
		{
			Name:    "Given a Numbers and non-Numbers should ignore non-Numbers and return calculator of Numbers",
			Numbers: "add,2,2",
			Sum:     4,
		},
		{
			Name:    "Given a float in a list of Numbers should return the calculator of only integers",
			Numbers: "0,2.4,2",
			Sum:     2,
		},
		{
			Name:    "Given a Numbers which higher than maxInt should return 0",
			Numbers: "3,9223372036854775807",
			Sum:     0,
		},
		{
			Name:    "Given a Numbers which smaller than minInt should return 0",
			Numbers: "-2,-9223372036854775808",
			Sum:     0,
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			calculate := calculator.New()
			got, err := calculate.Add(tt.Numbers)
			if err != nil {
				fmt.Println(err)
			}
			assert.Equal(t, tt.Sum, got)
		})
	}
}
