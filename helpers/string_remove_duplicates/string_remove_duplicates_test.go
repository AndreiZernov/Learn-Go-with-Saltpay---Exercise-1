package string_remove_duplicates_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/string_remove_duplicates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayRemoveDuplicates(t *testing.T) {
	formaterTest := []struct {
		name     string
		str      string
		expected string
	}{
		{name: "Given a string and a slice which contains a string should return true", str: "1,2,2,4,5,5,7,8,9,10", expected: "1,2,4,5,7,8,9,10"},
		{name: "Given a sting and a slice which no contains a string should return false", str: "1,1,1,1", expected: "1"},
		{name: "Given a sting and a slice with few values which contains a string should return true", str: "add,add,2", expected: "add,2"},
		{name: "Given a sting and a slice with few values which not contains a string should return false", str: "1", expected: "1"},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			newString := string_remove_duplicates.RemoveDuplicates(tt.str)
			assert.Equal(t, tt.expected, newString)
		})
	}
}
