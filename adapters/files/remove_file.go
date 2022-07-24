package files

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const failedToRemoveFileErrorMessage = "failed to remove file"

func RemoveFile(path string) error {
	path = filepath.Join(Root, path)
	err := os.Remove(path)
	return errors.Wrap(err, failedToRemoveFileErrorMessage)
}
