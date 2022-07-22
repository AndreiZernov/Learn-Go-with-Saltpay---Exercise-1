package testing_helpers_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/internals/testing_helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaptureOutput(t *testing.T) {
	t.Run("Should capture output", func(t *testing.T) {
		out := testing_helpers.CaptureOutput(func() {
			fmt.Println("test")
		})
		assert.Equal(t, "test\n", out)
	})
}
