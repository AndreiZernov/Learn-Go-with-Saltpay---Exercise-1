package fibonacci

import (
	"errors"
	"math"
)

type Fibonacci struct {
}

func New() *Fibonacci {
	return &Fibonacci{}
}

var ErrOverflow = errors.New("integer overflow")
var ErrOutsideOfSequence = errors.New("position outside of fibonacci sequence")

func (c Fibonacci) GetNumberFromNumericPosition(position int64) (int64, error) {
	switch {
	case position <= 0:
		return 0, ErrOutsideOfSequence
	case position == 1:
		return 0, nil
	case position >= 94:
		return 0, ErrOverflow
	default:
		numericPosition := position - 1

		firstTerm := math.Pow(math.Phi, float64(numericPosition))
		secondTerm := math.Pow(math.Phi-1, float64(numericPosition))
		result := math.Round((firstTerm + secondTerm) / math.Sqrt(5))

		return int64(result), nil
	}
}
