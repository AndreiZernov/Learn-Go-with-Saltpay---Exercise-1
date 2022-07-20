package middlewares_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationMiddleware(t *testing.T) {
	t.Setenv("AUTH_TOKEN", "SUPER_SECRET_API_KEY_1,SUPER_SECRET_API_KEY_2,SUPER_SECRET_API_KEY_3,SUPER_SECRET_API_KEY_4,SUPER_SECRET_API_KEY_5")

	authenticationMiddlewareTests := []struct {
		Name         string
		Queries      string
		ResponseCode int
		AuthKey      string
	}{
		{
			Name:         "Given a correct auth key should return 200 response",
			Queries:      "?num=2&num=3",
			ResponseCode: http.StatusOK,
			AuthKey:      "SUPER_SECRET_API_KEY_1",
		},
		{
			Name:         "Given a another correct auth key should return 200 response",
			Queries:      "?num=2&num=3",
			ResponseCode: http.StatusOK,
			AuthKey:      "SUPER_SECRET_API_KEY_2",
		},
		{
			Name:         "Given a wrong auth key should return 403 response",
			Queries:      "?num=2&num=3",
			ResponseCode: http.StatusForbidden,
			AuthKey:      "WRONG_API_KEY",
		},
		{
			Name:         "Given a no auth key should return 403 response",
			Queries:      "?num=2&num=3",
			ResponseCode: http.StatusForbidden,
			AuthKey:      "",
		},
	}

	for _, tt := range authenticationMiddlewareTests {
		t.Run(tt.Name, func(t *testing.T) {

			request, _ := http.NewRequest(http.MethodPost, "/add"+tt.Queries, nil)
			response := httptest.NewRecorder()

			request.Header.Set("Authorization", "Bearer "+tt.AuthKey)
			middlewares.AuthenticationMiddleware(http.HandlerFunc(handlers.AddRequestHandlerForQueries)).ServeHTTP(response, request)

			gotCode := response.Code

			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
	t.Run("Given no authentication should return 403", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/add?num=2&num=3", nil)
		response := httptest.NewRecorder()

		middlewares.AuthenticationMiddleware(http.HandlerFunc(handlers.AddRequestHandlerForQueries)).ServeHTTP(response, request)

		gotCode := response.Code

		assert.Equal(t, http.StatusForbidden, gotCode)
	})
}
