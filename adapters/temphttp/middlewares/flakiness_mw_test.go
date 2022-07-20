package middlewares_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFlakinessMiddleware(t *testing.T) {
	flakinessMiddlewareTests := []struct {
		Name              string
		queries           string
		responseCode      int
		expectedDuraction int
	}{
		{
			Name:              "Given a flakiness probability 1 should return 500 response",
			queries:           "?flakiness=1",
			responseCode:      http.StatusInternalServerError,
			expectedDuraction: 0,
		},
		{
			Name:              "Given a flakiness probability 0 should return 200 response",
			queries:           "?num=2&num=3&flakiness=0",
			responseCode:      http.StatusOK,
			expectedDuraction: 0,
		},
		{
			Name:              "Given the flakiness with specific error 404 should return 404 response",
			queries:           "?num=2&num=3&flakiness=1,404",
			responseCode:      http.StatusNotFound,
			expectedDuraction: 0,
		},
		{
			Name:              "Given the flakiness with specific error 400 should return 400 response",
			queries:           "?num=2&num=3&flakiness=1,400",
			responseCode:      http.StatusBadRequest,
			expectedDuraction: 0,
		},
		{
			Name:              "Given a delay of 2s should be delayed for 2s",
			queries:           "?flakiness=1,404,2s",
			responseCode:      http.StatusNotFound,
			expectedDuraction: 2,
		},
		{
			Name:              "Given a delay and the OK response should be delayed for 1s and then return 200",
			queries:           "?num=2&num=3&flakiness=1,200,1s",
			responseCode:      http.StatusOK,
			expectedDuraction: 1,
		},
	}

	for _, tt := range flakinessMiddlewareTests {
		t.Run(tt.Name, func(t *testing.T) {
			start := time.Now()
			request, _ := http.NewRequest(http.MethodPost, "/add/"+tt.queries, nil)
			response := httptest.NewRecorder()

			middlewares.FlakinessMiddleware(http.HandlerFunc(handlers.AddRequestHandlerForQueries)).ServeHTTP(response, request)

			gotCode := response.Code
			duration := int(time.Since(start).Seconds())

			assert.Equal(t, tt.expectedDuraction, duration)
			assert.Equal(t, tt.responseCode, gotCode)
		})
	}

}
