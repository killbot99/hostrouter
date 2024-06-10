package middleware

import (
	"log"
	"net/http"
)

// Logging is logger middleware
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging")
		next.ServeHTTP(w, r)
	})
}
