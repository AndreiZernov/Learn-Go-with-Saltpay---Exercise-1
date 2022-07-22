package middlewares

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const envLogName = "LOG_PATHNAME"

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			start             = time.Now()
			lrw               = newLoggingResponseWriter(w)
			statusCode        = lrw.statusCode
			accessLogPathname = os.Getenv(envLogName)
			authorization     = r.Header.Get("Authorization")
			key               = strings.TrimPrefix(authorization, "Bearer ")
			TimeFormat        = "2006-02-01T15:04:05Z"
		)

		escapedKey := strings.Replace(key, "\n", "", -1)
		escapedKey = strings.Replace(escapedKey, "\r", "", -1)

		if len(escapedKey) > 10 {
			escapedKey = escapedKey[0:10]
		}

		logData := fmt.Sprintf("%s %s %s %s %s %s %d",
			time.Now().Format(TimeFormat),
			r.Method,
			r.RequestURI,
			escapedKey,
			strconv.FormatInt(r.ContentLength, 10),
			strconv.Itoa(statusCode),
			time.Since(start).Milliseconds())
		fmt.Println(logData)

		err := files.WriteFile(accessLogPathname, logData+"\n")
		if err != nil {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r)
	})
}
