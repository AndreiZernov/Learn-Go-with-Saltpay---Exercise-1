package middlewares

import (
	"encoding/json"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"net/http"
	"os"
	"strings"
)

const envAuthKeysPathname = "AUTH_KEYS_PATHNAME"

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authorization    = r.Header.Get("Authorization")
			token            = strings.TrimPrefix(authorization, "Bearer ")
			authKeysPathname = os.Getenv(envAuthKeysPathname)
			stringOfKeys     = files.ReadFile(authKeysPathname)
			sliceOfKeys      = strings.Split(stringOfKeys, "\n")
		)
		err := json.NewEncoder(w).Encode(r)
		error_handler.HandlePanic(err)

		switch {
		case slices.Contains(sliceOfKeys, token) && token != "":
			next.ServeHTTP(w, r)
		case token == "":
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
