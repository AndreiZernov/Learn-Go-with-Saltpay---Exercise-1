package strings_helper

import (
	"strings"
)

func DataCleaner(stringToClean string) string {
	stringNoNewLine := strings.Replace(stringToClean, "\n", ",", -1)
	stringNoWhiteSpace := strings.Replace(stringNoNewLine, " ", "", -1)
	cleanedString := RemoveDuplicates(stringNoWhiteSpace)
	return cleanedString
}
