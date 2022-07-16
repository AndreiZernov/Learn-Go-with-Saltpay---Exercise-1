package sum

import (
	"errors"
	"math"
	"strconv"
)

const maxInt = math.MaxInt64
const minInt = math.MinInt64

var ErrOverflow = errors.New("integer overflow")

func Add(n []string) (int, error) {
	sum := 0

	for _, number := range n {
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
