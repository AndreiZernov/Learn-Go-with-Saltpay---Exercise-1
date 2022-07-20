package http

import (
	"math/rand"
	"net/http"
	"strconv"
)

const (
	min = 0.0
	max = 1.0
)

func flakinessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flakiness := r.URL.Query()["flakiness"]
		if len(flakiness) != 0 {
			var (
				probability, _ = strconv.ParseFloat(flakiness[0], 64)
				random         = min + rand.Float64()*(max-min)
			)

			if random <= probability {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
