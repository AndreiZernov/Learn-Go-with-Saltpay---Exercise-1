package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			start         = time.Now()
			lrw           = newLoggingResponseWriter(w)
			statusCode    = lrw.statusCode
			authorization = r.Header.Get("Authorization")
			token         = strings.TrimPrefix(authorization, "Bearer ")
		)

		log.Println(
			r.Method,
			r.RequestURI,
			token[0:9],
			strconv.FormatInt(r.ContentLength, 10),
			strconv.Itoa(statusCode),
			time.Since(start).Milliseconds(),
		)

		next.ServeHTTP(w, r)
	})
}
