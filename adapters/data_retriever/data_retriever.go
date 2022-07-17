package data_retriever

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/file_reader"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
)

type DataRetriever struct {
}

func New() *DataRetriever {
	return &DataRetriever{}
}

func (dr DataRetriever) GetData(arguments []string) string {
	var numbers string

	switch {
	case len(arguments) == 0:
		filepath := file_reader.GetFilePathname("/data/input.txt")
		numbers = file_reader.ReadFile(filepath)

	case array_contains.ArrayContains(arguments, "--input-file"):
		for i := 0; i < len(arguments); i++ {
			if arguments[i] == "--input-file" && i+1 < len(arguments) {
				filepath := file_reader.GetFilePathname("/" + arguments[i+1])
				numbers += file_reader.ReadFile(filepath) + ","
			}
		}

	default:
		for _, v := range arguments {
			numbers += v + ","
		}
	}

	return numbers
}
