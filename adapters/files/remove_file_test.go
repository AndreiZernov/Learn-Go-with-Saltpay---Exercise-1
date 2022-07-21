package files_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveFile(t *testing.T) {
	t.Run("Should remove file", func(t *testing.T) {
		pathname := "adapters/files/test_access_log.txt"

		files.WriteFile(pathname, "test-data")
		files.RemoveFile(pathname)

		path := filepath.Join(files.Root, pathname)
		_, err := os.Stat(path)
		assert.Equal(t, true, os.IsNotExist(err))
	})

	t.Run("Should exist file", func(t *testing.T) {
		pathname := "adapters/files/test_access_log.txt"

		files.WriteFile(pathname, "test-data")
		defer files.RemoveFile(pathname)

		path := filepath.Join(files.Root, pathname)
		_, err := os.Stat(path)
		assert.Equal(t, false, os.IsNotExist(err))
	})
}
