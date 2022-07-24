package slices

import (
	"github.com/pkg/errors"
	"strconv"
)

func ConvertToSliceOfNumbers(data []string) ([]int64, error) {
	var numbers []int64
	for _, numberStr := range data {
		number, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert string to int")
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}
