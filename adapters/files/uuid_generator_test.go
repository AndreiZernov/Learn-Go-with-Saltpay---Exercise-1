package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {

	t.Run("Should generate uuid", func(t *testing.T) {
		pathname := "test_authorised_api_access_keys.txt"
		t.Setenv("AUTH_KEYS_PATHNAME", pathname)

		err := UUIDGenerator()
		data := ReadFile(pathname)
		defer RemoveFile(pathname)

		assert.Equal(t, nil, err)
		assert.Equal(t, 37, len(data))
	})
}
