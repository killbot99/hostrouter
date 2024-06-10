package middleware

import (
	"log"
	"net/http"
)

func Routing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Routing")
		next.ServeHTTP(w, r)
	})
}
