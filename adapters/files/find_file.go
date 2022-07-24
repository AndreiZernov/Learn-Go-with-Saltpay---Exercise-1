package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const failedToFindFileErrorMessage = "failed to find file"

func FindFile(path string) error {
	path = filepath.Join(Root, path)
	_, err := os.Stat(path)
	return errors.Wrap(err, failedToFindFileErrorMessage)
}
