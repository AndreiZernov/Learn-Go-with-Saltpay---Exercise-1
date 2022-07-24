package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToSliceOfNumbers(t *testing.T) {
	t.Run("Given a slice of strings and a slice of int64 should return a slice of int64", func(t *testing.T) {
		data := []string{"1", "2", "3"}
		expected := []int64{1, 2, 3}
		actual, err := ConvertToSliceOfNumbers(data)
		if err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("Given a slice of strings and a slice of int64 should return an error", func(t *testing.T) {
		data := []string{"1", "2", "3", "a"}
		_, err := ConvertToSliceOfNumbers(data)
		if err == nil {
			t.Errorf("Expected error but got nil")
		}
	})
}
