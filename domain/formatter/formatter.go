package formatter

import (
	"fmt"
)

type Formatter struct {
}

func New() *Formatter {
	return &Formatter{}
}

func (f Formatter) GroupsOfThousands(number int) string {
	max := 9999
	min := -9999
	strNumber := fmt.Sprint(number)
	if number > max || number < min {
		for i := len(strNumber) - 3; i > 0; i -= 3 {
			if strNumber[:i] != "-" {
				strNumber = strNumber[:i] + "," + strNumber[i:]
			}
		}
	}
	return strNumber
}
