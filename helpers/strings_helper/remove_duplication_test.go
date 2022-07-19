package strings_helper_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
