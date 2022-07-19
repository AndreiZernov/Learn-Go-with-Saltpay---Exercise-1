package http

import (
	"encoding/json"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"net/http"
	"strings"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authorization = r.Header.Get("Authorization")
			token         = strings.Replace(authorization, "Bearer ", "", -1)
		)

		err := json.NewEncoder(w).Encode(r)
		error_handler.HandlePanic(err)

		switch token {
		case "SUPER_SECRET_API_KEY":
			next.ServeHTTP(w, r)
		case "":
			w.WriteHeader(http.StatusForbidden)
			err := json.NewEncoder(w).Encode("Missing auth token")
			error_handler.HandlePanic(err)
			return
		default:
			w.WriteHeader(http.StatusForbidden)
			err := json.NewEncoder(w).Encode("Invalid token")
			error_handler.HandlePanic(err)
			return
		}
	})
}
