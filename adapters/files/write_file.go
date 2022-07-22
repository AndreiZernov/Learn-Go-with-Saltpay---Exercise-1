package files

import (
	"bufio"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"os"
	"path/filepath"
)

func WriteFile(path, data string) {
	path = filepath.Join(Root, path)
	file, openErr := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	error_handler.AnnotatingError(openErr, "Failed to open file")
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	dataWriter.WriteString(data)
	dataWriter.Flush()
}
