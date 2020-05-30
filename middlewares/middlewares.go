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
		bearerToken := r.Header.Get("Authorization")
		isValid, _ := utils.ValidateToken(bearerToken)
		if isValid {
			m := r.Method
			log.Println(m)
			log.Println(r.RemoteAddr)
			next.ServeHTTP(w, r)
		}
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	})
}
