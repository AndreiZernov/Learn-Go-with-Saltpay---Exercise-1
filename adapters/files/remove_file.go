package files

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"os"
	"path/filepath"
)

func RemoveFile(path string) {
	path = filepath.Join(Root, path)
	err := os.Remove(path)
	error_handler.HandlePanic(err)
}
