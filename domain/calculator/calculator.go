package calculator

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Calculator struct {
}

func New() *Calculator {
	return &Calculator{}
}

const maxInt = math.MaxInt64
const minInt = math.MinInt64

var ErrOverflow = errors.New("integer overflow")

func (c Calculator) Add(n string) (int, error) {
	sum := 0
	newArray := strings.Split(n, ",")

	for _, number := range newArray {
		x, err := strconv.Atoi(number)

		if err == nil {
			if x > 0 {
				if sum > maxInt-x {
					return 0, ErrOverflow
				}
			} else {
				if sum < minInt-x {
					return 0, ErrOverflow
				}
			}
			sum += x
		}
	}

	return sum, nil
}
