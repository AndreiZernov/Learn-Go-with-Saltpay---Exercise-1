package data_retriever

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"github.com/pkg/errors"
)

const defaultFilePathname = "/data/input.txt"

type DataRetriever struct{}

func New() *DataRetriever {
	return &DataRetriever{}
}

func (dr DataRetriever) GetData(arguments []string) (string, error) {
	var numbers string
	switch {
	case len(arguments) == 0:
		data, err := files.ReadFile(defaultFilePathname)
		if err != nil {
			return "", errors.Wrap(err, "failed to read file from default path")
		}
		return data, nil

	case slices.Contains(arguments, "--input-file"):
		for i := 0; i < len(arguments); i++ {
			if arguments[i] == "--input-file" && i+1 < len(arguments) {
				data, err := files.ReadFile("/" + arguments[i+1])
				if err != nil {
					return "", errors.Wrap(err, "failed to read file from specified path")
				}
				numbers += data + ","
			}
		}
		return numbers, nil

	default:
		for _, v := range arguments {
			numbers += v + ","
		}
		return numbers, nil
	}
}
