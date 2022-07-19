package data_retriever

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/file_reader"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
)

type DataRetriever struct{}

func New() *DataRetriever {
	return &DataRetriever{}
}

func (dr DataRetriever) GetData(arguments []string) string {
	switch {
	case len(arguments) == 0:
		return dr.retrieveWithNoArgumentGiven()

	case slices.Contains(arguments, "--input-file"):
		return dr.retrieveWithInputFilesArguments(arguments)

	default:
		return dr.retrieveWithStringOfNumbersArguments(arguments)
	}
}

func (dr DataRetriever) retrieveWithNoArgumentGiven() string {
	return file_reader.ReadFile("/data/input.txt")
}

func (dr DataRetriever) retrieveWithInputFilesArguments(arguments []string) string {
	var numbers string
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "--input-file" && i+1 < len(arguments) {
			numbers += file_reader.ReadFile("/"+arguments[i+1]) + ","
		}
	}
	return numbers
}

func (dr DataRetriever) retrieveWithStringOfNumbersArguments(arguments []string) string {
	numbers := ""
	for _, v := range arguments {
		numbers += v + ","
	}
	return numbers
}
