package data_cleaner_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/data_cleaner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataCleaner(t *testing.T) {
	formaterTest := []struct {
		name     string
		str      string
		expected string
	}{
		{name: "Given a string with comma and space between items should return formatted string", str: "10, 20, 30, 40, 50", expected: "10,20,30,40,50"},
		{name: "Given a string with correct format should return the same string", str: "10,20,30,40,50", expected: "10,20,30,40,50"},
		{name: "Given a string with new line between items should return just numbers divided by commas", str: "10\n20\n30\n40\n50", expected: "10,20,30,40,50"},
		{name: "Given a string with duplicated numbers should return the string without duplication", str: "10,10,20,20,30,30", expected: "10,20,30"},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			newStr := data_cleaner.DataCleaner(tt.str)
			assert.Equal(t, tt.expected, newStr)
		})
	}
}
