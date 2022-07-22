package middlewares

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	min = 0.0
	max = 1.0
)

func FlakinessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		flakiness := r.URL.Query()["flakiness"]

		if len(flakiness) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		var (
			flakinessSlice = strings.Split(flakiness[0], ",")
			probability, _ = strconv.ParseFloat(flakinessSlice[0], 64)
			random         = min + rand.Float64()*(max-min)
			responseStatus = http.StatusInternalServerError
			err            error
		)

		if len(flakinessSlice) >= 2 {
			responseStatus, err = strconv.Atoi(flakinessSlice[1])
			if err != nil {
				responseStatus = http.StatusBadRequest
			}
		}

		if len(flakinessSlice) == 3 {
			delay := flakinessSlice[2]
			parsedDelay, _ := time.ParseDuration(delay)
			time.Sleep(parsedDelay)
		}

		if random <= probability {
			w.WriteHeader(responseStatus)
			return
		}

		next.ServeHTTP(w, r)
	})
}
