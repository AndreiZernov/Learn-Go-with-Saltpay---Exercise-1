package string_remove_duplicates

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"strings"
)

func RemoveDuplicates(str string) string {
	var list []string
	newArray := strings.Split(str, ",")

	for _, item := range newArray {
		if array_contains.ArrayContains(list, item) == false {
			list = append(list, item)
		}
	}
	return strings.Join(list[:], ",")
}
