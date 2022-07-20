package error_handler_test

import (
	"errors"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlePanic(t *testing.T) {
	t.Run("Given an error to the HandlePanic should call panic and return true", func(t *testing.T) {
		err := errors.New("failed")
		got := testPanic(func() { error_handler.HandlePanic(err) })
		assert.True(t, got)
	})

	t.Run("Given a no error to the HandlePanic should not call panic and return false", func(t *testing.T) {
		got := testPanic(func() { fmt.Println("") })
		assert.False(t, got)
	})
}

func testPanic(testFunc func()) (isPanic bool) {
	defer func() {
		if err := recover(); err != nil {
			isPanic = true
		}
	}()
	testFunc()
	return false
}
