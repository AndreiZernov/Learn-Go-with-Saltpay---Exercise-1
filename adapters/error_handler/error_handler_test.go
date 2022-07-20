package error_handler_test

import (
	"errors"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorHandler(t *testing.T) {
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

func TestHandleStatusBadRequest(t *testing.T) {
	t.Run("Given a bad request should return a 400 request status", func(t *testing.T) {
		response := httptest.NewRecorder()

		err := errors.New("400 Bad Request")

		error_handler.HandleStatusBadRequest(response, err)
		got := response.Code

		assert.Equal(t, http.StatusBadRequest, got)
	})

	t.Run("Given a request without error should return a 200 request status", func(t *testing.T) {
		response := httptest.NewRecorder()

		error_handler.HandleStatusBadRequest(response, nil)
		got := response.Code

		assert.Equal(t, http.StatusOK, got)
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
