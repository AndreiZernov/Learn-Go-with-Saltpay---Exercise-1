package strings_helper_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataCleaner(t *testing.T) {
	formaterTest := []struct {
		Name          string
		StringToClean string
		Expected      string
	}{
		{
			Name:          "Given a strings_helper with comma and space between items should return formatted strings_helper",
			StringToClean: "10, 20, 30, 40, 50",
			Expected:      "10,20,30,40,50",
		},
		{
			Name:          "Given a strings_helper with correct format should return the same strings_helper",
			StringToClean: "10,20,30,40,50",
			Expected:      "10,20,30,40,50",
		},
		{
			Name:          "Given a strings_helper with new line between items should return just numbers divided by commas",
			StringToClean: "10\n20\n30\n40\n50",
			Expected:      "10,20,30,40,50",
		},
		{
			Name:          "Given a strings_helper with duplicated numbers should return the strings_helper without duplication",
			StringToClean: "10,10,20,20,30,30",
			Expected:      "10,20,30",
		},
	}

	for _, tt := range formaterTest {
		t.Run(tt.Name, func(t *testing.T) {
			newStr := strings_helper.DataCleaner(tt.StringToClean)
			assert.Equal(t, tt.Expected, newStr)
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	formaterTest := []struct {
		Name          string
		StringToClean string
		Expected      string
	}{
		{
			Name:          "Given a string and a slice which contains a string should return true",
			StringToClean: "1,2,2,4,5,5,7,8,9,10",
			Expected:      "1,2,4,5,7,8,9,10",
		},
		{
			Name:          "Given a sting and a slice which no contains a string should return false",
			StringToClean: "1,1,1,1",
			Expected:      "1",
		},
		{
			Name:          "Given a sting and a slice with few values which contains a string should return true",
			StringToClean: "add,add,2",
			Expected:      "add,2",
		},
		{
			Name:          "Given a sting and a slice with few values which not contains a string should return false",
			StringToClean: "1",
			Expected:      "1",
		},
	}

	for _, tt := range formaterTest {
		t.Run(tt.Name, func(t *testing.T) {
			newString := strings_helper.RemoveDuplicates(tt.StringToClean)
			assert.Equal(t, tt.Expected, newString)
		})
	}
}
