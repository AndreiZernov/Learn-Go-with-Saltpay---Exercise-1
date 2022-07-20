package temphttp

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciRequestHandler(t *testing.T) {
	fibonacciTestsNoPanic := []struct {
		Name              string
		FibonacciRoute    string
		FibonacciResponse string
	}{
		{Name: "Given the position of the first number in the Fibonacci sequence should return 0", FibonacciRoute: "1", FibonacciResponse: "0 \n"},
		{Name: "Given the position of the third number in the Fibonacci sequence should return 1", FibonacciRoute: "3", FibonacciResponse: "1 \n"},
		{Name: "Given the position 93  should be able to calculate based on int64 limitation", FibonacciRoute: "93", FibonacciResponse: "7540113804746346429 \n"},
	}

	for _, tt := range fibonacciTestsNoPanic {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.FibonacciRoute, nil)
			response := httptest.NewRecorder()

			FibonacciRequestHandler(response, request)

			got := response.Body.String()

			assert.Equal(t, tt.FibonacciResponse, got)
		})
	}

	fibonacciTestsBadRequests := []struct {
		Name                  string
		FibonacciRoute        string
		FibonacciResponseCode int
	}{
		{Name: "Given the any number below or equal 0 should return code 400", FibonacciRoute: "0", FibonacciResponseCode: 400},
		{Name: "Given the position 94  should return overflow error and return code 400", FibonacciRoute: "94", FibonacciResponseCode: 400},
		{Name: "Given negative number should return code 400", FibonacciRoute: "-2", FibonacciResponseCode: 400},
	}

	for _, tt := range fibonacciTestsBadRequests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.FibonacciRoute, nil)
			response := httptest.NewRecorder()

			FibonacciRequestHandler(response, request)

			got := response.Code

			assert.Equal(t, tt.FibonacciResponseCode, got)
		})
	}
}
