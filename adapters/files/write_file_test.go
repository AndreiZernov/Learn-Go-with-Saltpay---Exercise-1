package files_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	t.Run("Should create, write, read and then delete file test_access_log.txt", func(t *testing.T) {
		pathname := "adapters/files/test_access_log.txt"
		files.WriteFile(pathname, "test-data")
		data := files.ReadFile(pathname)
		path := filepath.Join(files.Root, pathname)
		os.Remove(path)
		assert.Equal(t, "test-data\n", data)
	})
}
