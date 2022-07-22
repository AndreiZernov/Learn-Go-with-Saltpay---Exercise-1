package files_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testAccessLogPathname = "adapters/files/test_access_log.txt"

func TestRemoveFile(t *testing.T) {
	t.Run("Should remove file", func(t *testing.T) {
		files.WriteFile(testAccessLogPathname, "test-data")
		files.RemoveFile(testAccessLogPathname)

		err := files.FindFile(testAccessLogPathname)
		assert.Error(t, err)
	})

	t.Run("Should exist file", func(t *testing.T) {
		pathname := testAccessLogPathname

		files.WriteFile(pathname, "test-data")
		defer files.RemoveFile(pathname)

		err := files.FindFile(pathname)
		assert.NoError(t, err)
	})
}
