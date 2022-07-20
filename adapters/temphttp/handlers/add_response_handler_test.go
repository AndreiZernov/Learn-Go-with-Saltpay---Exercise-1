package handlers_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddResponseHandler(t *testing.T) {
	addResponseHandlerTests := []struct {
		Name         string
		Data         []string
		ResponseBody string
		ResponseCode int
	}{
		{
			Name:         "Given one number in body should return the message with the same number",
			Data:         []string{"2"},
			ResponseBody: "Sum of 2 equal 2 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given two numbers in body should return the message with the correct sum of them",
			Data:         []string{"2", "3"},
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given and empty body should return 400",
			Data:         []string{},
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
	}

	for _, tt := range addResponseHandlerTests {
		t.Run(tt.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			handlers.AddResponseHandler(response, tt.Data)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
}
