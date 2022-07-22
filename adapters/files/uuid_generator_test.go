package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testAuthKeysPathname = "test_authorised_api_access_keys.txt"

func TestUUIDGenerator(t *testing.T) {
	t.Run("Should generate uuid", func(t *testing.T) {
		t.Setenv("AUTH_KEYS_PATHNAME", testAuthKeysPathname)

		err := UUIDGenerator(1)
		data, _ := ReadFile(testAuthKeysPathname)
		defer RemoveFile(testAuthKeysPathname)

		assert.NoError(t, err)
		assert.Equal(t, 37, len(data))
	})
}
