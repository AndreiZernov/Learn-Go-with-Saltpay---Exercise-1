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
		Name             string
		Queries          string
		ResponseCode     int
		ExpectedDuration int
	}{
		{
			Name:             "Given a flakiness probability 1 should return 500 response",
			Queries:          "?flakiness=1",
			ResponseCode:     http.StatusInternalServerError,
			ExpectedDuration: 0,
		},
		{
			Name:             "Given a flakiness probability 0 should return 200 response",
			Queries:          "?num=2&num=3&flakiness=0",
			ResponseCode:     http.StatusOK,
			ExpectedDuration: 0,
		},
		{
			Name:             "Given the flakiness with specific error 404 should return 404 response",
			Queries:          "?num=2&num=3&flakiness=1,404",
			ResponseCode:     http.StatusNotFound,
			ExpectedDuration: 0,
		},
		{
			Name:             "Given a delay of 1s should be delayed for 1s",
			Queries:          "?flakiness=1,404,1s",
			ResponseCode:     http.StatusNotFound,
			ExpectedDuration: 1,
		},
	}

	for _, tt := range flakinessMiddlewareTests {
		t.Run(tt.Name, func(t *testing.T) {
			start := time.Now()
			request, _ := http.NewRequest(http.MethodPost, "/add/"+tt.Queries, nil)
			response := httptest.NewRecorder()

			middlewares.FlakinessMiddleware(http.HandlerFunc(handlers.AddRequestHandler)).ServeHTTP(response, request)

			gotCode := response.Code
			duration := int(time.Since(start).Seconds())

			assert.Equal(t, tt.ExpectedDuration, duration)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}

}
