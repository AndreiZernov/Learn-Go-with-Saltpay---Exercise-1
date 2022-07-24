package files

import (
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

const envAuthKeyName = "AUTH_KEYS_PATHNAME"
const failedToOpenFileErrorMessage = "failed to open file"

func UUIDGenerator(number int) error {
	var (
		result           []string
		authKeysPathname = os.Getenv(envAuthKeyName)
		path             = filepath.Join(Root, authKeysPathname)
		file, err        = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	)
	if err != nil {
		return errors.Wrap(err, failedToOpenFileErrorMessage)
	}
	defer file.Close()

	for i := 0; i < number; i++ {
		result = append(result, uuid.New()+"\n")
	}

	finalResult := strings.Join(result, "")
	file.WriteString(finalResult)

	return nil
}
