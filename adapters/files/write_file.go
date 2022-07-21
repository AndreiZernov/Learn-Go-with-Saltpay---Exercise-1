package files

import (
	"bufio"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"os"
	"path/filepath"
)

func WriteFile(path, data string) {
	path = filepath.Join(Root, path)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	error_handler.HandlePanic(err)

	dataWriter := bufio.NewWriter(file)
	_, err = dataWriter.WriteString(data)
	err = dataWriter.Flush()
	err = file.Close()
	error_handler.HandlePanic(err)
}
