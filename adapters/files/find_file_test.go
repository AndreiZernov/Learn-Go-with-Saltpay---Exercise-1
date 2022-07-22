package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFile(t *testing.T) {
	t.Run("should return error if file not found", func(t *testing.T) {
		err := FindFile("not_found.txt")
		assert.Error(t, err)
	})
	t.Run("should return nil if file found", func(t *testing.T) {
		err := FindFile("README.md")
		assert.NoError(t, err)
	})
}
