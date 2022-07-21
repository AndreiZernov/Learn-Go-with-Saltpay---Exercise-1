package slices_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	formaterStringsTest := []struct {
		Name            string
		SliceOfStrings  []string
		StringToContain string
		Expected        bool
	}{
		{
			Name:            "Given a string and a slice of strings which contains a string should return true",
			SliceOfStrings:  []string{"10"},
			StringToContain: "10",
			Expected:        true,
		},
		{
			Name:            "Given a string and a slice of strings which no contains a string should return false",
			SliceOfStrings:  []string{"10"},
			StringToContain: "20",
			Expected:        false,
		},

		{
			Name:            "Given a string and a slice of strings with few values which contains a string should return true",
			SliceOfStrings:  []string{"10", "20"},
			StringToContain: "20",
			Expected:        true,
		},
		{
			Name:            "Given a string and a slice of strings with few values which not contains a string should return false",
			SliceOfStrings:  []string{"10", "20"},
			StringToContain: "30",
			Expected:        false,
		},
	}

	for _, tt := range formaterStringsTest {
		t.Run(tt.Name, func(t *testing.T) {
			ok := slices.Contains(tt.SliceOfStrings, tt.StringToContain)
			assert.Equal(t, tt.Expected, ok)
		})
	}

	formaterIntegersTest := []struct {
		Name             string
		SliceOfIntegers  []int
		IntegerToContain int
		Expected         bool
	}{
		{
			Name:             "Given a IntegerToContain and a slice of Integers which contains a IntegerToContain should return true",
			SliceOfIntegers:  []int{10},
			IntegerToContain: 10,
			Expected:         true,
		},
		{
			Name:             "Given a IntegerToContain and a slice of Integers which no contains a IntegerToContain should return false",
			SliceOfIntegers:  []int{10},
			IntegerToContain: 20,
			Expected:         false,
		},

		{
			Name:             "Given a IntegerToContain and a slice of Integers with few values which contains a IntegerToContain should return true",
			SliceOfIntegers:  []int{10, 20},
			IntegerToContain: 20,
			Expected:         true,
		},
		{
			Name:             "Given a IntegerToContain and a slice of Integers with few values which not contains a IntegerToContain should return false",
			SliceOfIntegers:  []int{10, 20},
			IntegerToContain: 30,
			Expected:         false,
		},
	}

	for _, tt := range formaterIntegersTest {
		t.Run(tt.Name, func(t *testing.T) {
			ok := slices.Contains(tt.SliceOfIntegers, tt.IntegerToContain)
			assert.Equal(t, tt.Expected, ok)
		})
	}
}
