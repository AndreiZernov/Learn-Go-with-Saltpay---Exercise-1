package formatter

import (
	"fmt"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

const (
	max                     = 9999
	min                     = -9999
	thousandCommasFrequency = 3
)

func (f Formatter) GroupsOfThousands(number int64, format bool) string {
	strNumber := fmt.Sprint(number)
	if format && (number > max || number < min) {
		for i := len(strNumber) - thousandCommasFrequency; i > 0; i -= thousandCommasFrequency {
			if strNumber[:i] != "-" {
				strNumber = strNumber[:i] + "," + strNumber[i:]
			}
		}
	}
	return strNumber
}
