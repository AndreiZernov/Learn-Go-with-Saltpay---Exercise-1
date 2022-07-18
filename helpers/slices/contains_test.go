package slices_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	formaterStringsTest := []struct {
		name           string
		sliceOfStrings []string
		str            string
		expected       bool
	}{
		{name: "Given a string and a slice of strings which contains a string should return true", sliceOfStrings: []string{"10"}, str: "10", expected: true},
		{name: "Given a string and a slice of strings which no contains a string should return false", sliceOfStrings: []string{"10"}, str: "20", expected: false},
		{name: "Given a string and a slice of strings with few values which contains a string should return true", sliceOfStrings: []string{"10", "20"}, str: "20", expected: true},
		{name: "Given a string and a slice of strings with few values which not contains a string should return false", sliceOfStrings: []string{"10", "20"}, str: "30", expected: false},
	}

	for _, tt := range formaterStringsTest {
		t.Run(tt.name, func(t *testing.T) {
			ok := slices.Contains(tt.sliceOfStrings, tt.str)
			assert.Equal(t, tt.expected, ok)
		})
	}

	formaterIntegersTest := []struct {
		name            string
		sliceOfIntegers []int
		integer         int
		expected        bool
	}{
		{name: "Given a integer and a slice of integers which contains a integer should return true", sliceOfIntegers: []int{10}, integer: 10, expected: true},
		{name: "Given a integer and a slice of integers which no contains a integer should return false", sliceOfIntegers: []int{10}, integer: 20, expected: false},
		{name: "Given a integer and a slice of integers with few values which contains a integer should return true", sliceOfIntegers: []int{10, 20}, integer: 20, expected: true},
		{name: "Given a integer and a slice of integers with few values which not contains a integer should return false", sliceOfIntegers: []int{10, 20}, integer: 30, expected: false},
	}

	for _, tt := range formaterIntegersTest {
		t.Run(tt.name, func(t *testing.T) {
			ok := slices.Contains(tt.sliceOfIntegers, tt.integer)
			assert.Equal(t, tt.expected, ok)
		})
	}
}
