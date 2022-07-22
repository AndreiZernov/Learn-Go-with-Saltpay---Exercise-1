package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testAuthKeysPathname = "test_authorised_api_access_keys.txt"

func TestUUIDGenerator(t *testing.T) {
	t.Run("Should generate uuid", func(t *testing.T) {
		t.Setenv("AUTH_KEYS_PATHNAME", testAuthKeysPathname)

		UUIDGenerator(1)
		data := ReadFile(testAuthKeysPathname)
		defer RemoveFile(testAuthKeysPathname)

		assert.Equal(t, 37, len(data))
	})
}
