package http

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
				flakinessSlice = strings.Split(flakiness[0], ",")
				probability, _ = strconv.ParseFloat(flakinessSlice[0], 64)
				random         = min + rand.Float64()*(max-min)
				responseStatus = http.StatusInternalServerError
			)

			if len(flakinessSlice) == 2 {
				responseStatus, _ = strconv.Atoi(flakinessSlice[1])
			}

			if random <= probability {
				w.WriteHeader(responseStatus)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
