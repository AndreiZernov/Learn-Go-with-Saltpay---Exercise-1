package files

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func WriteFile(path, data string) error {
	path = filepath.Join(Root, path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, failedToOpenFileErrorMessage)
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	dataWriter.WriteString(data)
	dataWriter.Flush()
	return errors.Wrap(err, failedToReadFileErrorMessage)
}
