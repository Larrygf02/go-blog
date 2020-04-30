package middlewares

import "net/http"

func SetMiddlewareJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "appliction/json")
		next.ServeHTTP(w, r)
	})
}
