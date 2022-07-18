package data_retriever

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/file_reader"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
)

type DataRetriever struct {
}

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
	filepath := file_reader.GetFilePathname("/data/input.txt")
	return file_reader.ReadFile(filepath)
}

func (dr DataRetriever) retrieveWithInputFilesArguments(arguments []string) string {
	numbers := ""
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "--input-file" && i+1 < len(arguments) {
			filepath := file_reader.GetFilePathname("/" + arguments[i+1])
			numbers += file_reader.ReadFile(filepath) + ","
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
