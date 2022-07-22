package handlers_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciRequestHandler(t *testing.T) {
	fibonacciRequestHandlerTests := []struct {
		Name, Queries, ResponseBody string
		ResponseBodyCode            int
	}{
		{
			Name:             "Given the position of the first number in the Fibonacci sequence should return 0",
			Queries:          "1",
			ResponseBody:     "0",
			ResponseBodyCode: http.StatusOK,
		},
		{
			Name:             "Given the position of the third number in the Fibonacci sequence should return 1",
			Queries:          "3",
			ResponseBody:     "1",
			ResponseBodyCode: http.StatusOK,
		},
		{
			Name:             "Given the position 93 should be able to calculate based on int64 limitation",
			Queries:          "93",
			ResponseBody:     "7540113804746344448",
			ResponseBodyCode: http.StatusOK,
		},
		{
			Name:             "Given the any number below or equal 0 should return code 400",
			Queries:          "0",
			ResponseBody:     "0",
			ResponseBodyCode: http.StatusBadRequest,
		},
		{
			Name:             "Given the position 94  should return overflow error and return code 400",
			Queries:          "94",
			ResponseBody:     "0",
			ResponseBodyCode: http.StatusBadRequest,
		},
		{
			Name:             "Given negative number should return code 400",
			Queries:          "-2",
			ResponseBody:     "0",
			ResponseBodyCode: http.StatusBadRequest,
		},
	}

	for _, tt := range fibonacciRequestHandlerTests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.Queries, nil)
			response := httptest.NewRecorder()

			handlers.FibonacciRequestHandler(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseBodyCode, gotCode)
		})
	}
}
