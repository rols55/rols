package method

import (
	"net/http"

	"forum/shared/logger"
)

func Method(args ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if args == nil {
				next.ServeHTTP(w, r)
				return
			}
			for _, m := range args {
				if m == r.Method {
					next.ServeHTTP(w, r)
					return
				}
			}
			logger.Info("Request denied for method: '%v'", r.Method)
			http.Error(w, "Not Found 404", http.StatusNotFound)
		})
	}
}
