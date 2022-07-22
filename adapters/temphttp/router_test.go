package temphttp_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const envAuthKeysEnvName = "AUTH_KEYS_PATHNAME"
const envLogPathname = "LOG_PATHNAME"
const testAuthKeysPathname = "test_authorised_api_access_keys.txt"
const testAccessLogPathname = "adapters/files/test_access_log.txt"

func TestRouter(t *testing.T) {
	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)
	t.Setenv(envLogPathname, testAccessLogPathname)

	files.UUIDGenerator(1)
	defer files.RemoveFile(testAuthKeysPathname)

	routerTests := []struct {
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
			ResponseBody:     "",
			ResponseBodyCode: http.StatusBadRequest,
		},
		{
			Name:             "Given the position 94  should return overflow error and return code 400",
			Queries:          "94",
			ResponseBody:     "",
			ResponseBodyCode: http.StatusBadRequest,
		},
		{
			Name:             "Given negative number should return code 400",
			Queries:          "-2",
			ResponseBody:     "",
			ResponseBodyCode: http.StatusBadRequest,
		},
	}

	for _, tt := range routerTests {
		t.Run(tt.Name, func(t *testing.T) {
			authKeys, err := files.ReadFile(testAuthKeysPathname)
			if err != nil {
				t.Errorf("Cannot read auth keys file")
			}
			authKey := strings.Split(authKeys, "\n")[0]

			request, _ := http.NewRequest(http.MethodGet, "/fibonacci/"+tt.Queries, nil)
			request.Header.Set("Authorization", authKey)
			response := httptest.NewRecorder()

			temphttp.NewRouter().ServeHTTP(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseBodyCode, gotCode)
		})
	}
}
