package formatter

import (
	"strconv"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

func (f Formatter) GroupsOfThousands(number string) string {
	max := 9999
	min := -9999
	intNumber, _ := strconv.Atoi(number)
	if intNumber > max || intNumber < min {
		for i := len(number) - 3; i > 0; i -= 3 {
			number = number[:i] + "," + number[i:]
		}
	}
	return number
}
