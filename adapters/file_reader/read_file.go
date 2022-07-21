package file_reader

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func ReadFile(path string) string {
	path = filepath.Join(Root, path)
	content, err := os.ReadFile(path)
	error_handler.HandlePanic(err)
	return string(content)
}
