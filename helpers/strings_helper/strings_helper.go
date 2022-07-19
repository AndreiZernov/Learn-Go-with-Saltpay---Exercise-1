package strings_helper

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"strings"
)

func DataCleaner(stringToClean string) string {
	stringNoNewLine := strings.Replace(stringToClean, "\n", ",", -1)
	stringNoWhiteSpace := strings.Replace(stringNoNewLine, " ", "", -1)
	cleanedString := RemoveDuplicates(stringNoWhiteSpace)
	return cleanedString
}

func RemoveDuplicates(stringToClean string) string {
	var list []string
	newArray := strings.Split(stringToClean, ",")

	for _, item := range newArray {
		if slices.Contains(list, item) == false {
			list = append(list, item)
		}
	}
	return strings.Join(list[:], ",")
}
