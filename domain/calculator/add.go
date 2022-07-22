package calculator

import (
	"errors"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"math"
)

type Calculator struct {
}

func New() *Calculator {
	return &Calculator{}
}

const (
	maxInt = math.MaxInt64
	minInt = math.MinInt64
)

var ErrOverflow = errors.New("integer overflow")

func (c Calculator) Add(numbers []int64) (int64, error) {
	var (
		numbersWithNoDuplications []int64
		sum                       int64
	)

	// Removing duplications
	for _, number := range numbers {
		if slices.Contains(numbersWithNoDuplications, number) == false {
			numbersWithNoDuplications = append(numbersWithNoDuplications, number)
		}
	}

	for _, number := range numbersWithNoDuplications {
		if number > 0 {
			if sum > maxInt-number {
				return 0, ErrOverflow
			}
		} else {
			if sum < minInt-number {
				return 0, ErrOverflow
			}
		}
		sum += number
	}

	return sum, nil
}
