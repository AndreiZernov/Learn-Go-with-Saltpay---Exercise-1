package middlewares_test

import (
	"bytes"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
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
			out := captureOutput(func() {
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

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
