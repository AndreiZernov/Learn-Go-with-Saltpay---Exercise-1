package data_retriever

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"github.com/pkg/errors"
	"strings"
)

const defaultFilePathname = "/data/input.txt"

type DataRetriever struct{}

func New() *DataRetriever {
	return &DataRetriever{}
}

func (dr DataRetriever) GetNumbers(arguments []string) ([]int64, error) {
	var numbers []int64
	switch {
	case len(arguments) == 0:
		data, err := files.ReadFile(defaultFilePathname)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read file from default path")
		}

		splittedData := strings.FieldsFunc(data, dr.split)
		return slices.ConvertToSliceOfNumbers(splittedData)

	case slices.Contains(arguments, "--input-file"):
		for i := 0; i < len(arguments); i++ {
			if arguments[i] == "--input-file" && i+1 < len(arguments) {
				data, err := files.ReadFile("/" + arguments[i+1])
				if err != nil {
					return nil, errors.Wrap(err, "failed to read file from specified path")
				}

				splittedData := strings.FieldsFunc(data, dr.split)
				convertedData, err := slices.ConvertToSliceOfNumbers(splittedData)
				if err != nil {
					return nil, errors.Wrap(err, "failed to convert string to int")
				}

				numbers = append(numbers, convertedData[:]...)
			}
		}
		return numbers, nil

	default:
		for _, argument := range arguments {
			splittedData := strings.FieldsFunc(argument, dr.split)
			convertedData, err := slices.ConvertToSliceOfNumbers(splittedData)
			if err != nil {
				return nil, errors.Wrap(err, "failed to convert string to int")
			}

			numbers = append(numbers, convertedData[:]...)
		}
		return numbers, nil
	}
}

func (dr DataRetriever) split(r rune) bool {
	return r == '\n' || r == ',' || r == '\t' || r == ' ' || r == '.'
}
