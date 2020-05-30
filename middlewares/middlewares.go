package middlewares

import (
	"log"
	"net/http"
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
		m := r.Method
		log.Println(m)
		log.Println(r.RemoteAddr)
	})
}
