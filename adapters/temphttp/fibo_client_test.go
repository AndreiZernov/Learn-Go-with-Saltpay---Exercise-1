package temphttp_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestFiboClient(t *testing.T) {
	t.Setenv("SERVER_PORT", "8080")
	t.Setenv("API_ENDPOINT", "http://localhost")
	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)
	t.Setenv(envLogPathname, testAccessLogPathname)

	files.UUIDGenerator(10)
	defer files.RemoveFile(testAuthKeysPathname)
	defer files.RemoveFile(testAccessLogPathname)

	var (
		serverPort  = os.Getenv("SERVER_PORT")
		apiEndpoint = os.Getenv("API_ENDPOINT")
		requestURL  = fmt.Sprintf("%s:%s", apiEndpoint, serverPort)
	)

	fiboClientTests := []struct {
		Name, Argument, ResponseBody string
		ResponseBodyCode             int
	}{
		{
			Name:             "Given the position of the second number in the Fibonacci sequence should return 1",
			Argument:         "2",
			ResponseBody:     "1",
			ResponseBodyCode: http.StatusOK,
		},
	}

	for _, tt := range fiboClientTests {
		t.Run(tt.Name, func(t *testing.T) {
			fiboClient := temphttp.NewFiboClient(requestURL, &http.Client{})
			request, err := fiboClient.Call(tt.Argument)
			response := httptest.NewRecorder()

			if err != nil {
				t.Errorf("FiboClient.Call() error = %v", err)
			}

			router := temphttp.NewRouter()
			router.ServeHTTP(response, request)

			http.ListenAndServe(":"+serverPort, router)
			waitForServer()
			gotBody := response.Body.String()

			assert.Equal(t, tt.ResponseBody, gotBody)
		})
	}
}

func waitForServer() {
	for i := 0; i < 10; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", "8080"))
		if conn != nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
