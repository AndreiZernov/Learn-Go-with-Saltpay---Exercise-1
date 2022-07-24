package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"runtime"
)

const failedToReadFileErrorMessage = "failed to read file"

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func ReadFile(path string) (string, error) {
	path = filepath.Join(Root, path)
	content, err := os.ReadFile(path)
	return string(content), errors.Wrap(err, failedToReadFileErrorMessage)
}
