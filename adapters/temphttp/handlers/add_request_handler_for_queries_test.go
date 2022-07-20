package handlers_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddRequestHandlerForQueries(t *testing.T) {
	addRequestsHandlersForQueriesTests := []struct {
		Name         string
		queries      string
		ResponseBody string
		ResponseCode int
	}{
		{
			Name:         "Given one number in query should return the message with the same number",
			queries:      "?num=2",
			ResponseBody: "Sum of 2 equal 2 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given two numbers in query should return the message with the correct sum of them",
			queries:      "?num=2&num=3",
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong query key should ignore it and give the sum of correct one",
			queries:      "?num=2&num=3&wrongNum=20",
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong query key only should return 400",
			queries:      "?wrongNum=20",
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Given and empty query should return 400",
			queries:      "",
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
	}

	for _, tt := range addRequestsHandlersForQueriesTests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, "/add/"+tt.queries, nil)
			response := httptest.NewRecorder()

			handlers.AddRequestHandlerForQueries(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
}
