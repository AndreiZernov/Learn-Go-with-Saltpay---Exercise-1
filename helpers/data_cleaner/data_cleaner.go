package data_cleaner

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/string_remove_duplicates"
	"strings"
)

func DataCleaner(str string) string {
	stringNoNewLine := strings.Replace(str, "\n", ",", -1)
	stringNoWhiteSpace := strings.Replace(stringNoNewLine, " ", "", -1)
	stringNoDuplicates := string_remove_duplicates.RemoveDuplicates(stringNoWhiteSpace)
	return stringNoDuplicates
}
