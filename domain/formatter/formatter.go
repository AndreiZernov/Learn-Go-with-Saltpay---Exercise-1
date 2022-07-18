package formatter

import (
	"fmt"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

const max = 9999
const min = -9999
const thousandCommasFrequency = 3

func (f Formatter) GroupsOfThousands(number int) string {
	strNumber := fmt.Sprint(number)
	if number > max || number < min {
		for i := len(strNumber) - thousandCommasFrequency; i > 0; i -= thousandCommasFrequency {
			if strNumber[:i] != "-" {
				strNumber = strNumber[:i] + "," + strNumber[i:]
			}
		}
	}
	return strNumber
}
