package files

import (
	"bufio"
	"github.com/pborman/uuid"
	"os"
	"path/filepath"
	"strings"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"

func UUIDGenerator(number int) {
	authKeysPathname := os.Getenv(envAuthKeysPathname)
	path := filepath.Join(Root, authKeysPathname)
	file, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	dataWriter := bufio.NewWriter(file)

	var result []string
	for i := 0; i < number; i++ {
		result = append(result, uuid.New()+"\n")
	}
	finalResult := strings.Join(result, "")
	dataWriter.WriteString(finalResult)
	dataWriter.Flush()
	file.Close()
}
