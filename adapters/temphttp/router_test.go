package temphttp_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciRequestHandler(t *testing.T) {
	fibonacciStatus200Tests := []struct {
		Name         string
		queries      string
		responseBody string
	}{
		{Name: "Given the position of the first number in the Fibonacci sequence should return 0", queries: "1", responseBody: "0 \n"},
		{Name: "Given the position of the third number in the Fibonacci sequence should return 1", queries: "3", responseBody: "1 \n"},
		{Name: "Given the position 93  should be able to calculate based on int64 limitation", queries: "93", responseBody: "7540113804746346429 \n"},
	}

	for _, tt := range fibonacciStatus200Tests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.queries, nil)
			response := httptest.NewRecorder()

			temphttp.FibonacciRequestHandler(response, request)

			got := response.Body.String()
			assert.Equal(t, tt.responseBody, got)
		})
	}

	fibonacciTestsBadRequests := []struct {
		Name             string
		queries          string
		responseBodyCode int
	}{
		{Name: "Given the any number below or equal 0 should return code 400", queries: "0", responseBodyCode: 400},
		{Name: "Given the position 94  should return overflow error and return code 400", queries: "94", responseBodyCode: 400},
		{Name: "Given negative number should return code 400", queries: "-2", responseBodyCode: 400},
	}

	for _, tt := range fibonacciTestsBadRequests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.queries, nil)
			response := httptest.NewRecorder()

			temphttp.FibonacciRequestHandler(response, request)

			got := response.Code
			assert.Equal(t, tt.responseBodyCode, got)
		})
	}
}

func TestAddRequestHandlerForQueries(t *testing.T) {
	addRequestsHandlersForQueriesTests := []struct {
		Name         string
		queries      string
		responseBody string
		responseCode int
	}{
		{Name: "Given one number in query should return the message with the same number", queries: "?num=2", responseBody: "Sum of 2 equal 2 \n", responseCode: 200},
		{Name: "Given two numbers in query should return the message with the correct sum of them", queries: "?num=2&num=3", responseBody: "Sum of 2,3 equal 5 \n", responseCode: 200},
		{Name: "Given the wrong query key should ignore it and give the sum of correct one", queries: "?num=2&num=3&sum=20", responseBody: "Sum of 2,3 equal 5 \n", responseCode: 200},
		{Name: "Given the wrong query key only should return 400", queries: "?sum=20", responseBody: "", responseCode: 400},
		{Name: "Given and empty query should return 400", queries: "", responseBody: "", responseCode: 400},
	}

	for _, tt := range addRequestsHandlersForQueriesTests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/add/"+tt.queries, nil)
			response := httptest.NewRecorder()

			temphttp.AddRequestHandlerForQueries(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.responseBody, gotBody)
			assert.Equal(t, tt.responseCode, gotCode)
		})
	}
}
