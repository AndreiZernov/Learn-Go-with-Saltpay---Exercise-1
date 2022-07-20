package http

import (
	"encoding/json"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"net/http"
	"os"
	"strings"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authorization = r.Header.Get("Authorization")
			token         = strings.TrimPrefix(authorization, "Bearer ")
			stringOfKeys  = os.Getenv("AUTH_TOKEN")
			sliceOfKeys   = strings.Split(stringOfKeys, ",")
		)
		json.NewEncoder(w).Encode(r)

		switch {
		case slices.Contains(sliceOfKeys, token):
			next.ServeHTTP(w, r)
		case token == "":
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		default:
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}
	})
}
