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
