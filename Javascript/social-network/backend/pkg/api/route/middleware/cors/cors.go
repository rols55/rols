package cors

import (
	"net/http"
	"os"
)

var AllowOrigin = func() string {
	if allow, ok := os.LookupEnv("API_ALLOW_ORIGIN"); ok {
		return allow
	}
	return "http://localhost:3000"
}()

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Check if the request origin is allowed
		if origin == AllowOrigin {
			w.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// If it's an OPTIONS request, respond with success and return
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		// Serve the request
		h.ServeHTTP(w, r)
	})
}
