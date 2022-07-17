package string_replace

import "strings"

func StringReplace(str string) string {
	stringNoNewLine := strings.Replace(str, "\n", ",", -1)
	stringNoWhiteSpace := strings.Replace(stringNoNewLine, " ", "", -1)
	return stringNoWhiteSpace
}
