package file_reader

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ReadFile(filepath string) string {
	path := GetFilePathname(filepath)
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func GetFilePathname(path string) string {
	return filepath.Join(Root, path)
}
