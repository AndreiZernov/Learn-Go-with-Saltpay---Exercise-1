package commandLine

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/string_replace"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
	numbers    []string
)

func ReadFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string_replace.StringReplace(string(content))
}

func GetArguments() []string {
	toGetAllArgs := os.Args[1:]

	switch {
	case len(toGetAllArgs) == 0:
		filepath := filepath.Join(Root, "/data/", "input.txt")
		numbers = []string{ReadFile(filepath)}
	case array_contains.ArrayContains(toGetAllArgs, "--input-file"):
		var path string

		for i := 0; i < len(toGetAllArgs); i++ {
			if toGetAllArgs[i] == "--input-file" {
				path = filepath.Join(Root, "/", toGetAllArgs[i+1])
				numbers = append(numbers, ReadFile(path))
			}
		}
	default:
		numbers = toGetAllArgs
	}

	return numbers
}
