package slices_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	formaterStringsTest := []struct {
		Name           string
		SliceOfStrings []string
		str            string
		Expected       bool
	}{
		{
			Name:           "Given a string and a slice of strings which contains a string should return true",
			SliceOfStrings: []string{"10"},
			str:            "10",
			Expected:       true,
		},
		{
			Name:           "Given a string and a slice of strings which no contains a string should return false",
			SliceOfStrings: []string{"10"},
			str:            "20",
			Expected:       false,
		},

		{
			Name:           "Given a string and a slice of strings with few values which contains a string should return true",
			SliceOfStrings: []string{"10", "20"},
			str:            "20",
			Expected:       true,
		},
		{
			Name:           "Given a string and a slice of strings with few values which not contains a string should return false",
			SliceOfStrings: []string{"10", "20"},
			str:            "30",
			Expected:       false,
		},
	}

	for _, tt := range formaterStringsTest {
		t.Run(tt.Name, func(t *testing.T) {
			ok := slices.Contains(tt.SliceOfStrings, tt.str)
			assert.Equal(t, tt.Expected, ok)
		})
	}

	formaterIntegersTest := []struct {
		Name            string
		SliceOfIntegers []int
		Integer         int
		Expected        bool
	}{
		{
			Name:            "Given a Integer and a slice of Integers which contains a Integer should return true",
			SliceOfIntegers: []int{10},
			Integer:         10,
			Expected:        true,
		},
		{
			Name:            "Given a Integer and a slice of Integers which no contains a Integer should return false",
			SliceOfIntegers: []int{10},
			Integer:         20,
			Expected:        false,
		},

		{
			Name:            "Given a Integer and a slice of Integers with few values which contains a Integer should return true",
			SliceOfIntegers: []int{10, 20},
			Integer:         20,
			Expected:        true,
		},
		{
			Name:            "Given a Integer and a slice of Integers with few values which not contains a Integer should return false",
			SliceOfIntegers: []int{10, 20},
			Integer:         30,
			Expected:        false,
		},
	}

	for _, tt := range formaterIntegersTest {
		t.Run(tt.Name, func(t *testing.T) {
			ok := slices.Contains(tt.SliceOfIntegers, tt.Integer)
			assert.Equal(t, tt.Expected, ok)
		})
	}
}
