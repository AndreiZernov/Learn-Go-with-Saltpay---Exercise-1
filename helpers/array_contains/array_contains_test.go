package array_contains_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayContains(t *testing.T) {
	formaterTest := []struct {
		name     string
		slice    []string
		str      string
		expected bool
	}{
		{name: "Given a string and a slice which contains a string should return true", slice: []string{"10"}, str: "10", expected: true},
		{name: "Given a sting and a slice which no contains a string should return false", slice: []string{"10"}, str: "20", expected: false},
		{name: "Given a sting and a slice with few values which contains a string should return true", slice: []string{"10", "20"}, str: "20", expected: true},
		{name: "Given a sting and a slice with few values which not contains a string should return false", slice: []string{"10", "20"}, str: "30", expected: false},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			ok := array_contains.ArrayContains(tt.slice, tt.str)
			assert.Equal(t, tt.expected, ok)
		})
	}
}
