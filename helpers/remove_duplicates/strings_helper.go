package strings_helper

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"strings"
)

func DataCleaner(str string) string {
	stringNoNewLine := strings.Replace(str, "\n", ",", -1)
	stringNoWhiteSpace := strings.Replace(stringNoNewLine, " ", "", -1)
	stringNoDuplicates := RemoveDuplicates(stringNoWhiteSpace)
	return stringNoDuplicates
}

func RemoveDuplicates(str string) string {
	var list []string
	newArray := strings.Split(str, ",")

	for _, item := range newArray {
		if slices.Contains(list, item) == false {
			list = append(list, item)
		}
	}
	return strings.Join(list[:], ",")
}
