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

const (
	maxInt = math.MaxInt64
)

var ErrOverflow = errors.New("integer overflow")
var ErrOutsideOfSequence = errors.New("position outside of fibonacci sequence")

func (c Fibonacci) GetNumberFromNumericPosition(position int64) (int64, error) {
	if position <= 0 {
		return 0, ErrOutsideOfSequence
	}

	var (
		fibonacciSequence            = make([]int64, position+1, position+2)
		i, currentNumber, prevNumber int64
	)

	fibonacciSequence[0] = 0
	fibonacciSequence[1] = 1

	for i = 2; i <= position; i++ {
		currentNumber = fibonacciSequence[i]
		prevNumber = fibonacciSequence[i-1]

		if currentNumber > maxInt-prevNumber {
			return 0, ErrOverflow
		}
		fibonacciSequence[i] = fibonacciSequence[i-1] + fibonacciSequence[i-2]
	}

	if position == 1 {
		return fibonacciSequence[0], nil
	}

	return fibonacciSequence[position-1], nil
}
