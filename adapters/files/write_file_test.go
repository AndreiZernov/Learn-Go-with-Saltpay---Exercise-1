package files_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteFile(t *testing.T) {
	t.Run("Should create, write, read and then delete file test_access_log.txt", func(t *testing.T) {
		err := files.WriteFile(testAccessLogPathname, "test-data")
		if err != nil {
			t.Fatal(err)
		}
		data, _ := files.ReadFile(testAccessLogPathname)

		errRemove := files.RemoveFile(testAccessLogPathname)
		if err != nil {
			t.Fatal(errRemove)
		}
		assert.Equal(t, "test-data", data)

	})
}
