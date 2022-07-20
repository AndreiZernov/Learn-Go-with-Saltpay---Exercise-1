package error_handler_test

import (
	"errors"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
