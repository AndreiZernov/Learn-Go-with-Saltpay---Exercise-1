package middlewares_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	envAuthKeysEnvName   = "AUTH_KEYS_PATHNAME"
	testAuthKeysPathname = "test_authorised_api_access_keys.txt"
)

func TestAuthenticationMiddleware(t *testing.T) {
	pathname := "test_authorised_api_access_keys.txt"
	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)

	files.UUIDGenerator(1)

	data, err := files.ReadFile(pathname)
	if err != nil {
		t.Errorf("Error reading file %s: %v", pathname, err)
	}
	authKeys := strings.Split(data, "\n")
	authKey := authKeys[0]
	defer files.RemoveFile(pathname)

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
			AuthKey:      authKey,
		},
		{
			Name:         "Given a another correct auth key should return 200 response",
			Queries:      "?num=2&num=3",
			ResponseCode: http.StatusOK,
			AuthKey:      authKey,
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
			middlewares.AuthenticationMiddleware(http.HandlerFunc(handlers.AddRequestHandler)).ServeHTTP(response, request)

			gotCode := response.Code

			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
	t.Run("Given no authentication should return 403", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/add?num=2&num=3", nil)
		response := httptest.NewRecorder()

		middlewares.AuthenticationMiddleware(http.HandlerFunc(handlers.AddRequestHandler)).ServeHTTP(response, request)

		gotCode := response.Code

		assert.Equal(t, http.StatusForbidden, gotCode)
	})
}
