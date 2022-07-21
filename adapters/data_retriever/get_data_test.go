package data_retriever_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataRetriever_GetData(t *testing.T) {
	adderTest := []struct {
		Name      string
		Arguments []string
		Numbers   string
	}{
		{
			Name:      "Given an Arguments --input-file data/input.txt should return Numbers inside",
			Arguments: []string{"--input-file", "data/input.txt"},
			Numbers:   "4\n5\n32\n100\n867543,",
		},
		{
			Name:      "Given an Arguments --input-file twice should return string of Numbers merged from both files",
			Arguments: []string{"--input-file", "data/input.txt", "--input-file", "data/input2.csv"},
			Numbers:   "4\n5\n32\n100\n867543,4,5,32,100,867543,",
		},
		{
			Name:      "Given an Arguments as Numbers in one string should return string of Numbers as it is",
			Arguments: []string{"-2, 3, 4"},
			Numbers:   "-2, 3, 4,",
		},
		{
			Name:      "Given an Arguments as Numbers in wto string should return merged string of Numbers",
			Arguments: []string{"-2, 3, 4", "-2, 3, 4"},
			Numbers:   "-2, 3, 4,-2, 3, 4,",
		},
		{
			Name:      "Given no Arguments should return default Numbers from data/input.txt file",
			Arguments: []string{},
			Numbers:   "4\n5\n32\n100\n867543",
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			dataRetriever := data_retriever.New()
			got := dataRetriever.GetData(tt.Arguments)
			assert.Equal(t, tt.Numbers, got)
		})
	}
}
