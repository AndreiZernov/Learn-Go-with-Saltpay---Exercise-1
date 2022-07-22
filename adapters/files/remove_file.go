package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func RemoveFile(path string) error {
	path = filepath.Join(Root, path)
	err := os.Remove(path)
	return errors.Wrap(err, "failed to remove file")
}
