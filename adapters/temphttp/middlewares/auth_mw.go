package middlewares

import (
	"encoding/json"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"log"
	"net/http"
	"os"
	"strings"
)

const envAuthKeysEnvName = "AUTH_KEYS_PATHNAME"

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			authorization    = r.Header.Get("Authorization")
			token            = strings.TrimPrefix(authorization, "Bearer ")
			authKeysPathname = os.Getenv(envAuthKeysEnvName)
		)
		stringOfKeys, err := files.ReadFile(authKeysPathname)
		if err != nil {
			log.Fatal(err)
		}

		sliceOfKeys := strings.Split(stringOfKeys, "\n")

		json.NewEncoder(w).Encode(r)

		switch {
		case slices.Contains(sliceOfKeys, token) && token != "":
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
