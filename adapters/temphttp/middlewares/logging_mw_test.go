package middlewares_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/internals/testing_helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	loggingMiddlewareTests := []struct {
		Name         string
		Queries      string
		ResponseCode int
		RequestURI   string
	}{
		{
			Name:       "Given a flakiness probability 1 should return 500 response",
			RequestURI: "/add",
		},
		{
			Name:       "Given a flakiness probability 0 should return 200 response",
			RequestURI: "/fibonacci",
		},
	}

	for _, tt := range loggingMiddlewareTests {
		t.Run(tt.Name, func(t *testing.T) {
			response := httptest.NewRecorder()
			out := testing_helpers.CaptureOutput(func() {
				request, _ := http.NewRequest(http.MethodPost, tt.RequestURI, nil)

				request.Header.Set("Authorization", "Bearer SUPER_SECRET_API_KEY_1")
				request.RequestURI = tt.RequestURI
				middlewares.LoggingMiddleware(http.HandlerFunc(handlers.AddRequestHandler)).ServeHTTP(response, request)
			})

			outputSlice := strings.Split(out, " ")
			assert.Equal(t, tt.RequestURI, outputSlice[2])

		})
	}

}
