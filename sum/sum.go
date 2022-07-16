package sum

import "strconv"

type Numbers []string

func (n Numbers) add() int {
	sum := 0

	for _, number := range n {
		x, err := strconv.Atoi(number)
		if err == nil {
			sum += x
		}
	}
	return sum
}
