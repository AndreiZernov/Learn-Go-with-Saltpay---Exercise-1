package data_retriever_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataRetriever_GetData(t *testing.T) {
	adderTest := []struct {
		name      string
		arguments []string
		numbers   string
	}{
		{name: "Given an arguments --input-file data/input.txt should return numbers inside", arguments: []string{"--input-file", "data/input.txt"}, numbers: "4,5,32,100,867543,"},
		{name: "Given an arguments --input-file twice should return string of numbers merged from both files", arguments: []string{"--input-file", "data/input.txt", "--input-file", "data/input2.csv"}, numbers: "4,5,32,100,867543,4,5,32,100,867543,"},
		{name: "Given an arguments as numbers in one string should return string of numbers as it is", arguments: []string{"-2, 3, 4"}, numbers: "-2, 3, 4,"},
		{name: "Given an arguments as numbers in wto string should return merged string of numbers", arguments: []string{"-2, 3, 4", "-2, 3, 4"}, numbers: "-2, 3, 4,-2, 3, 4,"},
		{name: "Given no arguments should return default numbers from data/input.txt file", arguments: []string{}, numbers: "4,5,32,100,867543"},
	}

	for _, tt := range adderTest {
		t.Run(tt.name, func(t *testing.T) {
			dataRetriever := data_retriever.New()
			got := dataRetriever.GetData(tt.arguments)
			assert.Equal(t, tt.numbers, got)
		})
	}
}
