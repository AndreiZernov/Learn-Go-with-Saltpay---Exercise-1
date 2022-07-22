package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func FindFile(path string) error {
	path = filepath.Join(Root, path)
	_, err := os.Stat(path)
	return errors.Wrap(err, "failed to find file")
}
