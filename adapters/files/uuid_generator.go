package files

import (
	"bufio"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

const envAuthKeysEnvName = "AUTH_KEYS_PATHNAME"

func UUIDGenerator(number int) error {
	var (
		result           []string
		authKeysPathname = os.Getenv(envAuthKeysEnvName)
		path             = filepath.Join(Root, authKeysPathname)
		file, err        = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)

	for i := 0; i < number; i++ {
		result = append(result, uuid.New()+"\n")
	}

	finalResult := strings.Join(result, "")
	dataWriter.WriteString(finalResult)
	dataWriter.Flush()

	return nil
}
