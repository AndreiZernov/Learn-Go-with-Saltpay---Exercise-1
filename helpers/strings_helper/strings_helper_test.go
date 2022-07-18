package strings_helper_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataCleaner(t *testing.T) {
	formaterTest := []struct {
		name     string
		str      string
		expected string
	}{
		{name: "Given a strings_helper with comma and space between items should return formatted strings_helper", str: "10, 20, 30, 40, 50", expected: "10,20,30,40,50"},
		{name: "Given a strings_helper with correct format should return the same strings_helper", str: "10,20,30,40,50", expected: "10,20,30,40,50"},
		{name: "Given a strings_helper with new line between items should return just numbers divided by commas", str: "10\n20\n30\n40\n50", expected: "10,20,30,40,50"},
		{name: "Given a strings_helper with duplicated numbers should return the strings_helper without duplication", str: "10,10,20,20,30,30", expected: "10,20,30"},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			newStr := strings_helper.DataCleaner(tt.str)
			assert.Equal(t, tt.expected, newStr)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
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
			newString := strings_helper.RemoveDuplicates(tt.str)
			assert.Equal(t, tt.expected, newString)
		})
	}
}
