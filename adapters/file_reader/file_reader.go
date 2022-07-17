package file_reader

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/string_replace"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ReadFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string_replace.StringReplace(string(content))
}

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func GetFilePathname(path string) string {
	return filepath.Join(Root, path)
}
