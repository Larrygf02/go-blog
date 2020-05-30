package middlewares

import (
	"log"
	"net/http"

	"github.com/larrygf02/go-blog/utils"
)

func SetMiddlewareJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "appliction/json")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func TokenMiddlewareJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path != "/login" {
			bearerToken := r.Header.Get("Authorization")
			log.Println(bearerToken)
			isValid, _ := utils.ValidateToken(bearerToken)
			if isValid {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
