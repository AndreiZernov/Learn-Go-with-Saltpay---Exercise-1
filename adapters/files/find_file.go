package files

import (
	"os"
	"path/filepath"
)

func FindFile(path string) error {
	path = filepath.Join(Root, path)
	_, err := os.Stat(path)
	return err
}
