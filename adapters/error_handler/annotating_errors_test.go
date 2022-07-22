package error_handler_test

import (
	"errors"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/internals/testing_helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnnotatingError(t *testing.T) {
	t.Run("Given an error to the AnnotatingError should call error", func(t *testing.T) {
		err := errors.New("failed")
		out := testing_helpers.CaptureOutput(func() {
			error_handler.AnnotatingError(err, "test error")
		})
		assert.Equal(t, "test error: failed\n", out)
	})
	t.Run("Given no error should not print anything", func(t *testing.T) {
		out := testing_helpers.CaptureOutput(func() {
			error_handler.AnnotatingError(nil, "test error")
		})
		assert.Equal(t, "", out)
	})
}
