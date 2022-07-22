package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func ReadFile(path string) (string, error) {
	path = filepath.Join(Root, path)
	content, err := os.ReadFile(path)
	return string(content), errors.Wrap(err, "failed to read file")
}
